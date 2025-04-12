package status

import (
	"fmt"
	"time"

	"github.com/ngn13/website/api/config"
	"github.com/ngn13/website/api/database"
	"github.com/ngn13/website/api/util"
)

type Type struct {
	conf *config.Type
	db   *database.Type

	ticker     *time.Ticker
	updateChan chan int
	closeChan  chan int

	disabled bool
	timeout  time.Duration
	limit    time.Duration
}

func (s *Type) check() {
	var (
		services []database.Service
		service  database.Service
		err      error
	)

	for s.db.ServiceNext(&service) {
		services = append(services, service)
	}

	for i := range services {
		if err = s.check_service(&services[i]); err != nil {
			util.Fail("failed to check the service status for \"%s\": %s", services[i].Name, err.Error())
		}

		if err = s.db.ServiceUpdate(&services[i]); err != nil {
			util.Fail("failed to update service status for \"%s\": %s", services[i].Name, err.Error())
		}
	}
}

func (s *Type) loop() {
	s.check()

	for {
		select {
		case <-s.closeChan:
			close(s.updateChan)
			s.ticker.Stop()
			s.closeChan <- 0
			return

		case <-s.updateChan:
			s.check()

		case <-s.ticker.C:
			s.check()
		}
	}
}

func (s *Type) Setup(conf *config.Type, db *database.Type) error {
	var (
		dur time.Duration
		err error
	)

	if conf.Interval == "" || conf.Timeout == "" || conf.Limit == "" {
		s.disabled = true
		return nil
	}

	if dur, err = util.GetDuration(conf.Interval); err != nil {
		return err
	}

	if s.timeout, err = util.GetDuration(conf.Timeout); err != nil {
		return err
	}

	if s.limit, err = util.GetDuration(conf.Limit); err != nil {
		return err
	}

	s.conf = conf
	s.db = db

	s.ticker = time.NewTicker(dur)
	s.updateChan = make(chan int)
	s.closeChan = make(chan int)

	s.disabled = false

	return nil
}

func (s *Type) Run() error {
	if s.ticker == nil || s.updateChan == nil || s.closeChan == nil {
		return fmt.Errorf("you either didn't call Setup() or you called it and it failed")
	}

	if s.disabled {
		go s.check()
		return nil
	}

	go s.loop()
	return nil
}

func (s *Type) Check() {
	if !s.disabled {
		s.updateChan <- 0
	}
}

func (s *Type) Stop() {
	// tell loop() to stop
	s.closeChan <- 0

	// wait till loop() stops
	for {
		select {
		case <-s.closeChan:
			close(s.closeChan)
			return
		}
	}
}
