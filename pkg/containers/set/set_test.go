package set_test

import (
	"fmt"
	"testing"

	"github.com/kaschnit/go-ds/pkg/containers/enumerable"
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
	t.Parallel()

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
	t.Parallel()

	sets := getSetsForTest[int]()
	for _, s := range sets {
		assert.True(t, s.Empty())

		s.Add(1)
		assert.False(t, s.Empty())

		s.Remove(1)
		assert.True(t, s.Empty())
	}
}

func TestSize(t *testing.T) {
	t.Parallel()

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
			sets := getSetsForTest(testCase.initial...)
			for _, q := range sets {
				t.Run(fmt.Sprintf("%T", q), func(t *testing.T) {
					assert.Equal(t, testCase.expected, q.Size())
					assert.Equal(t, len(testCase.initial), q.Size())
				})
			}
		})
	}
}

func TestClearNonEmpty(t *testing.T) {
	t.Parallel()

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
			sets := getSetsForTest(testCase.initial...)
			for _, s := range sets {
				t.Run(fmt.Sprintf("%T", s), func(t *testing.T) {
					assert.False(t, s.Empty())

					s.Clear()
					assert.True(t, s.Empty())

					s.Clear()
					assert.True(t, s.Empty())

					s.Add(1.2345)
					assert.False(t, s.Empty())

					s.Clear()
					assert.True(t, s.Empty())
				})
			}
		})
	}
}

func TestClearEmpty(t *testing.T) {
	t.Parallel()

	sets := getSetsForTest[string]()
	for _, s := range sets {
		t.Run(fmt.Sprintf("%T", s), func(t *testing.T) {
			assert.True(t, s.Empty())

			s.Clear()
			assert.True(t, s.Empty())

			s.Clear()
			assert.True(t, s.Empty())

			s.Add("hello")
			assert.False(t, s.Empty())

			s.Clear()
			assert.True(t, s.Empty())
		})
	}
}

func TestAddNewItem(t *testing.T) {
	t.Parallel()

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
			sets := getSetsForTest(testCase.initial...)
			for _, s := range sets {
				t.Run(fmt.Sprintf("%T", s), func(t *testing.T) {
					contains := s.Contains(testCase.pushItem)
					assert.False(t, contains)

					prevSize := s.Size()
					s.Add(testCase.pushItem)
					assert.Equal(t, prevSize+1, s.Size())

					contains = s.Contains(testCase.pushItem)
					assert.True(t, contains)
				})
			}
		})
	}
}

func TestAddExistingItem(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		initial []string
		addItem string
	}{
		{
			name:    "1 items pushing a string",
			initial: []string{"abc"},
			addItem: "abc",
		},
		{
			name:    "3 items pushing a string",
			initial: []string{"1", "two", "3"},
			addItem: "3",
		},
		{
			name:    "3 items pushing an empty string",
			initial: []string{"1", "", "3"},
			addItem: "",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			sets := getSetsForTest(testCase.initial...)
			for _, s := range sets {
				t.Run(fmt.Sprintf("%T", s), func(t *testing.T) {
					contains := s.Contains(testCase.addItem)
					assert.True(t, contains)

					prevSize := s.Size()
					s.Add(testCase.addItem)
					assert.Equal(t, prevSize, s.Size())

					contains = s.Contains(testCase.addItem)
					assert.True(t, contains)
				})
			}
		})
	}
}

func TestAddAll(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		initial      []int
		addItems     []int
		expectedSize int
	}{
		{
			name:         "adding no items",
			initial:      []int{9, 8, 7},
			addItems:     []int{},
			expectedSize: 3,
		},
		{
			name:         "adding one item",
			initial:      []int{9, 8, 7},
			addItems:     []int{1},
			expectedSize: 4,
		},
		{
			name:         "adding some items that don't exist",
			initial:      []int{9, 8, 7},
			addItems:     []int{500, 1000},
			expectedSize: 5,
		},
		{
			name:         "adding some items that exist and don't exist",
			initial:      []int{9, 8, 7},
			addItems:     []int{9, 1000},
			expectedSize: 4,
		},
		{
			name:         "adding some items that already exist",
			initial:      []int{9, 8, 7},
			addItems:     []int{9, 7},
			expectedSize: 3,
		},
		{
			name:         "adding one item to an empty list",
			initial:      []int{},
			addItems:     []int{-52},
			expectedSize: 1,
		},
		{
			name:         "addding some items to an empty set",
			initial:      []int{},
			addItems:     []int{1, 2, 3},
			expectedSize: 3,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			sets := getSetsForTest(testCase.initial...)
			for _, s := range sets {
				t.Run(fmt.Sprintf("%T", s), func(t *testing.T) {
					s.AddAll(testCase.addItems...)
					assert.Equal(t, testCase.expectedSize, s.Size())
					assert.True(t, s.ContainsAll(testCase.initial...))
					assert.True(t, s.ContainsAll(testCase.addItems...))
				})
			}
		})
	}
}

