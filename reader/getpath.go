package reader

func getPath(m map[string]interface{}, paths ...string) (value interface{}, present bool) {
	if len(paths) == 0 || len(m) == 0 {
		return nil, false
	}

	// Loop througt all the key except last one
	it := m
	var ok bool
	for _, path := range paths[:len(paths)-1] {
		it, ok = it[path].(map[string]interface{})
		if !ok {
			return nil, false
		}
	}

	// return last interface value
	value, present = it[paths[len(paths)-1]]
	return
}
