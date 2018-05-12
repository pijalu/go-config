package microcli

import (
	"fmt"
	"strings"

	"github.com/imdario/mergo"
	"github.com/micro/cli"
	"github.com/pijalu/go-config/changeset"
	"github.com/pijalu/go-config/parser"
	"github.com/pijalu/go-config/parser/noop"
	"github.com/pijalu/go-config/source"
)

type clisrc struct {
	opts source.Options
	ctx  *cli.Context
}

const sourceName = "microcli"

var prs parser.Parser

func init() {
	prs = noop.NewParser()
}

func (c *clisrc) Load() (interface{}, error) {
	var changes map[string]interface{}

	for _, name := range c.ctx.GlobalFlagNames() {
		tmp := toEntry(name, c.ctx.GlobalGeneric(name))
		mergo.Map(&changes, tmp) // need to sort error handling
	}

	for _, name := range c.ctx.FlagNames() {
		tmp := toEntry(name, c.ctx.Generic(name))
		mergo.Map(&changes, tmp) // need to sort error handling
	}
	return changes, nil
}

func (c *clisrc) Read() (*changeset.ChangeSet, error) {
	data, err := c.Load()
	if err != nil {
		return nil, fmt.Errorf("Failed to read: %v", err)
	}
	return prs.Parse(sourceName, data)
}

func toEntry(name string, v interface{}) map[string]interface{} {
	n := strings.ToLower(name)
	keys := strings.FieldsFunc(n, split)
	reverse(keys)
	tmp := make(map[string]interface{})
	for i, k := range keys {
		if i == 0 {
			tmp[k] = v
		} else {
			tmp = map[string]interface{}{k: tmp}
		}
	}
	return tmp
}

func reverse(ss []string) {
	for i := len(ss)/2 - 1; i >= 0; i-- {
		opp := len(ss) - 1 - i
		ss[i], ss[opp] = ss[opp], ss[i]
	}
}

func split(r rune) bool {
    return r == '-' || r == '_'
}

func (c *clisrc) Watch() (source.Watcher, error) {
	return source.NewNoopWatcher()
}

func (c *clisrc) String() string {
	return sourceName
}

// NewSource returns a config source for integrating parsed flags from a micro/cli.Context.
// Hyphens are delimiters for nesting, and all keys are lowercased.
//
// Example:
//      cli.StringFlag{Name: "db-host"},
//
//
//      {
//          "database": {
//              "host": "localhost"
//          }
//      }
func NewSource(ctx *cli.Context, opts ...source.Option) source.Source {
	var options source.Options
	for _, o := range opts {
		o(&options)
	}

	return &clisrc{opts: options, ctx: ctx}
}
