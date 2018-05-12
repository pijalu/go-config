package reader

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/pijalu/go-config/changeset"
)

type structWithStringer struct{}

func (t structWithStringer) String() string {
	return "ts"
}

func TestMap(t *testing.T) {
	expected := map[string]interface{}{"key": "value"}
	actual := NewValues(&changeset.ChangeSet{
		Data: expected,
	}).Map()

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %v but got %v",
			expected,
			actual)
	}
}

func TestGet(t *testing.T) {
	expected := "value"

	vs := NewValues(&changeset.ChangeSet{
		Data: map[string]interface{}{
			"k1": map[string]interface{}{
				"k2": expected,
			},
		},
	})

	actual := vs.Get("k1", "k2").String("nope")
	if actual != expected {
		t.Fatalf("expected %v but got %v",
			expected,
			actual)
	}
}

func TestAsString(t *testing.T) {
	type testCase struct {
		v        interface{}
		ok       bool
		expected string
	}

	for _, c := range []testCase{
		// Test string
		{
			v:        "hello",
			ok:       true,
			expected: "hello",
		},
		// int should not be stringified
		{
			v:  1,
			ok: false,
		},
		// check a struct that implement fmt.Stringer interface
		{
			v:        structWithStringer{},
			ok:       true,
			expected: "ts",
		},
	} {
		fmt.Printf("")
		actual, ok := asString(c.v)
		if ok != c.ok {
			t.Fatalf("expected %v for convert but was %v (case %v)",
				c.ok,
				ok,
				c)
		}
		if ok && actual != c.expected {
			t.Fatalf("expected %v but got %v (case %v)",
				c.expected,
				actual,
				c)
		}
	}
}

func TestString(t *testing.T) {
	type testCase struct {
		v        interface{}
		def      string
		expected string
	}

	for _, c := range []testCase{
		// Test nil
		{
			v:        nil,
			def:      "default",
			expected: "default",
		},
		// Test string
		{
			v:        "hello",
			def:      "default",
			expected: "hello",
		},
		// int should not be stringified, so expect default
		{
			v:        1,
			def:      "default",
			expected: "default",
		},
		// check a struct that implement fmt.Stringer interface
		{
			v:        structWithStringer{},
			def:      "default",
			expected: "ts",
		},
	} {
		actual := newValue(c.v).String(c.def)
		if actual != c.expected {
			t.Fatalf("expected %v but got %v (case %v)",
				c.expected,
				actual,
				c)
		}
	}
}

func TestBool(t *testing.T) {
	type testCase struct {
		v        interface{}
		def      bool
		expected bool
	}

	for _, c := range []testCase{
		// Check nil
		{
			v:        nil,
			def:      true,
			expected: true,
		},
		// Check convert for actual bool
		{
			v:        true,
			def:      false,
			expected: true,
		},
		{
			v:        false,
			def:      true,
			expected: false,
		},
		// Check convert for stringed boolean
		{
			v:        "true",
			def:      false,
			expected: true,
		},
		{
			v:        "false",
			def:      true,
			expected: false,
		},
		// int should not be boolean
		{
			v:        1,
			def:      false,
			expected: false,
		},
		{
			v:        0,
			def:      true,
			expected: true,
		},
	} {
		actual := newValue(c.v).Bool(c.def)
		if actual != c.expected {
			t.Fatalf("expected %v but got %v with test case %v",
				c.expected,
				actual,
				c)
		}
	}
}

func TestInt(t *testing.T) {
	type testCase struct {
		v        interface{}
		def      int
		expected int
	}

	for _, c := range []testCase{
		// Check nil
		{
			v:        nil,
			def:      0,
			expected: 0,
		},
		// Check convert for actual type
		{
			v:        1,
			def:      0,
			expected: 1,
		},
		// Check convert for stringed type
		{
			v:        "1",
			def:      0,
			expected: 1,
		},
		// string should not work
		{
			v:        "one",
			def:      0,
			expected: 0,
		},
	} {
		actual := newValue(c.v).Int(c.def)
		if actual != c.expected {
			t.Fatalf("expected %v but got %v with test case %v",
				c.expected,
				actual,
				c)
		}
	}
}

func TestFloat64(t *testing.T) {
	type testCase struct {
		v        interface{}
		def      float64
		expected float64
	}

	for _, c := range []testCase{
		// Check nil
		{
			v:        nil,
			def:      0,
			expected: 0,
		},
		// Check convert for actual type
		{
			v:        1.0,
			def:      0,
			expected: 1,
		},
		// Check convert for stringed type
		{
			v:        "1.0",
			def:      0,
			expected: 1.0,
		},
		// string should not work
		{
			v:        "one",
			def:      0,
			expected: 0,
		},
	} {
		actual := newValue(c.v).Float64(c.def)
		if actual != c.expected {
			t.Fatalf("expected %v but got %v with test case %v",
				c.expected,
				actual,
				c)
		}
	}
}

