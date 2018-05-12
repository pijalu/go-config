// Package noop implements a noparse parser.
// This "parser" is intended for source already providing "map" data
package noop

import (
	"errors"
	"time"

	"github.com/pijalu/go-config/changeset"
	"github.com/pijalu/go-config/parser"
)

// Parser is a map parser
type Parser struct{}

// NewParser a new Noop parser
func NewParser() parser.Parser {
	return &Parser{}
}

// Parse parses data, expected to be map[string]interfaces and return a matching changeset
func (m *Parser) Parse(src string, data interface{}) (*changeset.ChangeSet, error) {
	parsedMap, ok := data.(map[string]interface{})
	if !ok {
		return nil, errors.New("data is not a map")
	}

	// Create ChangeSet
	return (&changeset.ChangeSet{
		Data:      parsedMap,
		Timestamp: time.Now(),
		Source:    src,
	}).RecalculateChecksum(), nil
}
