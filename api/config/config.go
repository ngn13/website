package config

import (
	"fmt"
	"net/url"
	"os"
)

type Type struct {
	Options []Option
	Count   int
}

func (c *Type) Find(name string, typ uint8) (*Option, error) {
	for i := 0; i < c.Count; i++ {
		if c.Options[i].Name != name {
			continue
		}

		if c.Options[i].Type != typ {
			return nil, fmt.Errorf("bad option type")
		}

		return &c.Options[i], nil
	}

	return nil, fmt.Errorf("option not found")
}

func (c *Type) Load() (err error) {
	var (
		env_val  string
		env_name string
		opt      *Option
		exists   bool
	)

	// default options
	c.Options = []Option{
		{Name: "debug", Value: "false", Type: OPTION_TYPE_BOOL, Required: true},                         // should display debug messgaes?
		{Name: "app_url_clear", Value: "http://localhost:7001/", Type: OPTION_TYPE_URL, Required: true}, // frontend application URL for the website
		{Name: "password", Value: "", Type: OPTION_TYPE_STR, Required: true},                            // admin password
		{Name: "host", Value: "0.0.0.0:7002", Type: OPTION_TYPE_STR, Required: true},                    // host the server should listen on
		{Name: "ip_header", Value: "X-Real-IP", Type: OPTION_TYPE_STR, Required: false},                 // header that should be checked for obtaining the client IP
		{Name: "interval", Value: "1h", Type: OPTION_TYPE_STR, Required: false},                         // service status check interval
		{Name: "timeout", Value: "15s", Type: OPTION_TYPE_STR, Required: false},                         // timeout for the service status check
		{Name: "limit", Value: "5s", Type: OPTION_TYPE_STR, Required: false},                            // if the service responds slower than this limit, it will be marked as "slow"
	}
	c.Count = len(c.Options)

	for i := 0; i < c.Count; i++ {
		opt = &c.Options[i]

		env_name = opt.Env()

		if env_val, exists = os.LookupEnv(env_name); exists {
			opt.Value = env_val
		}

		if opt.Value == "" && opt.Required {
			return fmt.Errorf("please specify a value for the config option \"%s\" (\"%s\")", opt.Name, env_name)
		}

		if err = opt.Load(); err != nil {
			return fmt.Errorf("failed to load option \"%s\" (\"%s\"): %s", opt.Name, env_name, err.Error())
		}
	}

	return nil
}

func (c *Type) GetStr(name string) string {
	var (
		opt *Option
		err error
	)

	if opt, err = c.Find(name, OPTION_TYPE_STR); err != nil {
		return ""
	}

	return opt.TypeValue.Str
}

func (c *Type) GetBool(name string) bool {
	var (
		opt *Option
		err error
	)

	if opt, err = c.Find(name, OPTION_TYPE_BOOL); err != nil {
		return false
	}

	return opt.TypeValue.Bool
}

func (c *Type) GetURL(name string) *url.URL {
	var (
		opt *Option
		err error
	)

	if opt, err = c.Find(name, OPTION_TYPE_URL); err != nil {
		return nil
	}

	return opt.TypeValue.URL
}
