package envvar

import (
	"fmt"
	"os"
	"strings"

	"github.com/imdario/mergo"
	"github.com/pijalu/go-config/changeset"
	"github.com/pijalu/go-config/parser"
	"github.com/pijalu/go-config/parser/noop"
	"github.com/pijalu/go-config/source"
)

const sourceName = "envvar"

var (
	DefaultPrefixes = []string{}
)

type envvar struct {
	prefixes         []string
	strippedPrefixes []string
	opts             source.Options
}

var prs parser.Parser

func init() {
	prs = noop.NewParser()
}

func (e *envvar) Load() (interface{}, error) {
	var changes map[string]interface{}

	for _, env := range os.Environ() {

		if len(e.prefixes) > 0 || len(e.strippedPrefixes) > 0 {
			notFound := true

			if _, ok := matchPrefix(e.prefixes, env); ok {
				notFound = false
			}

			if match, ok := matchPrefix(e.strippedPrefixes, env); ok {
				env = strings.TrimPrefix(env, match)
				notFound = false
			}

			if notFound {
				continue
			}
		}

		pair := strings.SplitN(env, "=", 2)
		value := pair[1]
		keys := strings.Split(strings.ToLower(pair[0]), "_")
		reverse(keys)

		tmp := make(map[string]interface{})
		for i, k := range keys {
			if i == 0 {
				tmp[k] = value
				continue
			}

			tmp = map[string]interface{}{k: tmp}
		}

		if err := mergo.Map(&changes, tmp); err != nil {
			return nil, err
		}
	}
	return changes, nil
}

func (e *envvar) Read() (*changeset.ChangeSet, error) {
	data, err := e.Load()
	if err != nil {
		return nil, fmt.Errorf("Failed to read: %v", err)
	}
	return prs.Parse(sourceName, data)
}

func matchPrefix(pre []string, s string) (string, bool) {
	for _, p := range pre {
		if strings.HasPrefix(s, p) {
			return p, true
		}
	}

	return "", false
}

func reverse(ss []string) {
	for i := len(ss)/2 - 1; i >= 0; i-- {
		opp := len(ss) - 1 - i
		ss[i], ss[opp] = ss[opp], ss[i]
	}
}

func (e *envvar) Watch() (source.Watcher, error) {
	return newWatcher()
}

func (e *envvar) String() string {
	return sourceName
}

// NewSource returns a config source for parsing ENV variables.
// Underscores are delimiters for nesting, and all keys are lowercased.
//
// Example:
//      "DATABASE_SERVER_HOST=localhost" will convert to
//
//      {
//          "database": {
//              "server": {
//                  "host": "localhost"
//              }
//          }
//      }
func NewSource(opts ...source.Option) source.Source {
	var options source.Options
	for _, o := range opts {
		o(&options)
	}

	var sp []string
	var pre []string
	if options.Context != nil {
		if p, ok := options.Context.Value(strippedPrefixKey{}).([]string); ok {
			sp = p
		}

		if p, ok := options.Context.Value(prefixKey{}).([]string); ok {
			pre = p
		}

		if len(sp) > 0 || len(pre) > 0 {
			pre = append(pre, DefaultPrefixes...)
		}
	}
	return &envvar{prefixes: pre, strippedPrefixes: sp, opts: options}
}
