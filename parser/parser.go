// Package parser defines main Parser interface and basic functionality to parse a
package parser

import (
	"fmt"

	"github.com/pijalu/go-config/source"
)

// Parser defines a basic parser
type Parser interface {
	// Parse parses a given data into a changeset. It returns error in case of error
	Parse(source string, data interface{}) (*source.ChangeSet, error)
}

// parsers is the map of configured parsers
var parsers = map[string]Parser{}

// Add adds a parser for a given source
func Add(source string, parser Parser) error {
	if _, present := parsers[source]; present {
		return fmt.Errorf("a parser is already defined for %s",
			source)
	}
	parsers[source] = parser

	return nil
}

// Remove removes a parser for a given source
func Remove(source string) {
	delete(parsers, source)
}

// Reset resets all parsers. This method should only be used for testing
func Reset() {
	parsers = map[string]Parser{}
}

// Parse parses a given source data
func Parse(source string, data interface{}) (*source.ChangeSet, error) {
	parser, ok := parsers[source]
	if !ok {
		return nil, fmt.Errorf("no parser configured for %s", source)
	}
	return parser.Parse(source, data)
}
