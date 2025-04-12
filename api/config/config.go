package config

import (
	"fmt"
	"net/url"

	"github.com/ngn13/ortam"
)

type Type struct {
	Debug    bool     // should display debug messgaes?
	AppUrl   *url.URL // frontend application URL for the website
	Password string   // admin password
	Host     string   // host the server should listen on
	IPHeader string   // header that should be checked for obtaining the client IP
	Interval string   // service status check interval
	Timeout  string   // timeout for the service status check
	Limit    string   // if the service responds slower than this limit, it will be marked as "slow"
}

func Load() (*Type, error) {
	var conf = Type{
		Debug:    false,
		Password: "",
		Host:     "0.0.0.0:7002",
		IPHeader: "X-Real-IP",
		Interval: "1h",
		Timeout:  "15s",
		Limit:    "5s",
	}

	if err := ortam.Load(&conf, "WEBSITE"); err != nil {
		return nil, err
	}

	if conf.AppUrl == nil {
		conf.AppUrl, _ = url.Parse("http://localhost:7001/")
	}

	if conf.Password == "" {
		return nil, fmt.Errorf("password is not specified")
	}

	if conf.Host == "" {
		return nil, fmt.Errorf("host address is not specified")
	}

	return &conf, nil
}
