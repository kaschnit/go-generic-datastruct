package hashset_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/kaschnit/go-ds/pkg/containers/set"
	"github.com/kaschnit/go-ds/pkg/containers/set/hashset"
	"github.com/stretchr/testify/assert"
)

// Ensure that HashSet implements Set
var _ set.Set[int] = hashset.New(1)

func TestHashSetString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		set  *hashset.HashSet[int]
	}{
		{
			name: "empty set",
			set:  hashset.New[int](),
		},
		{
			name: "set with 1 item",
			set:  hashset.New(987654321),
		},
		{
			name: "set with a few items",
			set:  hashset.New(100, 1145, -202, 5, 6, 7),
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			resultLines := strings.Split(testCase.set.String(), "\n")
			assert.Len(t, resultLines, 2, "expected 2 lines in HashSet.String() output")
			assert.Equal(t, resultLines[0], "HashSet")

			// Set does not guarantee ordering
			testCase.set.ForEach(func(_, value int) {
				assert.Contains(t, resultLines[1], fmt.Sprintf("%d", value))
			})
		})
	}
}
