package flag

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/imdario/mergo"
	"github.com/pijalu/go-config/changeset"
	"github.com/pijalu/go-config/parser"
	"github.com/pijalu/go-config/parser/noop"
	"github.com/pijalu/go-config/source"
)

const sourceName = "flag"

type flagsrc struct {
	opts source.Options
}

var prs parser.Parser

func init() {
	prs = noop.NewParser()
}

func (fs *flagsrc) Load() (interface{}, error) {
	if !flag.Parsed() {
		return nil, errors.New("flags not parsed")
	}

	var changes map[string]interface{}
	flag.Visit(func(f *flag.Flag) {
		n := strings.ToLower(f.Name)
		keys := strings.Split(n, "-")
		reverse(keys)

		tmp := make(map[string]interface{})
		for i, k := range keys {
			if i == 0 {
				tmp[k] = f.Value.String()
				continue
			}

			tmp = map[string]interface{}{k: tmp}
		}

		mergo.Map(&changes, tmp) // need to sort error handling
		return
	})

	return changes, nil
}

func (fs *flagsrc) Read() (*changeset.ChangeSet, error) {
	data, err := fs.Load()
	if err != nil {
		return nil, fmt.Errorf("Failed to read: %v", err)
	}
	return prs.Parse(sourceName, data)
}

func reverse(ss []string) {
	for i := len(ss)/2 - 1; i >= 0; i-- {
		opp := len(ss) - 1 - i
		ss[i], ss[opp] = ss[opp], ss[i]
	}
}

func (fs *flagsrc) Watch() (source.Watcher, error) {
	return source.NewNoopWatcher()
}

func (fs *flagsrc) String() string {
	return sourceName
}

// NewSource returns a config source for integrating parsed flags.
// Hyphens are delimiters for nesting, and all keys are lowercased.
//
// Example:
//      dbhost := flag.String("database-host", "localhost", "the db host name")
//
//      {
//          "database": {
//              "host": "localhost"
//          }
//      }
func NewSource(opts ...source.Option) source.Source {
	var options source.Options
	for _, o := range opts {
		o(&options)
	}

	return &flagsrc{opts: options}
}
