package parser

import (
	"errors"
	"reflect"
	"testing"

	"github.com/pijalu/go-config/changeset"
)

func TestReset(t *testing.T) {
	Add("test", &MockParser{})
	Reset()
	if len(parsers) != 0 {
		t.Fatalf("We don't have expected count of parsers: expected 0, got %d",
			len(parsers))
	}
}

func TestAdd(t *testing.T) {
	Reset()
	Add("test", &MockParser{})
	if len(parsers) != 1 {
		t.Fatalf("We don't have expected count of parsers: expected 1, got %d",
			len(parsers))
	}
	if _, ok := parsers["test"]; !ok {
		t.Fatalf("Got the wrong parser !")
	}
}

func TestRemove(t *testing.T) {
	Reset()
	Add("test", &MockParser{})
	Add("other", &MockParser{})

	Remove("test")
	if len(parsers) != 1 {
		t.Fatalf("We don't have expected count of parsers: expected 0, got %d",
			len(parsers))
	}
	if _, ok := parsers["other"]; !ok {
		t.Fatalf("Didn't get expected parser !")
	}
}

func TestParse(t *testing.T) {
	Reset()
	parser := &MockParser{
		RError:     nil,
		RChangeSet: &changeset.ChangeSet{Source: "test"},
	}
	Add("test", parser)

	cs, err := Parse("test", "Some data")
	// make sure our mock isn't broken...
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if cs.Source != "test" {
		t.Fatalf("Unexpected source: %s", cs.Source)
	}

	expectedCall := []MockParserCall{{
		Source: "test",
		Data:   "Some data",
	}}
	if !reflect.DeepEqual(parser.ParseCalls, expectedCall) {
		t.Fatalf("Call was not valid: expected %v but got %v",
			expectedCall,
			parser.ParseCalls)
	}
}

func TestParseWithError(t *testing.T) {
	Reset()
	expectedErr := errors.New("test")
	parser := &MockParser{
		RError:     expectedErr,
		RChangeSet: nil,
	}
	Add("test", parser)

	if _, err := Parse("test", "Some data"); err != expectedErr {
		t.Fatalf("expected error %v but got %v", expectedErr, err)
	}
}

func TestParseWhenUnknowSource(t *testing.T) {
	Reset()

	if _, err := Parse("test", "Some data"); err == nil {
		t.Fatalf("expected error got none")
	}
}
