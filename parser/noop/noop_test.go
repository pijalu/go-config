package noop

import (
	"reflect"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	startTime := time.Now()

	input := make(map[string]interface{})
	input["hello"] = "world"

	m := Parser{}
	actual, err := m.Parse("test", input)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(input, actual.Data) {
		t.Fatalf("Data is not as expected: want %v but have %v",
			input,
			actual)
	}
	if actual.Checksum == "" {
		t.Fatalf("Checksum is not set")
	}
	if actual.Source != "test" {
		t.Fatalf("Unexpected source: want %s but have %v",
			"test",
			actual.Source)
	}
	if !startTime.Before(actual.Timestamp) {
		t.Fatalf("Unexpected timestamp: want it to be after %v but it's %v",
			startTime,
			actual.Timestamp)
	}
}
