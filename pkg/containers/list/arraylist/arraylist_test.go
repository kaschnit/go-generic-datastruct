package arraylist_test

import (
	"testing"

	"github.com/kaschnit/go-ds/pkg/containers/iterator"
	"github.com/kaschnit/go-ds/pkg/containers/list"
	"github.com/kaschnit/go-ds/pkg/containers/list/arraylist"
	"github.com/stretchr/testify/assert"
)

// Ensure that ArrayList implements List
var _ list.List[int] = &arraylist.ArrayList[int]{}

// Ensure that ArrayList implements ReverseIterable
var _ iterator.ReverseIterable[int, string] = &arraylist.ArrayList[string]{}

func TestString(t *testing.T) {
	tests := []struct {
		name     string
		list     *arraylist.ArrayList[int]
		expected string
	}{
		{
			name:     "empty list",
			list:     arraylist.New[int](),
			expected: "ArrayList\n",
		},
		{
			name:     "list with 1 item",
			list:     arraylist.New(987654321),
			expected: "ArrayList\n987654321",
		},
		{
			name:     "list with a few items",
			list:     arraylist.New(100, 1145, -202, 5, 6, 7),
			expected: "ArrayList\n100,1145,-202,5,6,7",
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, testCase.list.String())
		})
	}
}
