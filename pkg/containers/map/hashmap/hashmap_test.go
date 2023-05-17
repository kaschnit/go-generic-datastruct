package hashmap_test

import (
	"strings"
	"testing"

	"github.com/kaschnit/go-ds/pkg/compare"
	mapp "github.com/kaschnit/go-ds/pkg/containers/map"
	"github.com/kaschnit/go-ds/pkg/containers/map/entry"
	"github.com/kaschnit/go-ds/pkg/containers/map/hashmap"
	"github.com/stretchr/testify/assert"
)

// Ensure that HashMap implements Map.
var _ mapp.Map[string, int] = &hashmap.HashMap[string, string, int]{}

func TestBuilderPut(t *testing.T) {
	t.Parallel()

	m := hashmap.NewBuilder[int, int, string](compare.IdentityHashKey[int]).
		Put(1, "a").
		Put(2, "c").
		Put(3, "e").
		Build()
	assert.Equal(t, 3, m.Size())

	val, ok := m.Get(2)
	assert.True(t, ok)
	assert.Equal(t, "c", val)

	val, ok = m.Get(1)
	assert.True(t, ok)
	assert.Equal(t, "a", val)

	val, ok = m.Get(3)
	assert.True(t, ok)
	assert.Equal(t, "e", val)

	_, ok = m.Get(12345)
	assert.False(t, ok)
}

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
	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			resultLines := strings.Split(testCase.mapping.String(), "\n")
			assert.Len(t, resultLines, 2, "expected 2 lines in HashMap.String() output")
			assert.Equal(t, "HashMap", resultLines[0])

			// Map does not guarantee ordering
			testCase.mapping.ForEach(func(key int, value string) {
				assert.Contains(t, resultLines[1], entry.NewRef(key, value).String())
			})
		})
	}
}

func TestNewFromMap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		mapping          map[string]int
		expectedContains []string
	}{
		{
			name:             "empty map",
			mapping:          map[string]int{},
			expectedContains: []string{},
		},
		{
			name: "non-empty map",
			mapping: map[string]int{
				"foo": 48765,
				"bar": 1,
			},
			expectedContains: []string{"foo", "bar"},
		},
	}

	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			m := hashmap.NewFromMap(testCase.mapping)
			assert.Equal(t, len(testCase.expectedContains), m.Size())
			assert.True(t, m.ContainsAllKeys(testCase.expectedContains...))
		})
	}
}
