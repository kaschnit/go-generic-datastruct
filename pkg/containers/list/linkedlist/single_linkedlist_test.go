package linkedlist_test

import (
	"testing"

	"github.com/kaschnit/go-ds/pkg/containers/list"
	"github.com/kaschnit/go-ds/pkg/containers/list/linkedlist"
	"github.com/stretchr/testify/assert"
)

// Ensure that SingleLinkedList implements List
var _ list.List[int] = &linkedlist.SingleLinkedList[int]{}

func TestSingleLinkedString(t *testing.T) {
	tests := []struct {
		name     string
		list     *linkedlist.SingleLinkedList[int]
		expected string
	}{
		{
			name:     "empty list",
			list:     linkedlist.NewSingleLinked[int](),
			expected: "SingleLinkedList\n",
		},
		{
			name:     "list with 1 item",
			list:     linkedlist.NewSingleLinked(987654321),
			expected: "SingleLinkedList\n987654321",
		},
		{
			name:     "list with a few items",
			list:     linkedlist.NewSingleLinked(100, 1145, -202, 5, 6, 7),
			expected: "SingleLinkedList\n100,1145,-202,5,6,7",
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, testCase.list.String())
		})
	}
}
