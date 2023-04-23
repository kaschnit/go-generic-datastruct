package maputil_test

import (
	"testing"

	"github.com/kaschnit/go-ds/pkg/maputil"
	"github.com/stretchr/testify/assert"
)

func TestForEach(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		mapping  map[string]int
		expected int
	}{
		{
			name:     "do nothing",
			mapping:  map[string]int{},
			expected: 0,
		},
		{
			name:     "add a single number to its key's length",
			mapping:  map[string]int{"foo": 12},
			expected: 15,
		},
		{
			name: "add a few numbers to their keys' lengths",
			mapping: map[string]int{
				"a":     -100,
				"hello": 300,
				"":      57,
			},
			expected: 263,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			total := 0
			maputil.ForEach(testCase.mapping, func(key string, value int) {
				total += len(key) + value
			})
			assert.Equal(t, testCase.expected, total)
		})
	}
}

func TestFilter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		mapping  map[string]int
		expected map[string]int
	}{
		{
			name: "negative at index 0",
			mapping: map[string]int{
				"a": -100,
				"b": 300,
				"c": 57,
			},
			expected: map[string]int{"a": -100},
		},
		{
			name: "negative at index 1",
			mapping: map[string]int{
				"a": 100,
				"b": -300,
				"c": 57,
			},
			expected: map[string]int{"b": -300},
		},
		{
			name: "negative at index 2",
			mapping: map[string]int{
				"a": 100,
				"b": 300,
				"c": -57,
			},
			expected: map[string]int{"c": -57},
		},
		{
			name: "negative at index 2 and 3",
			mapping: map[string]int{
				"abc": 100,
				"def": -300,
				"ghi": -57,
			},
			expected: map[string]int{
				"def": -300,
				"ghi": -57,
			},
		},
		{
			name: "negative at index 1 and 3",
			mapping: map[string]int{
				"bbb":  -100,
				"aaa":  300,
				"abab": -100,
			},
			expected: map[string]int{
				"bbb":  -100,
				"abab": -100,
			},
		},
		{
			name: "all negatives",
			mapping: map[string]int{
				"a":  -100,
				"ee": -300,
				"f":  -57,
			},
			expected: map[string]int{
				"a":  -100,
				"ee": -300,
				"f":  -57,
			},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := maputil.Filter(testCase.mapping, func(_ string, value int) bool {
				return value < 0
			})
			assert.Equal(t, testCase.expected, result)
		})
	}
}

func TestAnyAll(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		mapping     map[string]int
		expectedAny bool
		expectedAll bool
	}{
		{
			name:        "no values",
			mapping:     map[string]int{},
			expectedAny: false,
			expectedAll: true,
		},
		{
			name:        "no matches with 1 item",
			mapping:     map[string]int{"abc": 12},
			expectedAny: false,
			expectedAll: false,
		},
		{
			name:        "no matches with 1 item",
			mapping:     map[string]int{"a": -12},
			expectedAny: false,
			expectedAll: false,
		},
		{
			name:        "match with 1 item",
			mapping:     map[string]int{"hello": -12},
			expectedAny: true,
			expectedAll: true,
		},
		{
			name: "matches at first key",
			mapping: map[string]int{
				"foo": -100,
				"bar": 300,
				"baz": 57,
			},
			expectedAny: true,
			expectedAll: false,
		},
		{
			name: "matches at second key",
			mapping: map[string]int{
				"foo":  100,
				"bar":  -300,
				"aaaa": 57,
			},
			expectedAny: true,
			expectedAll: false,
		},
		{
			name: "no matches values with 3 items",
			mapping: map[string]int{
				"foo": 100,
				"aa":  300,
				"ddd": 57,
			},
			expectedAny: false,
			expectedAll: false,
		},
		{
			name: "all matches",
			mapping: map[string]int{
				"aa": -100,
				"bb": -300,
				"ff": -57,
			},
			expectedAny: true,
			expectedAll: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			isMatch := func(key string, value int) bool {
				return len(key) > 1 && value < 0
			}
			t.Run("Any", func(t *testing.T) {
				assert.Equal(t, testCase.expectedAny, maputil.Any(testCase.mapping, isMatch))
			})
			t.Run("All", func(t *testing.T) {
				assert.Equal(t, testCase.expectedAll, maputil.All(testCase.mapping, isMatch))
			})
		})
	}
}

func TestFindOk(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		mapping       map[string]int
		expectedKey   string
		expectedValue int
	}{
		{
			name: "negative at index 0",
			mapping: map[string]int{
				"foo": -100,
				"bar": 300,
				"baz": 57,
			},
			expectedKey:   "foo",
			expectedValue: -100,
		},
		{
			name: "negative at index 1",
			mapping: map[string]int{
				"a": 100,
				"b": -300,
				"c": 57,
			},
			expectedKey:   "b",
			expectedValue: -300,
		},
		{
			name: "negative at index 2",
			mapping: map[string]int{
				"x": 100,
				"y": 300,
				"z": -57,
			},
			expectedKey:   "z",
			expectedValue: -57,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			key, val, ok := maputil.Find(testCase.mapping, func(_ string, value int) bool {
				return value < 0
			})
			assert.True(t, ok)
			assert.Equal(t, testCase.expectedKey, key)
			assert.Equal(t, testCase.expectedValue, val)
		})
	}
}

func TestFindNotOk(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		mapping map[string]int
	}{
		{
			name:    "no values",
			mapping: map[string]int{},
		},
		{
			name:    "no negative values with 1 item",
			mapping: map[string]int{"abc": 12},
		},
		{
			name: "no negatives with 3 items",
			mapping: map[string]int{
				"a":   100,
				"foo": 300,
				"bar": 57,
			},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			_, _, ok := maputil.Find(testCase.mapping, func(_ string, value int) bool {
				return value < 0
			})
			assert.False(t, ok)
		})
	}
}

func TestMap(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		mapping  map[string]string
		expected []int
	}{
		{
			name:     "no items",
			mapping:  map[string]string{},
			expected: []int{},
		},
		{
			name:     "one items",
			mapping:  map[string]string{"123456789": "f"},
			expected: []int{9},
		},
		{
			name: "a few items",
			mapping: map[string]string{
				"54321":        "aaaa",
				"":             "eeee",
				"222444666888": "",
			},
			expected: []int{5, 0, 12},
		},
		{
			name: "a few more items",
			mapping: map[string]string{
				"a":       "eeee",
				"foo":     "",
				"bar":     "",
				"abcdefg": "p",
			},
			expected: []int{1, 3, 3, 7},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := maputil.Map(testCase.mapping, func(key string, _ string) int {
				return len(key)
			})
			assert.ElementsMatch(t, testCase.expected, result)
		})
	}
}
