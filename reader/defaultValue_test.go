package reader

import (
	"fmt"
	"testing"
)

type structWithStringer struct{}

func (t structWithStringer) String() string {
	return "ts"
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

func TestBool(t *testing.T) {
	type testCase struct {
		v        interface{}
		def      bool
		expected bool
	}

	for _, c := range []testCase{
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
