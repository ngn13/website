package util

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func Render(file string, data interface{}) ([]byte, error) {
	var (
		rendered *bytes.Buffer
		tmpl     *template.Template
		content  []byte
		err      error
	)

	if content, err = os.ReadFile(file); err != nil {
		return nil, err
	}

	if tmpl, err = template.New("template").Parse(string(content)); err != nil {
		return nil, err
	}

	rendered = bytes.NewBuffer(nil)
	err = tmpl.Execute(rendered, data)

	return rendered.Bytes(), err
}

func GetDuration(d string) (time.Duration, error) {
	var (
		d_num uint64
		err   error
	)

	d_num_end := d[len(d)-1]
	d_num_str := strings.TrimSuffix(d, string(d_num_end))

	if d_num, err = strconv.ParseUint(d_num_str, 10, 64); err != nil {
		return 0, err
	}

	switch d_num_end {
	case 's':
		return time.Duration(d_num) * (time.Second), nil

	case 'm':
		return time.Duration(d_num) * (time.Second * 60), nil

	case 'h':
		return time.Duration(d_num) * ((time.Second * 60) * 60), nil
	}

	return 0, fmt.Errorf("invalid time duration format")
}

func GetSHA1(s string) string {
	hasher := sha1.New()
	return hex.EncodeToString(hasher.Sum([]byte(s)))
}
