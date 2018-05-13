// Package mapm allows a basic merge of 2 map[string]interface{}
package mapm

import (
	"fmt"
	"reflect"
)

// Merge merges 2 maps and return the merge results.
// Please note that trg map will be changed during merge
func Merge(trg map[string]interface{}, src map[string]interface{}) (map[string]interface{}, error) {
	// Create map if needed
	if trg == nil && src != nil {
		trg = make(map[string]interface{})
	}

	for k, v := range src {
		item, present := trg[k]
		// if not present, no merge
		if !present {
			trg[k] = v
			continue
		}

		// They are different type, so ignore merge
		if reflect.TypeOf(item) != reflect.TypeOf(v) {
			continue
		}

		// Let's look deeper
		switch reflect.TypeOf(v).Kind() {
		// Append to slice
		case reflect.Slice:
			reflect.ValueOf(trg).SetMapIndex(
				reflect.ValueOf(k),
				reflect.AppendSlice(
					reflect.ValueOf(item),
					reflect.ValueOf(v)))

			// Merge map
		case reflect.Map:
			itemMap, ok := item.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("could not convert item %v (%T) to map",
					item,
					item)
			}
			sourceMap, ok := v.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("could not convert item %v (%T) to map",
					v,
					v)
			}
			var err error
			if trg[k], err = Merge(itemMap, sourceMap); err != nil {
				return nil, fmt.Errorf("failed to merge key %v: %v",
					k,
					err)
			}
			// Replace
		default:
			trg[k] = v
		}
	}
	return trg, nil
}
