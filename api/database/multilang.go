package database

import (
	"encoding/json"
	"reflect"
	"strings"
	"unicode"
)

type Multilang struct {
	En string `json:"en"` // english
	Tr string `json:"tr"` // turkish
}

func (ml *Multilang) Supports(lang string) bool {
	ml_ref := reflect.ValueOf(ml).Elem()

	for i := 0; i < reflect.Indirect(ml_ref).NumField(); i++ {
		if name := reflect.Indirect(ml_ref).Type().Field(i).Name; strings.ToLower(name) == lang {
			return true
		}
	}

	return false
}

func (ml *Multilang) Get(lang string) string {
	r := []rune(lang)
	r[0] = unicode.ToUpper(r[0])
	l := string(r)

	ml_ref := reflect.ValueOf(ml)
	return reflect.Indirect(ml_ref).FieldByName(l).String()
}

func (ml *Multilang) Empty() bool {
	ml_ref := reflect.ValueOf(ml)

	for i := 0; i < reflect.Indirect(ml_ref).NumField(); i++ {
		if field := reflect.Indirect(ml_ref).Field(i); field.String() != "" {
			return false
		}
	}

	return true
}

func (ml *Multilang) Dump() (string, error) {
	if data, err := json.Marshal(ml); err != nil {
		return "", err
	} else {
		return string(data), nil
	}
}

func (ml *Multilang) Load(s string) error {
	return json.Unmarshal([]byte(s), ml)
}
