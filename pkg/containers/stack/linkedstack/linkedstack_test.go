package linkedstack_test

import (
	"testing"

	"github.com/kaschnit/go-ds/pkg/containers/stack"
	"github.com/kaschnit/go-ds/pkg/containers/stack/linkedstack"
	"github.com/stretchr/testify/assert"
)

// Ensure that LinkedStack implements Stack
var _ stack.Stack[int] = &linkedstack.LinkedStack[int]{}

func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		stack    *linkedstack.LinkedStack[int]
		expected string
	}{
		{
			name:     "empty stack",
			stack:    linkedstack.New[int](),
			expected: "LinkedStack\n",
		},
		{
			name:     "stack with 1 item",
			stack:    linkedstack.New(987654321),
			expected: "LinkedStack\n987654321",
		},
		{
			name:     "stack with a few items",
			stack:    linkedstack.New(100, 1145, -202, 5, 6, 7),
			expected: "LinkedStack\n100,1145,-202,5,6,7",
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, testCase.stack.String())
		})
	}
}
