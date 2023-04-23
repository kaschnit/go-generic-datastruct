package sliceutil_test

import (
	"testing"

	"github.com/kaschnit/go-ds/pkg/sliceutil"
	"github.com/stretchr/testify/assert"
)

func TestForEach(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		slice    []int
		expected int
	}{
		{
			name:     "sum nothing",
			slice:    []int{},
			expected: 0,
		},
		{
			name:     "sum a single number",
			slice:    []int{12},
			expected: 12,
		},
		{
			name:     "sum a few numbers",
			slice:    []int{-100, 300, 57},
			expected: 257,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			total := 0
			sliceutil.ForEach(testCase.slice, func(_ int, value int) {
				total += value
			})
			assert.Equal(t, testCase.expected, total)
		})
	}
}

func TestFilter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		slice    []int
		expected []int
	}{
		{
			name:     "negative at index 0",
			slice:    []int{-100, 300, 57},
			expected: []int{-100},
		},
		{
			name:     "negative at index 1",
			slice:    []int{100, -300, 57},
			expected: []int{-300},
		},
		{
			name:     "negative at index 2",
			slice:    []int{100, 300, -57},
			expected: []int{-57},
		},
		{
			name:     "negative at index 2 and 3",
			slice:    []int{100, -300, -57},
			expected: []int{-300, -57},
		},
		{
			name:     "negative at index 1 and 3",
			slice:    []int{-100, 300, -100},
			expected: []int{-100, -100},
		},
		{
			name:     "all negatives",
			slice:    []int{-100, -300, -57},
			expected: []int{-100, -300, -57},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := sliceutil.Filter(testCase.slice, func(_ int, value int) bool {
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
		slice       []int
		expectedAny bool
		expectedAll bool
	}{
		{
			name:        "no values",
			slice:       []int{},
			expectedAny: false,
			expectedAll: true,
		},
		{
			name:        "no negative values with 1 item",
			slice:       []int{12},
			expectedAny: false,
			expectedAll: false,
		},
		{
			name:        "negative at index 0",
			slice:       []int{-100, 300, 57},
			expectedAny: true,
			expectedAll: false,
		},
		{
			name:        "negative at index 1",
			slice:       []int{100, -300, 57},
			expectedAny: true,
			expectedAll: false,
		},
		{
			name:        "negative at index 2",
			slice:       []int{100, 300, -57},
			expectedAny: true,
			expectedAll: false,
		},
		{
			name:        "no negative values with 3 items",
			slice:       []int{100, 300, 57},
			expectedAny: false,
			expectedAll: false,
		},
		{
			name:        "all negatives",
			slice:       []int{-100, -300, -57},
			expectedAny: true,
			expectedAll: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			isNegative := func(_ int, value int) bool {
				return value < 0
			}
			t.Run("Any", func(t *testing.T) {
				assert.Equal(t, testCase.expectedAny, sliceutil.Any(testCase.slice, isNegative))
			})
			t.Run("All", func(t *testing.T) {
				assert.Equal(t, testCase.expectedAll, sliceutil.All(testCase.slice, isNegative))
			})
		})
	}
}

func TestFindOk(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		slice         []int
		expectedIndex int
		expectedValue int
	}{
		{
			name:          "negative at index 0",
			slice:         []int{-100, 300, 57},
			expectedIndex: 0,
			expectedValue: -100,
		},
		{
			name:          "negative at index 1",
			slice:         []int{100, -300, 57},
			expectedIndex: 1,
			expectedValue: -300,
		},
		{
			name:          "negative at index 2",
			slice:         []int{100, 300, -57},
			expectedIndex: 2,
			expectedValue: -57,
		},
		{
			name:          "negative at index 2 and 3",
			slice:         []int{100, -300, -57},
			expectedIndex: 1,
			expectedValue: -300,
		},
		{
			name:          "negative at index 1 and 3",
			slice:         []int{-100, 300, -100},
			expectedIndex: 0,
			expectedValue: -100,
		},
		{
			name:          "all negatives",
			slice:         []int{-100, -300, -57},
			expectedIndex: 0,
			expectedValue: -100,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			idx, val, ok := sliceutil.Find(testCase.slice, func(_ int, value int) bool {
				return value < 0
			})
			assert.True(t, ok)
			assert.Equal(t, testCase.expectedIndex, idx)
			assert.Equal(t, testCase.expectedValue, val)
		})
	}
}

func TestFindNotOk(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		slice []int
	}{
		{
			name:  "no values",
			slice: []int{},
		},
		{
			name:  "no negative values with 1 item",
			slice: []int{12},
		},
		{
			name:  "no negatives with 3 items",
			slice: []int{100, 300, 57},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			_, _, ok := sliceutil.Find(testCase.slice, func(_ int, value int) bool {
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
		slice    []string
		expected []int
	}{
		{
			name:     "no items",
			slice:    []string{},
			expected: []int{},
		},
		{
			name:     "one items",
			slice:    []string{"123456789"},
			expected: []int{9},
		},
		{
			name:     "a few items",
			slice:    []string{"54321", "", "222444666888"},
			expected: []int{5, 0, 12},
		},
		{
			name:     "a few more items",
			slice:    []string{"a", "foo", "bar", "abcdefg"},
			expected: []int{1, 3, 3, 7},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := sliceutil.Map(testCase.slice, func(_ int, v string) int {
				return len(v)
			})
			assert.Equal(t, testCase.expected, result)
		})
	}
}

func TestMapMap(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		slice    []string
		expected map[string]int
	}{
		{
			name:     "no items",
			slice:    []string{},
			expected: map[string]int{},
		},
		{
			name:     "one item",
			slice:    []string{"123456789"},
			expected: map[string]int{"123456789": 9},
		},
		{
			name:  "a few items",
			slice: []string{"54321", "", "222444666888"},
			expected: map[string]int{
				"54321":        5,
				"":             0,
				"222444666888": 12,
			},
		},
		{
			name:  "a few more items",
			slice: []string{"a", "foo", "bar", "abcdefg"},
			expected: map[string]int{
				"a":       1,
				"foo":     3,
				"bar":     3,
				"abcdefg": 7,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := sliceutil.MapMap(testCase.slice, func(_ int, v string) (string, int) {
				return v, len(v)
			})
			assert.Equal(t, testCase.expected, result)
		})
	}
}
