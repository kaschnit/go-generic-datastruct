package set_test

import (
	"fmt"
	"testing"

	"github.com/kaschnit/go-ds/pkg/containers/set"
	"github.com/kaschnit/go-ds/pkg/containers/set/hashset"
	"github.com/stretchr/testify/assert"
)

func getSetsForTest[T comparable](values ...T) []set.Set[T] {
	return []set.Set[T]{
		hashset.New(values...),
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
			sets := getSetsForTest(testCase.initial...)
			for _, s := range sets {
				t.Run(fmt.Sprintf("%T", s), func(t *testing.T) {
					assert.False(t, s.Empty())
				})
			}
		})
	}
}

func TestEmptyTrue(t *testing.T) {
	sets := getSetsForTest[int]()
	for _, s := range sets {
		assert.True(t, s.Empty())

		s.Add(1)
		assert.False(t, s.Empty())

		s.Remove(1)
		assert.True(t, s.Empty())
	}
}