func TestDuration(t *testing.T) {
	type testCase struct {
		v        interface{}
		def      time.Duration
		expected time.Duration
	}

	for _, c := range []testCase{
		// Check nil
		{
			v:        nil,
			def:      time.Minute * 1,
			expected: time.Minute * 1,
		},
		// Check convert for actual type
		{
			v:        time.Hour * 1,
			def:      time.Minute * 1,
			expected: time.Hour * 1,
		},
		// Check convert for stringed type
		{
			v:        "24h",
			def:      time.Minute * 1,
			expected: time.Hour * 24,
		},
		// string should not work
		{
			v:        "one",
			def:      time.Minute * 1,
			expected: time.Minute * 1,
		},
	} {
		actual := newValue(c.v).Duration(c.def)
		if actual != c.expected {
			t.Fatalf("expected %v but got %v with test case %v",
				c.expected,
				actual,
				c)
		}
	}
}

func TestStringSlice(t *testing.T) {
	type testCase struct {
		v        interface{}
		def      []string
		expected []string
	}

	for _, c := range []testCase{
		// Check convert for actual type
		{
			v:        nil,
			def:      []string{"3", "4"},
			expected: []string{"3", "4"},
		},
		// Check convert for actual type
		{
			v:        []string{"1", "2"},
			def:      []string{"3", "4"},
			expected: []string{"1", "2"},
		},
		// Check convert for non convertable item
		{
			v:        "nope",
			def:      []string{"1", "2"},
			expected: []string{"1", "2"},
		},
	} {
		actual := newValue(c.v).StringSlice(c.def)
		if !reflect.DeepEqual(c.expected, actual) {
			t.Fatalf("expected %v but got %v with test case %v",
				c.expected,
				actual,
				c)
		}
	}
}

func TestStringMap(t *testing.T) {
	type testCase struct {
		v        interface{}
		def      map[string]string
		expected map[string]string
	}

	for _, c := range []testCase{
		// Check nil
		{
			v:        nil,
			def:      map[string]string{"aKey": "aValue"},
			expected: map[string]string{"aKey": "aValue"},
		},
		// Check convert for actual type
		{
			v: map[string]interface{}{
				"k1": map[string]interface{}{
					"k2": "value",
				},
			},
			def:      map[string]string{"aKey": "aValue"},
			expected: map[string]string{"k1.k2": "value"},
		},
		// Check convert for a stringer
		{
			v: map[string]interface{}{
				"k1": map[string]interface{}{
					"k2": structWithStringer{},
				},
			},
			def:      map[string]string{"aKey": "aValue"},
			expected: map[string]string{"k1.k2": "ts"},
		},
		// Check convert for a non-string value
		{
			v: map[string]interface{}{
				"k1": map[string]interface{}{
					"k2": 1,
				},
			},
			def:      map[string]string{"aKey": "aValue"},
			expected: map[string]string{"k1.k2": "1"},
		},
	} {
		actual := newValue(c.v).StringMap(c.def)
		if !reflect.DeepEqual(c.expected, actual) {
			t.Fatalf("expected %v but got %v with test case %v",
				c.expected,
				actual,
				c)
		}
	}
}

func TestScan(t *testing.T) {
	type ts struct {
		Question string
		Answer   int
	}

	expected := ts{
		Question: `what is the meaning of life everything and the universe`,
		Answer:   42,
	}

	m := map[string]interface{}{
		"question": expected.Question,
		"answer":   expected.Answer,
		"notused":  69,
	}

	actual := ts{}
	if err := newValue(m).Scan(&actual); err != nil {
		t.Fatalf("failed to scan: %v", err)
	}
	if actual != expected {
		t.Fatalf("expected %v but got %v",
			expected,
			actual)
	}
}

func TestScanEmpty(t *testing.T) {
	type ts struct {
		Question string
		Answer   int
	}

	expected := ts{}
	m := map[string]interface{}{}

	actual := ts{}
	if err := newValue(m).Scan(&actual); err != nil {
		t.Fatalf("failed to scan: %v", err)
	}
	if actual != expected {
		t.Fatalf("expected %v but got %v",
			expected,
			actual)
	}
}

func TestScanAsNotAMap(t *testing.T) {
	type ts struct {
		Question string
		Answer   int
	}

	if err := newValue(1).Scan(&ts{}); err == nil {
		t.Fatal("Expected scan to fail but it didn't")
	}
}
