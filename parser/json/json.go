// Package json implement a json parser
package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/pijalu/go-config/parser/noop"
	"github.com/pijalu/go-config/source"
)

// Parser object
type Parser struct{}

// Parse a json data struct
func (m *Parser) parseJSON(data interface{}) (map[string]interface{}, error) {
	var bytes []byte

	// Load data (support string and array of bytes
	switch reflect.TypeOf(data).Kind() {
	case reflect.String:
		bytes = []byte(data.(string))
	default:
		var ok bool
		if bytes, ok = data.([]byte); !ok {
			return nil, errors.New("data is not a byte array")
		}
	}

	jsonData := make(map[string]interface{})
	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		return nil, fmt.Errorf("could not parse json: %v", err)
	}

	return jsonData, nil
}

// Parse parses data, expected to be []bytes and return a matching changeset
func (m *Parser) Parse(src string, data interface{}) (*source.ChangeSet, error) {
	jsonData, err := m.parseJSON(data)
	if err != nil {
		return nil, err
	}

	// Let noop parser package the map
	return (&noop.Parser{}).Parse(src, jsonData)
}
