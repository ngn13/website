package status

import (
	"net/http"
	"net/http/httptrace"
	"net/url"
	"time"

	"github.com/ngn13/website/api/database"
	"github.com/ngn13/website/api/util"
)

const (
	STATUS_RES_DOWN = 0 // service is down
	STATUS_RES_OK   = 1 // service is up
	STATUS_RES_SLOW = 2 // service is up, but slow
	STATUS_RES_NONE = 3 // service doesn't support status checking/status checking is disabled
)

func (s *Type) check_http_service(service *database.Service) (r uint8, err error) {
	var (
		req *http.Request
		res *http.Response

		start   time.Time
		elapsed time.Duration
	)

	r = STATUS_RES_NONE

	if req, err = http.NewRequest("GET", service.CheckURL, nil); err != nil {
		return
	}

	trace := &httptrace.ClientTrace{
		GetConn:              func(_ string) { start = time.Now() },
		GotFirstResponseByte: func() { elapsed = time.Since(start) },
	}

	http.DefaultClient.Timeout = s.timeout
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	res, err = http.DefaultClient.Do(req)

	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		util.Debg("marking service \"%s\" as down (%s)", service.Name, err.Error())
		err = nil
		r = STATUS_RES_DOWN
	} else if res.StatusCode != 200 {
		util.Debg("marking service \"%s\" as down (status code %d)", service.Name, res.StatusCode)
		r = STATUS_RES_DOWN
	} else if elapsed.Microseconds() > s.limit.Microseconds() {
		r = STATUS_RES_SLOW
	} else {
		r = STATUS_RES_OK
	}

	return
}

func (s *Type) check_service(service *database.Service) error {
	var (
		res uint8
		url *url.URL
		err error
	)

	if s.disabled || service.CheckURL == "" {
		err = nil
		goto fail
	}

	if url, err = url.Parse(service.CheckURL); err != nil {
		return err
	}

	switch url.Scheme {
	case "https":
		if res, err = s.check_http_service(service); err != nil {
			goto fail
		}

	case "http":
		if res, err = s.check_http_service(service); err != nil {
			goto fail
		}

	default:
		// unsupported protocol
		err = nil
		goto fail
	}

	service.CheckTime = uint64(time.Now().Unix())
	service.CheckRes = res
	return nil

fail:
	service.CheckTime = 0
	service.CheckRes = STATUS_RES_NONE
	return err
}
