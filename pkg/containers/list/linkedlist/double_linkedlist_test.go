package linkedlist_test

import (
	"testing"

	"github.com/kaschnit/go-ds/pkg/containers/list"
	"github.com/kaschnit/go-ds/pkg/containers/list/linkedlist"
	"github.com/stretchr/testify/assert"
)

// Ensure that DoubleLinkedList implements List
var _ list.List[int] = &linkedlist.DoubleLinkedList[int]{}

func TestDoubleLinkedString(t *testing.T) {
	tests := []struct {
		name     string
		list     *linkedlist.DoubleLinkedList[int]
		expected string
	}{
		{
			name:     "empty list",
			list:     linkedlist.NewDoubleLinked[int](),
			expected: "DoubleLinkedList\n",
		},
		{
			name:     "list with 1 item",
			list:     linkedlist.NewDoubleLinked(987654321),
			expected: "DoubleLinkedList\n987654321",
		},
		{
			name:     "list with a few items",
			list:     linkedlist.NewDoubleLinked(100, 1145, -202, 5, 6, 7),
			expected: "DoubleLinkedList\n100,1145,-202,5,6,7",
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, testCase.list.String())
		})
	}
}
