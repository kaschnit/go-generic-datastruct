package compare_test

import (
	"testing"

	"github.com/kaschnit/go-ds/pkg/compare"
	"github.com/stretchr/testify/assert"
)

// Ensure that OrderedComparator implements Comparator
var _ compare.Comparator[int] = compare.OrderedComparator[int]

func TestOrderedComparator(t *testing.T) {
	tests := []struct {
		name     string
		left     int
		right    int
		expected compare.Priority
	}{
		{
			name:     "left less than right",
			left:     0,
			right:    1,
			expected: compare.PriorityRightHigher,
		},
		{
			name:     "left way less than right",
			left:     7,
			right:    9000,
			expected: compare.PriorityRightHigher,
		},
		{
			name:     "left and right equal",
			left:     100,
			right:    100,
			expected: compare.PriorityEqual,
		},
		{
			name:     "left and right equal",
			left:     -73,
			right:    -73,
			expected: compare.PriorityEqual,
		},
		{
			name:     "left greater than right",
			left:     5,
			right:    4,
			expected: compare.PriorityLeftHigher,
		},
		{
			name:     "left way greater than right",
			left:     9000,
			right:    7,
			expected: compare.PriorityLeftHigher,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, compare.OrderedComparator(testCase.left, testCase.right))
		})
	}
}
