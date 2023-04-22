package arraystack_test

import (
	"testing"

	"github.com/kaschnit/go-ds/pkg/containers/stack"
	"github.com/kaschnit/go-ds/pkg/containers/stack/arraystack"
	"github.com/stretchr/testify/assert"
)

// Ensure that ArrayStack implements Stack
var _ stack.Stack[int] = &arraystack.ArrayStack[int]{}

func TestString(t *testing.T) {
	tests := []struct {
		name     string
		stack    *arraystack.ArrayStack[int]
		expected string
	}{
		{
			name:     "empty stack",
			stack:    arraystack.New[int](),
			expected: "ArrayStack\n",
		},
		{
			name:     "stack with 1 item",
			stack:    arraystack.New(987654321),
			expected: "ArrayStack\n987654321",
		},
		{
			name:     "stack with a few items",
			stack:    arraystack.New(100, 1145, -202, 5, 6, 7),
			expected: "ArrayStack\n100,1145,-202,5,6,7",
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, testCase.stack.String())
		})
	}
}
