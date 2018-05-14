// Package source is the interface for sources
package source

import (
	"context"
)

// Source is the source from which config is loaded
type Source interface {
	Load() (interface{}, error)
	Read() (*ChangeSet, error)
	Watch() (Watcher, error)
	String() string
}

// Watcher watches a source for changes
type Watcher interface {
	Next() (*ChangeSet, error)
	Stop() error
}

type Options struct {
	// for alternative data
	Context context.Context
}

type Option func(o *Options)
