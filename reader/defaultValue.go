package reader

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/pijalu/go-config/changeset"
)

type defaultValues struct {
	cs *changeset.ChangeSet
}

type defaultValue struct {
	v interface{}
}

// NewValues builds a new value
func NewValues(s *changeset.ChangeSet) Values {
	return &defaultValues{
		cs: s,
	}
}

// Map returns values as a map[string]interface{}
func (d *defaultValues) Map() map[string]interface{} {
	return d.cs.Data
}

// Get return a value pointed by given path
func (d *defaultValues) Get(path ...string) Value {
	v, present := getPath(d.cs.Data, path...)
	if !present {
		return newValue(nil)
	}
	return newValue(v)
}

// newValue builds a new Value
func newValue(v interface{}) Value {
	return &defaultValue{
		v: v,
	}
}

// asString convert an interface to a string
func asString(v interface{}) (string, bool) {
	s, ok := v.(string)
	if !ok {
		b, ok := v.(fmt.Stringer)
		if !ok {
			return "", false
		}
		s = b.String()
	}
	return s, true
}

// Bool return boolean value
func (v *defaultValue) Bool(def bool) bool {
	if v.v == nil {
		return def
	}

	// Try first a bool cast
	b, ok := v.v.(bool)
	if !ok {
		s, ok := asString(v.v)
		if !ok {
			return def
		}
		// parse
		b, err := strconv.ParseBool(s)
		if err != nil {
			return def
		}
		return b
	}
	return b
}

// Int returns int value
func (v *defaultValue) Int(def int) int {
	if v.v == nil {
		return def
	}

	b, ok := v.v.(int)
	if !ok {
		s, ok := asString(v.v)
		if !ok {
			return def
		}
		b, err := strconv.ParseInt(s, 0, 0)
		if err != nil {
			return def
		}
		return int(b)
	}
	return b
}

// String returns the value as a string, or def if it cannot be converted
func (v *defaultValue) String(def string) string {
	if v.v == nil {
		return def
	}

	b, ok := asString(v.v)
	if !ok {
		return def
	}
	return b
}

// Float64 returns the value as a float, or def if it cannot be converted
func (v *defaultValue) Float64(def float64) float64 {
	if v.v == nil {
		return def
	}
	b, ok := v.v.(float64)
	if !ok {
		s, ok := asString(v.v)
		if !ok {
			return def
		}
		var err error
		b, err = strconv.ParseFloat(s, 64)
		if err != nil {
			return def
		}
	}
	return b
}

// Duration returns the value as a  duration, or def if it cannot be converted
func (v *defaultValue) Duration(def time.Duration) time.Duration {
	if v.v == nil {
		return def
	}
	s, ok := asString(v.v)
	if !ok {
		return def
	}
	d, err := time.ParseDuration(s)
	if err != nil {
		return def
	}
	return d
}

// StringSlice returns value as a string slice, or def if it cannot be converted
func (v *defaultValue) StringSlice(def []string) []string {
	if v.v == nil {
		return def
	}
	s, ok := v.v.([]string)
	if !ok {
		return def
	}
	return s
}

// StringMap returns value as a map of string of def if it cannot be converted
func (v *defaultValue) StringMap(def map[string]string) map[string]string {
	if v.v == nil {
		return def
	}
	m, ok := v.v.(map[string]interface{})
	if !ok {
		return def
	}
	r := make(map[string]string)
	for k, v := range flatMap(m) {
		var ok bool
		if r[k], ok = asString(v); !ok {
			r[k] = fmt.Sprintf("%v", v)
		}
	}
	return r
}

// Scan fills a struct with given value
func (v *defaultValue) Scan(val interface{}) error {
	if v.v == nil {
		return errors.New("value is nil")
	}
	m, ok := v.v.(map[string]interface{})
	if !ok {
		return errors.New("value is not a map")
	}

	return mapstructure.Decode(m, val)
}
