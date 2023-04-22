package stack_test

import (
	"fmt"
	"testing"

	"github.com/kaschnit/go-ds/pkg/containers/stack"
	"github.com/kaschnit/go-ds/pkg/containers/stack/arraystack"
	"github.com/kaschnit/go-ds/pkg/containers/stack/linkedstack"
	"github.com/stretchr/testify/assert"
)

func getStacksForTest[T any](values ...T) []stack.Stack[T] {
	return []stack.Stack[T]{
		arraystack.New(values...),
		linkedstack.New(values...),
	}
}

func TestEmptyFalse(t *testing.T) {
	tests := []struct {
		name    string
		initial []string
	}{
		{
			name:    "3 items",
			initial: []string{"a", "b", "c"},
		},
		{
			name:    "1 item",
			initial: []string{"hello"},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			stacks := getStacksForTest(testCase.initial...)
			for _, s := range stacks {
				t.Run(fmt.Sprintf("%T", s), func(t *testing.T) {
					assert.False(t, s.Empty())
				})
			}
		})
	}
}

func TestEmptyTrue(t *testing.T) {
	stacks := getStacksForTest[int]()
	for _, s := range stacks {
		t.Run(fmt.Sprintf("%T", s), func(t *testing.T) {
			assert.True(t, s.Empty())

			s.Push(1)
			assert.False(t, s.Empty())

			s.Pop()
			assert.True(t, s.Empty())
		})
	}
}

func TestSize(t *testing.T) {
	tests := []struct {
		name     string
		initial  []float64
		expected int
	}{
		{
			name:     "3 items",
			initial:  []float64{1.2, 2.3, 999.999},
			expected: 3,
		},
		{
			name:     "1 item",
			initial:  []float64{7.000},
			expected: 1,
		},
		{
			name:     "0 items",
			initial:  []float64{},
			expected: 0,
		},
		{
			name:     "6 items",
			initial:  []float64{2.5, 1.000, -5.444, 0.1, 500, 12},
			expected: 6,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			stacks := getStacksForTest(testCase.initial...)
			for _, s := range stacks {
				t.Run(fmt.Sprintf("%T", s), func(t *testing.T) {
					assert.Equal(t, testCase.expected, s.Size())
					assert.Equal(t, len(testCase.initial), s.Size())
				})
			}
		})
	}
}

func TestClearNonEmpty(t *testing.T) {
	tests := []struct {
		name    string
		initial []float64
	}{
		{
			name:    "3 items",
			initial: []float64{1.2, 2.3, 999.999},
		},
		{
			name:    "1 item",
			initial: []float64{7.000},
		},
		{
			name:    "6 items",
			initial: []float64{2.5, 1.000, -5.444, 0.1, 500, 12},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			stacks := getStacksForTest(testCase.initial...)
			for _, s := range stacks {
				t.Run(fmt.Sprintf("%T", s), func(t *testing.T) {
					assert.False(t, s.Empty())

					s.Clear()
					assert.True(t, s.Empty())

					s.Clear()
					assert.True(t, s.Empty())

					s.Push(1.2345)
					assert.False(t, s.Empty())

					s.Clear()
					assert.True(t, s.Empty())
				})
			}
		})
	}
}

func TestClearEmpty(t *testing.T) {
	stacks := getStacksForTest[string]()
	for _, s := range stacks {
		assert.True(t, s.Empty())

		s.Clear()
		assert.True(t, s.Empty())

		s.Clear()
		assert.True(t, s.Empty())

		s.Push("hello")
		assert.False(t, s.Empty())

		s.Clear()
		assert.True(t, s.Empty())
	}
}

func TestPush(t *testing.T) {
	tests := []struct {
		name     string
		initial  []string
		pushItem string
	}{
		{
			name:     "3 items pushing a string",
			initial:  []string{"1", "two", "3"},
			pushItem: "foo",
		},
		{
			name:     "3 items pushing an empty string",
			initial:  []string{"1", "two", "3"},
			pushItem: "",
		},
		{
			name:     "0 items pushing a string",
			initial:  []string{},
			pushItem: "hello",
		},
		{
			name:     "0 items pushing an empty string",
			initial:  []string{},
			pushItem: "",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			stacks := getStacksForTest(testCase.initial...)
			for _, s := range stacks {
				t.Run(fmt.Sprintf("%T", s), func(t *testing.T) {
					prevSize := s.Size()
					s.Push(testCase.pushItem)
					assert.Equal(t, prevSize+1, s.Size())

					actual, ok := s.Peek()
					assert.True(t, ok)
					assert.Equal(t, testCase.pushItem, actual)
				})
			}
		})
	}
}

func TestPushAll(t *testing.T) {
	tests := []struct {
		name         string
		initial      []int
		pushItems    []int
		expectedPeek int
	}{
		{
			name:         "pushing no items",
			initial:      []int{9, 8, 7},
			pushItems:    []int{},
			expectedPeek: 7,
		},
		{
			name:         "pushing one item",
			initial:      []int{9, 8, 7},
			pushItems:    []int{1},
			expectedPeek: 1,
		},
		{
			name:         "pushing some items",
			initial:      []int{9, 8, 7},
			pushItems:    []int{500, 1000},
			expectedPeek: 1000,
		},
		{
			name:         "pushing one item onto an empty stack",
			initial:      []int{},
			pushItems:    []int{-52},
			expectedPeek: -52,
		},
		{
			name:         "pushing some items onto an empty stack",
			initial:      []int{},
			pushItems:    []int{1, 2, 3},
			expectedPeek: 3,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			stacks := getStacksForTest(testCase.initial...)
			for _, s := range stacks {
				t.Run(fmt.Sprintf("%T", s), func(t *testing.T) {
					prevSize := s.Size()
					s.PushAll(testCase.pushItems...)
					assert.Equal(t, prevSize+len(testCase.pushItems), s.Size())

					actual, ok := s.Peek()
					assert.True(t, ok)
					assert.Equal(t, testCase.expectedPeek, actual)
				})
			}
		})
	}
}

func TestPopUntilEmpty(t *testing.T) {
	vals := []int{1, 2, 3}
	stacks := getStacksForTest(vals...)
	for _, s := range stacks {
		val, ok := s.Pop()
		assert.True(t, ok)
		assert.Equal(t, 3, val)

		val, ok = s.Pop()
		assert.True(t, ok)
		assert.Equal(t, 2, val)

		val, ok = s.Pop()
		assert.True(t, ok)
		assert.Equal(t, 1, val)

		val, ok = s.Pop()
		assert.False(t, ok, "Expected not ok but was ok, val was %v", val)
	}
}
