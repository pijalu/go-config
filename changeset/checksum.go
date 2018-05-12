package changeset

import (
	"crypto/md5"
	"encoding/gob"
	"fmt"
)

// RecalculateChecksum recalculate checkum
// It returns the updated changeset
func (c *ChangeSet) RecalculateChecksum() *ChangeSet {
	h := md5.New()

	enc := gob.NewEncoder(h)
	enc.Encode(c.Data)

	c.Checksum = fmt.Sprintf("%x", h.Sum(nil))

	return c
}
