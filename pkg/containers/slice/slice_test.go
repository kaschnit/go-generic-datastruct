package slice_test

import (
	"testing"

	"github.com/kaschnit/go-ds/pkg/containers/enumerable"
	"github.com/kaschnit/go-ds/pkg/containers/slice"
	"github.com/stretchr/testify/assert"
)

var _ enumerable.Enumerable[int, string] = slice.Slice[string]([]string{})

func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		slice    slice.Slice[int]
		expected string
	}{
		{
			name:     "empty list",
			slice:    []int{},
			expected: "Slice\n",
		},
		{
			name:     "list with 1 item",
			slice:    []int{987654321},
			expected: "Slice\n987654321",
		},
		{
			name:     "list with a few items",
			slice:    []int{100, 1145, -202, 5, 6, 7},
			expected: "Slice\n100,1145,-202,5,6,7",
		},
	}
	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, testCase.expected, testCase.slice.String())
		})
	}
}

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
	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			total := 0
			slice.Slice[int](testCase.slice).ForEach(func(_ int, value int) {
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
		slice    slice.Slice[int]
		expected slice.Slice[int]
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
	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			result := testCase.slice.Filter(func(_ int, value int) bool {
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

	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			isNegative := func(_ int, value int) bool {
				return value < 0
			}
			t.Run("Any", func(t *testing.T) {
				assert.Equal(t, testCase.expectedAny, slice.Slice[int](testCase.slice).Any(isNegative))
			})
			t.Run("All", func(t *testing.T) {
				assert.Equal(t, testCase.expectedAll, slice.Slice[int](testCase.slice).All(isNegative))
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
	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			idx, val, ok := slice.Slice[int](testCase.slice).Find(func(_ int, value int) bool {
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
	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			_, _, ok := slice.Slice[int](testCase.slice).Find(func(_ int, value int) bool {
				return value < 0
			})
			assert.False(t, ok)
		})
	}
}
