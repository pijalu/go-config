// Package memory is a memory source
package memory

import (
	"sync"
	"time"

	"github.com/pborman/uuid"
	"github.com/pijalu/go-config/source"
)

type memory struct {
	sync.RWMutex
	ChangeSet *source.ChangeSet
	Watchers  map[string]*watcher
}

func (s *memory) Load() (interface{}, error) {
	return s.ChangeSet.Data, nil
}

func (s *memory) Read() (*source.ChangeSet, error) {
	s.RLock()
	cs := &source.ChangeSet{
		Timestamp: s.ChangeSet.Timestamp,
		Data:      s.ChangeSet.Data,
		Checksum:  s.ChangeSet.Checksum,
		Source:    s.ChangeSet.Source,
	}
	s.RUnlock()
	return cs, nil
}

func (s *memory) Watch() (source.Watcher, error) {
	w := &watcher{
		Id:      uuid.NewUUID().String(),
		Updates: make(chan *source.ChangeSet, 100),
		Source:  s,
	}

	s.Lock()
	s.Watchers[w.Id] = w
	s.Unlock()
	return w, nil
}

// Update allows manual updates of the config data.
func (s *memory) Update(data map[string]interface{}) {
	s.Lock()
	// update changeset
	s.ChangeSet = (&source.ChangeSet{
		Timestamp: time.Now(),
		Data:      data,
		Source:    "memory",
	}).RecalculateChecksum()

	// update watchers
	for _, w := range s.Watchers {
		select {
		case w.Updates <- s.ChangeSet:
		default:
		}
	}
	s.Unlock()
}

func (s *memory) String() string {
	return "memory"
}

func NewSource(opts ...source.Option) source.Source {
	var options source.Options
	for _, o := range opts {
		o(&options)
	}

	var data map[string]interface{}

	if options.Context != nil {
		d, ok := options.Context.Value(dataKey{}).(map[string]interface{})
		if ok {
			data = d
		}
	}

	s := &memory{
		Watchers: make(map[string]*watcher),
	}
	s.Update(data)
	return s
}
