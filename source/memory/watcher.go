package memory

import "github.com/pijalu/go-config/changeset"

type watcher struct {
	Id      string
	Updates chan *changeset.ChangeSet
	Source  *memory
}

func (w *watcher) Next() (*changeset.ChangeSet, error) {
	cs := <-w.Updates
	return cs, nil
}

func (w *watcher) Stop() error {
	w.Source.Lock()
	delete(w.Source.Watchers, w.Id)
	w.Source.Unlock()
	return nil
}
