package yaml

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/pijalu/go-config/changeset"
	"github.com/pijalu/go-config/parser"
	"github.com/pijalu/go-config/parser/noop"

	"gopkg.in/yaml.v2"
)

// Parser object
type Parser struct{}

// NewParser return a json parser
func NewParser() parser.Parser {
	return &Parser{}
}

func stringifyKeys(m map[interface{}]interface{}) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	for k, v := range m {
		ks, ok := k.(string)
		if !ok {
			kss, ok := k.(fmt.Stringer)
			if !ok {
				return nil, fmt.Errorf("cannot stringify key %s (typ %T)",
					k,
					k)
			}
			ks = kss.String()
		}
		subMap, ok := v.(map[interface{}]interface{})
		if ok {
			subMapWithStringKey, err := stringifyKeys(subMap)
			if err != nil {
				return nil, err
			}
			out[ks] = subMapWithStringKey
		} else {
			out[ks] = v
		}
	}
	return out, nil
}

// Parse a YAML data struct
func (m *Parser) parseYAML(data interface{}) (map[string]interface{}, error) {
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

	yamlData := make(map[interface{}]interface{})
	if err := yaml.Unmarshal(bytes, &yamlData); err != nil {
		return nil, fmt.Errorf("could not parse yaml: %v", err)
	}

	stringifiedMap, err := stringifyKeys(yamlData)
	if err != nil {
		return nil, fmt.Errorf("could not convert keys of yaml: %v", err)
	}

	return stringifiedMap, nil
}

// Parse parses data, expected to be []bytes and return a matching changeset
func (m *Parser) Parse(src string, data interface{}) (*changeset.ChangeSet, error) {
	yamlData, err := m.parseYAML(data)
	if err != nil {
		return nil, err
	}

	// Let noop parser package the map
	return (&noop.Parser{}).Parse(src, yamlData)
}
