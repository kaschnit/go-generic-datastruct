package entry_test

import (
	"testing"

	"github.com/kaschnit/go-ds/pkg/containers/map/entry"
	"github.com/stretchr/testify/assert"
)

func TestEntryString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		value    entry.Entry[any, any]
		expected string
	}{
		{
			name:     "empty string key and value",
			value:    entry.New[any, any]("", ""),
			expected: "Entry{Key:, Value:}",
		},
		{
			name:     "empty string value with int key",
			value:    entry.New[any, any](0, ""),
			expected: "Entry{Key:0, Value:}",
		},
		{
			name:     "non empty string value with int key",
			value:    entry.New[any, any](987654321, "foo"),
			expected: "Entry{Key:987654321, Value:foo}",
		},
		{
			name:     "non empty string value with string kye",
			value:    entry.New[any, any]("hello", "world"),
			expected: "Entry{Key:hello, Value:world}",
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, testCase.value.String())
		})
	}
}

func TestNewFromMap(t *testing.T) {
	tests := []struct {
		name           string
		mapping        map[string]int
		expectedResult []entry.Entry[string, int]
	}{
		{
			name:           "empty map",
			mapping:        map[string]int{},
			expectedResult: []entry.Entry[string, int]{},
		},
		{
			name: "non-empty map",
			mapping: map[string]int{
				"foo": 48765,
				"bar": 1,
			},
			expectedResult: []entry.Entry[string, int]{
				entry.New("foo", 48765),
				entry.New("bar", 1),
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			entries := entry.NewFromMap(testCase.mapping)
			assert.Equal(t, len(testCase.expectedResult), len(entries))
			assert.ElementsMatch(t, testCase.expectedResult, entries)
		})
	}
}
