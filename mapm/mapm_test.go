package mapm

import (
	"reflect"
	"testing"
)

type testCase struct {
	a        map[string]interface{}
	b        map[string]interface{}
	expected map[string]interface{}
	error    bool
}

func TestMerge(t *testing.T) {
	cases := []testCase{
		// Simple case
		{
			a: map[string]interface{}{
				"a": "b",
			},
			b: map[string]interface{}{
				"a": "c",
			},
			expected: map[string]interface{}{
				"a": "c",
			},
			error: false,
		},
		// Array
		{
			a: map[string]interface{}{
				"a": []string{"1", "2", "3"},
			},
			b: map[string]interface{}{
				"a": []string{"4", "5", "6"},
			},
			expected: map[string]interface{}{
				"a": []string{"1", "2", "3", "4", "5", "6"},
			},
			error: false,
		},
		// Multi layer
		{
			a: map[string]interface{}{
				"k1": "v1",
				"k2": map[string]interface{}{
					"k2.1": "v2",
				},
			},
			b: map[string]interface{}{
				"k1": "v1bis",
				"k2": map[string]interface{}{
					"k2.2": "v3",
				},
			},
			expected: map[string]interface{}{
				"k1": "v1bis",
				"k2": map[string]interface{}{
					"k2.1": "v2",
					"k2.2": "v3",
				},
			},
			error: false,
		},
	}

	for idx, tc := range cases {
		actual, err := Merge(tc.a, tc.b)
		if tc.error && err == nil {
			t.Errorf("expected failed merge but was successful for test case %d: %v",
				idx,
				tc)
			continue
		}
		if !reflect.DeepEqual(tc.expected, actual) {
			t.Errorf("expected %v but was %v for test case %d",
				tc.expected,
				tc.a,
				idx)
			continue
		}
	}

}
