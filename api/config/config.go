package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/ngn13/website/api/util"
)

type Option struct {
	Name     string
	Value    string
	Required bool
}

func (o *Option) Env() string {
	return strings.ToUpper(fmt.Sprintf("API_%s", o.Name))
}

var options []Option = []Option{
	{Name: "password", Value: "", Required: true},
	{Name: "frontend_url", Value: "http://localhost:5173/", Required: true},
}

func Load() bool {
	var val string

	for i := range options {
		if val = os.Getenv(options[i].Env()); val == "" {
			continue
		}

		options[i].Value = val
		options[i].Required = false
	}

	for i := range options {
		if options[i].Required && options[i].Value == "" {
			util.Fail("please specify the required config option \"%s\" (\"%s\")", options[i].Name, options[i].Env())
			return false
		}

		if options[i].Required && options[i].Value != "" {
			util.Fail("using the default value \"%s\" for required config option \"%s\" (\"%s\")", options[i].Value, options[i].Name, options[i].Env())
		}
	}

	return true
}

func Get(name string) string {
	for i := range options {
		if options[i].Name != name {
			continue
		}
		return options[i].Value
	}
	return ""
}
