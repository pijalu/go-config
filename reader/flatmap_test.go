package reader

import (
	"reflect"
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
	if !reflect.DeepEqual(expectedMap, actual) {
		t.Fatalf("Expected %v but  got %v",
			expectedMap,
			actual)
	}
}
