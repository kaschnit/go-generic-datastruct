package concurrentmap_test

import (
	"fmt"
	"strings"
	"testing"

	mapp "github.com/kaschnit/go-ds/pkg/containers/map"
	"github.com/kaschnit/go-ds/pkg/containers/map/concurrentmap"
	"github.com/kaschnit/go-ds/pkg/containers/map/entry"
	"github.com/kaschnit/go-ds/pkg/containers/map/hashmap"
	"github.com/stretchr/testify/assert"
)

func TestConcurrentMapString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		mapping           mapp.Map[string, int]
		expectedFirstLine string
	}{
		{
			name:              "empty hashmap",
			mapping:           hashmap.New[string, int](),
			expectedFirstLine: "HashMap",
		},
		{
			name:              "hashmap with 1 item",
			mapping:           hashmap.New(entry.New("foo", 987654321)),
			expectedFirstLine: "HashMap",
		},
		{
			name: "hashmap with a few items",
			mapping: hashmap.New(
				entry.New("abc", 100),
				entry.New("def", 1145),
				entry.New("ghi", -202),
				entry.New("jkl", 5),
				entry.New("mno", 6),
				entry.New("pqr", 7),
			),
			expectedFirstLine: "HashMap",
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			m := concurrentmap.MakeThreadSafe(testCase.mapping)
			resultLines := strings.Split(m.String(), "\n")
			assert.Len(t, resultLines, 2, "expected 2 lines in HashMap.String() output")
			assert.Equal(t, fmt.Sprintf("[Concurrent]%s", testCase.expectedFirstLine), resultLines[0])

			// Map does not guarantee ordering
			testCase.mapping.ForEach(func(key string, value int) {
				assert.Contains(t, resultLines[1], entry.NewRef(key, value).String())
			})
		})
	}
}
