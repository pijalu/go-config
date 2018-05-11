package reader

import "testing"

func TestGetPath(t *testing.T) {
	m := map[string]interface{}{
		"key1": map[string]interface{}{
			"key1_1": "value1_1",
			"key1_2": "value1_2",
		},
		"key2": "value2",
	}

	type testCase struct {
		path     []string
		expected string
		present  bool
	}

	testCases := []testCase{
		{
			path:     []string{"key1", "key1_1"},
			expected: "value1_1",
			present:  true,
		},
		{
			path:     []string{"key1", "key1_2"},
			expected: "value1_2",
			present:  true,
		},
		{
			path:     []string{"key2"},
			expected: "value2",
			present:  true,
		},
		{
			path:    []string{"key1", "key1_1", "key_nope"},
			present: false,
		},
	}

	for _, tc := range testCases {
		actual, present := getPath(m, tc.path...)
		if present != tc.present {
			t.Fatalf("expected present %v but was %v for test case %v",
				tc.present,
				present,
				tc)
		}
		if present && actual != tc.expected {
			t.Fatalf("expected %s but got %v for test case %v",
				tc.expected,
				actual,
				tc)
		}
	}

}
