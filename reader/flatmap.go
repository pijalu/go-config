package reader

import "fmt"

// flatten map with building merged keys. Duplicated entry will be lost
func flatMap(m map[string]interface{}) map[string]interface{} {
	rm := make(map[string]interface{})

	for k, v := range m {
		// we have a submap
		if subMap, ok := m[k].(map[string]interface{}); ok {
			for sk, sv := range flatMap(subMap) {
				rm[fmt.Sprintf("%s.%s", k, sk)] = sv
			}
		} else {
			rm[k] = v
		}
	}
	return rm
}
