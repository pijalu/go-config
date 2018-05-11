package reader

import (
	"testing"
)

func TestFlatMap(t *testing.T) {
	inputMap := map[string]interface{}{
		"key1": map[string]interface{}{
			"key1_1": "value1_1",
			"key1_2": "value1_2",
		},
		"key2": "value2",
	}

	expectedMap := map[string]interface{}{
		"key1.key1_1": "value1_1",
		"key1.key1_2": "value1_2",
		"key2":        "value2",
	}

	actual := flatMap(inputMap)
	if len(expectedMap) != len(actual) {
		t.Fatalf("Expected %d elems but got %d",
			len(expectedMap),
			len(actual))
	}

	for k, v := range expectedMap {
		if actual[k] != v {
			t.Fatalf("expected %v for key %v but got %v",
				v,
				k,
				actual[k])
		}
	}
}
