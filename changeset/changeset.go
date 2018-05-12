package changeset

import "time"

// ChangeSet represents a set of changes
type ChangeSet struct {
	Data      map[string]interface{}
	Checksum  string
	Timestamp time.Time
	Source    string
}
