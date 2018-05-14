// Package reader parses change sets and provides config values
package reader

import (
	"time"

	"github.com/pijalu/go-config/source"
)

// Reader is an interface for merging changesets
type Reader interface {
	Parse(...*source.ChangeSet) (*source.ChangeSet, error)
	Values(*source.ChangeSet) (Values, error)
	String() string
}

// Values is returned by the reader
type Values interface {
	Get(path ...string) Value
	Map() map[string]interface{}
}

// Value represents a value of any type
type Value interface {
	Checksum() string
	Bool(def bool) bool
	Int(def int) int
	String(def string) string
	Float64(def float64) float64
	Duration(def time.Duration) time.Duration
	StringSlice(def []string) []string
	StringMap(def map[string]string) map[string]string
	Scan(val interface{}) error
}
