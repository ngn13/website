package config

import (
	"fmt"
	"net/url"
	"strings"
)

const (
	OPTION_TYPE_STR  = 0
	OPTION_TYPE_BOOL = 1
	OPTION_TYPE_URL  = 2
)

type Option struct {
	Name      string
	Value     string
	Required  bool
	Type      uint8
	TypeValue struct {
		URL  *url.URL
		Str  string
		Bool bool
	}
}

func (o *Option) Env() string {
	return strings.ToUpper(fmt.Sprintf("WEBSITE_%s", o.Name))
}

func (o *Option) Load() (err error) {
	err = nil

	switch o.Type {
	case OPTION_TYPE_STR:
		o.TypeValue.Str = o.Value

	case OPTION_TYPE_BOOL:
		o.TypeValue.Bool = "1" == o.Value || "true" == strings.ToLower(o.Value)

	case OPTION_TYPE_URL:
		o.TypeValue.URL, err = url.Parse(o.Value)

	default:
		return fmt.Errorf("invalid option type")
	}

	return err
}
