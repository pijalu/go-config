package consul

import (
	"errors"
	"log"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/watch"
	"github.com/pijalu/go-config/source"
)

type watcher struct {
	name        string
	stripPrefix string

	wp   *watch.Plan
	ch   chan *source.ChangeSet
	exit chan bool
}

func newWatcher(key, addr, name string, stripPrefix string) (source.Watcher, error) {
	w := &watcher{
		name:        name,
		stripPrefix: stripPrefix,
		ch:          make(chan *source.ChangeSet),
		exit:        make(chan bool),
	}

	wp, err := watch.Parse(map[string]interface{}{"type": "keyprefix", "prefix": key})
	if err != nil {
		return nil, err
	}

	wp.Handler = w.handle

	// wp.Run is a blocking call and will prevent newWatcher from returning
	go wp.Run(addr)

	w.wp = wp

	return w, nil
}

func (w *watcher) handle(idx uint64, data interface{}) {
	if data == nil {
		return
	}

	kvs, ok := data.(api.KVPairs)
	if !ok {
		return
	}

	cs, err := prs.Parse(sourceName, makeMap(kvs, w.stripPrefix))
	if err != nil {
		log.Printf("error during parse: %v", err)
	}

	w.ch <- cs
}

func (w *watcher) Next() (*source.ChangeSet, error) {
	select {
	case cs := <-w.ch:
		return cs, nil
	case <-w.exit:
		return nil, errors.New("watcher stopped")
	}
}

func (w *watcher) Stop() error {
	select {
	case <-w.exit:
		return nil
	default:
		w.wp.Stop()
		close(w.exit)
	}
	return nil
}