func TestRemoveUntilEmpty(t *testing.T) {
	t.Parallel()

	vals := []int{1, 2, 3}
	sets := getSetsForTest(vals...)
	for _, s := range sets {
		t.Run(fmt.Sprintf("%T", s), func(t *testing.T) {
			assert.Equal(t, 3, s.Size())
			assert.True(t, s.Contains(1))

			ok := s.Remove(1)
			assert.True(t, ok)
			assert.False(t, s.Contains(1))
			assert.Equal(t, 2, s.Size())

			ok = s.Remove(1)
			assert.False(t, ok)
			assert.False(t, s.Contains(1))

			assert.Equal(t, 2, s.Size())
			assert.True(t, s.Contains(2))

			ok = s.Remove(2)
			assert.True(t, ok)
			assert.False(t, s.Contains(2))
			assert.Equal(t, 1, s.Size())

			ok = s.Remove(2)
			assert.False(t, ok)
			assert.False(t, s.Contains(2))

			assert.Equal(t, 1, s.Size())
			assert.True(t, s.Contains(3))

			ok = s.Remove(3)
			assert.True(t, ok)
			assert.False(t, s.Contains(3))
			assert.Equal(t, 0, s.Size())

			ok = s.Remove(3)
			assert.False(t, ok)
			assert.False(t, s.Contains(3))

			assert.Equal(t, 0, s.Size())
		})
	}
}

func TestRemoveAll(t *testing.T) {
	t.Parallel()

	vals := []int{1, 2, 3}
	sets := getSetsForTest(vals...)
	for _, s := range sets {
		t.Run(fmt.Sprintf("%T", s), func(t *testing.T) {
			assert.Equal(t, 3, s.Size())
			assert.True(t, s.ContainsAll(1, 2, 3))

			removed := s.RemoveAll(1, 2, 3, 4, 5)
			assert.Equal(t, 0, s.Size())
			assert.Equal(t, 3, removed)

			removed = s.RemoveAll(1, 2, 3, 4, 5)
			assert.Equal(t, 0, s.Size())
			assert.Equal(t, 0, removed)

			s.Add(-77)
			assert.Equal(t, 1, s.Size())

			removed = s.RemoveAll()
			assert.Equal(t, 1, s.Size())
			assert.Equal(t, 0, removed)

			removed = s.RemoveAll(-77)
			assert.Equal(t, 0, s.Size())
			assert.Equal(t, 1, removed)

			s.AddAll(10, 20)
			assert.Equal(t, 2, s.Size())

			removed = s.RemoveAll(10, 30)
			assert.Equal(t, 1, s.Size())
			assert.Equal(t, 1, removed)

			removed = s.RemoveAll(10, 20, 30)
			assert.Equal(t, 0, s.Size())
			assert.Equal(t, 1, removed)
		})
	}
}

func TestAnyAll(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		initial     []int
		expectedAny bool
		expectedAll bool
	}{
		{
			name:        "no values",
			initial:     []int{},
			expectedAny: false,
			expectedAll: true,
		},
		{
			name:        "no negative values with 1 item",
			initial:     []int{12},
			expectedAny: false,
			expectedAll: false,
		},
		{
			name:        "negative at index 0",
			initial:     []int{-100, 300, 57},
			expectedAny: true,
			expectedAll: false,
		},
		{
			name:        "negative at index 1",
			initial:     []int{100, -300, 57},
			expectedAny: true,
			expectedAll: false,
		},
		{
			name:        "negative at index 2",
			initial:     []int{100, 300, -57},
			expectedAny: true,
			expectedAll: false,
		},
		{
			name:        "no negative values with 3 items",
			initial:     []int{100, 300, 57},
			expectedAny: false,
			expectedAll: false,
		},
		{
			name:        "all negatives",
			initial:     []int{-100, -300, -57},
			expectedAny: true,
			expectedAll: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			sets := getSetsForTest(testCase.initial...)
			for _, s := range sets {
				t.Run(fmt.Sprintf("%T", s), func(t *testing.T) {
					isNegative := func(_ int, value int) bool {
						return value < 0
					}
					t.Run("Any", func(t *testing.T) {
						assert.Equal(t, testCase.expectedAny, s.Any(isNegative))
					})
					t.Run("All", func(t *testing.T) {
						assert.Equal(t, testCase.expectedAll, s.All(isNegative))
					})
				})
			}
		})
	}
}

