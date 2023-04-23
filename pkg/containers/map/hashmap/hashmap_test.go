package hashmap_test

import (
	"strings"
	"testing"

	mapp "github.com/kaschnit/go-ds/pkg/containers/map"
	"github.com/kaschnit/go-ds/pkg/containers/map/entry"
	"github.com/kaschnit/go-ds/pkg/containers/map/hashmap"
	"github.com/stretchr/testify/assert"
)

// Ensure that HashMap implements Map
var _ mapp.Map[string, int] = &hashmap.HashMap[string, string, int]{}

func TestHashMapString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		mapping *hashmap.HashMap[int, int, string]
	}{
		{
			name:    "empty hashmap",
			mapping: hashmap.New[int, string](),
		},
		{
			name:    "hashmap with 1 item",
			mapping: hashmap.New(entry.New(987654321, "foo")),
		},
		{
			name: "hashmap with a few items",
			mapping: hashmap.New(
				entry.New(100, "abc"),
				entry.New(1145, "abc"),
				entry.New(-202, "abc"),
				entry.New(5, "abc"),
				entry.New(6, "abc"),
				entry.New(7, "abc"),
			),
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			resultLines := strings.Split(testCase.mapping.String(), "\n")
			assert.Len(t, resultLines, 2, "expected 2 lines in HashMap.String() output")
			assert.Equal(t, resultLines[0], "HashMap")

			// Map does not guarantee ordering
			testCase.mapping.ForEach(func(key int, value string) {
				assert.Contains(t, resultLines[1], entry.NewRef(key, value).String())
			})
		})
	}
}
