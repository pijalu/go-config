// Package file is a file source. Expected format is json
package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pijalu/go-config/parser"
	"github.com/pijalu/go-config/parser/json"
	"github.com/pijalu/go-config/parser/xml"
	"github.com/pijalu/go-config/parser/yaml"
	"github.com/pijalu/go-config/source"
)

type file struct {
	path    string
	opts    source.Options
	modTime time.Time
}

var (
	DefaultPath = "config.json"
)

var parsers = make(map[string]parser.Parser)

func init() {
	parsers["json"] = json.NewParser()
	parsers["yaml"] = yaml.NewParser()
	parsers["xml"] = xml.NewParser()
}

func (f *file) Load() (interface{}, error) {
	fh, err := os.Open(f.path)
	if err != nil {
		return nil, err
	}
	defer fh.Close()
	b, err := ioutil.ReadAll(fh)
	if err != nil {
		return nil, err
	}
	info, err := fh.Stat()
	if err != nil {
		return nil, err
	}
	f.modTime = info.ModTime()

	return b, nil
}

func (f *file) Read() (*source.ChangeSet, error) {
	data, err := f.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to read: %v", err)
	}
	ext := strings.ToLower(
		strings.TrimPrefix(filepath.Ext(f.path), "."))
	p, present := parsers[ext]
	if !present {
		return nil, fmt.Errorf("could not find parser for %s (ext %s)", f.path, ext)
	}

	cs, err := p.Parse("file/"+ext, data)
	// Update mod time based on file
	if err != nil {
		cs.Timestamp = f.modTime
		cs.RecalculateChecksum()
	}

	return cs, err
}

func (f *file) String() string {
	return "file"
}

func (f *file) Watch() (source.Watcher, error) {
	if _, err := os.Stat(f.path); err != nil {
		return nil, err
	}
	return newWatcher(f)
}

func NewSource(opts ...source.Option) source.Source {
	var options source.Options
	for _, o := range opts {
		o(&options)
	}
	path := DefaultPath
	if options.Context != nil {
		f, ok := options.Context.Value(filePathKey{}).(string)
		if ok {
			path = f
		}
	}
	return &file{opts: options, path: path}
}
