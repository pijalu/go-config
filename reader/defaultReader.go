package reader

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/pijalu/go-config/source"
	"github.com/pijalu/mergo"
)

// defaultReader is a default reader
type defaultReader struct{}

// NewReader returns a new reader
func NewReader() Reader {
	return &defaultReader{}
}

// String returns reader name
func (d *defaultReader) String() string {
	return "default"
}

// Parse load and merge a series of ChangeSet
func (d *defaultReader) Parse(changeSets ...*source.ChangeSet) (*source.ChangeSet, error) {
	if len(changeSets) == 0 {
		return nil, nil
	}

	merged := make(map[string]interface{})
	sources := make([]string, 0, len(changeSets))

	for _, cs := range changeSets {
		sources = append(sources, cs.Source)
		if err := mergo.Map(&merged, cs.Data); err != nil {
			return nil, fmt.Errorf("failed to merge data: %v", err)
		}
	}

	// Return merged change set
	return (&source.ChangeSet{
		Timestamp: time.Now(),
		Data:      merged,
		Source:    strings.Join(sources, ";"),
	}).RecalculateChecksum(), nil
}

// Values returns value from changeset
func (d *defaultReader) Values(c *source.ChangeSet) (Values, error) {
	if c == nil {
		return nil, errors.New("changeset is nil")
	}
	// TODO
	return nil, nil
}
