package enumerable_test

import (
	"fmt"
	"testing"

	"github.com/kaschnit/go-ds/pkg/containers/enumerable"
	"github.com/kaschnit/go-ds/pkg/containers/list/arraylist"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	testCases := []struct {
		name     string
		items    []string
		expected []int
	}{
		{
			name:     "no items",
			items:    []string{},
			expected: []int{},
		},
		{
			name:     "one items",
			items:    []string{"123456789"},
			expected: []int{9},
		},
		{
			name:     "a few items",
			items:    []string{"54321", "", "222444666888"},
			expected: []int{5, 0, 12},
		},
		{
			name:     "a few more items",
			items:    []string{"a", "foo", "bar", "abcdefg"},
			expected: []int{1, 3, 3, 7},
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("ArrayList %s", testCase.name), func(t *testing.T) {
			l := arraylist.New(testCase.items...)
			result := enumerable.Map[int, string](l, func(_ int, v string) int {
				return len(v)
			})
			assert.Equal(t, testCase.expected, result)
		})
	}
}