func TestContainsAnyContainsAll(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		initial     []int
		values      []int
		expectedAny bool
		expectedAll bool
	}{
		{
			name:        "no values",
			initial:     []int{},
			values:      []int{},
			expectedAny: false,
			expectedAll: true,
		},
		{
			name:        "no values with attempted match",
			initial:     []int{},
			values:      []int{1, 2, 3},
			expectedAny: false,
			expectedAll: false,
		},
		{
			name:        "no values match with 1 item",
			initial:     []int{12},
			values:      []int{},
			expectedAny: false,
			expectedAll: true,
		},
		{
			name:        "no values with 3 items",
			initial:     []int{100, -300, 57},
			values:      []int{},
			expectedAny: false,
			expectedAll: true,
		},
		{
			name:        "one match with 3 itesm",
			initial:     []int{100, 300, -57},
			values:      []int{500, 700, 900, -57},
			expectedAny: true,
			expectedAll: false,
		},
		{
			name:        "no matches values with 3 items",
			initial:     []int{100, 300, 57},
			values:      []int{500, 700},
			expectedAny: false,
			expectedAll: false,
		},
		{
			name:        "1 of 3 values all match",
			initial:     []int{-100, -300, -57},
			values:      []int{-300, 500},
			expectedAny: true,
			expectedAll: false,
		},
		{
			name:        "1 of 3 values all match",
			initial:     []int{-100, -300, -57},
			values:      []int{-300},
			expectedAny: true,
			expectedAll: true,
		},
		{
			name:        "2 of 3 values all match",
			initial:     []int{-100, -300, -57},
			values:      []int{-100, -300},
			expectedAny: true,
			expectedAll: true,
		},
		{
			name:        "all values match",
			initial:     []int{-100, -300, -57},
			values:      []int{-100, -300, -57},
			expectedAny: true,
			expectedAll: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			sets := getSetsForTest(testCase.initial...)
			for _, s := range sets {
				t.Run(fmt.Sprintf("%T", s), func(t *testing.T) {
					t.Run("ContainsAny", func(t *testing.T) {
						assert.Equal(t, testCase.expectedAny, s.ContainsAny(testCase.values...))
					})
					t.Run("ContainsAll", func(t *testing.T) {
						assert.Equal(t, testCase.expectedAll, s.ContainsAll(testCase.values...))
					})
				})
			}
		})
	}
}

func TestFindOk(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		values        []int
		expectedValue int
	}{
		{
			name:          "negative at index 0",
			values:        []int{-100, 300, 57},
			expectedValue: -100,
		},
		{
			name:          "negative at index 1",
			values:        []int{100, -300, 57},
			expectedValue: -300,
		},
		{
			name:          "negative at index 2",
			values:        []int{100, 300, -57},
			expectedValue: -57,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			sets := getSetsForTest(testCase.values...)
			for _, s := range sets {
				isNegative := func(_ int, value int) bool {
					return value < 0
				}
				t.Run(fmt.Sprintf("%T", s), func(t *testing.T) {
					key, val, ok := s.Find(isNegative)
					assert.True(t, ok)
					assert.Equal(t, key, val)
					assert.Equal(t, testCase.expectedValue, val)
				})
			}
		})
	}
}

func TestFindNotOk(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		values []int
	}{
		{
			name:   "no values",
			values: []int{},
		},
		{
			name:   "no negative values with 1 item",
			values: []int{12},
		},
		{
			name:   "no negatives with 3 items",
			values: []int{100, 300, 57},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			sets := getSetsForTest(testCase.values...)
			for _, s := range sets {
				isNegative := func(_ int, value int) bool {
					return value < 0
				}
				t.Run(fmt.Sprintf("%T", s), func(t *testing.T) {
					_, _, ok := s.Find(isNegative)
					assert.False(t, ok)
				})
			}
		})
	}
}

func TestKeysValues(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		values       []int
		expectedKeys []int
	}{
		{
			name:         "no values",
			values:       []int{},
			expectedKeys: []int{},
		},
		{
			name:         "1 item",
			values:       []int{12},
			expectedKeys: []int{12},
		},
		{
			name:         "3 items",
			values:       []int{100, 300, 57},
			expectedKeys: []int{100, 300, 57},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			sets := getSetsForTest(testCase.values...)
			for _, s := range sets {
				t.Run(fmt.Sprintf("%T Keys", s), func(t *testing.T) {
					result := []int{}
					for index := range s.Keys(nil) {
						result = append(result, index)
					}
					assert.ElementsMatch(t, testCase.expectedKeys, result)
				})
				t.Run(fmt.Sprintf("%T Values", s), func(t *testing.T) {
					result := []int{}
					for index := range s.Values(nil) {
						result = append(result, index)
					}
					assert.ElementsMatch(t, testCase.expectedKeys, result)
				})
			}
		})
	}
}

func TestItems(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		values        []int
		expectedItems []enumerable.KeyValue[int, int]
	}{
		{
			name:          "no values",
			values:        []int{},
			expectedItems: []enumerable.KeyValue[int, int]{},
		},
		{
			name:   "1 item",
			values: []int{12},
			expectedItems: []enumerable.KeyValue[int, int]{
				{Key: 12, Value: 12},
			},
		},
		{
			name:   "3 items",
			values: []int{100, 300, 57},
			expectedItems: []enumerable.KeyValue[int, int]{
				{Key: 100, Value: 100},
				{Key: 300, Value: 300},
				{Key: 57, Value: 57},
			},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			sets := getSetsForTest(testCase.values...)
			for _, s := range sets {
				t.Run(fmt.Sprintf("%T", s), func(t *testing.T) {
					result := []enumerable.KeyValue[int, int]{}
					for item := range s.Items(nil) {
						result = append(result, item)
					}
					assert.ElementsMatch(t, testCase.expectedItems, result)
				})
			}
		})
	}
}
