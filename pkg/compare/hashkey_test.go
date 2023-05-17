package compare_test

import (
	"testing"

	"github.com/kaschnit/go-ds/pkg/compare"
	"github.com/stretchr/testify/assert"
)

// Ensure that OrderedComparator implements Comparator.
var (
	_ compare.HashKey[int, int]       = compare.IdentityHashKey[int]
	_ compare.HashKey[string, string] = compare.IdentityHashKey[string]
)

func TestIdentityHashKey(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    any
		expected any
	}{
		{
			name:     "numeric input 1",
			input:    1,
			expected: 1,
		},
		{
			name:     "numeric input 2",
			input:    -8000,
			expected: -8000,
		},
		{
			name:     "string input 1",
			input:    "",
			expected: "",
		},
		{
			name:     "string input 2",
			input:    "foo",
			expected: "foo",
		},
	}
	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, testCase.expected, compare.IdentityHashKey(testCase.input))
		})
	}
}
