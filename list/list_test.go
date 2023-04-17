package list_test

import (
	"fmt"
	"testing"

	"github.com/kaschnit/go-ds/enumerable"
	"github.com/kaschnit/go-ds/list"
	"github.com/kaschnit/go-ds/list/arraylist"
	"github.com/kaschnit/go-ds/list/linkedlist"
	"github.com/stretchr/testify/assert"
)

func getListsForTest[T any](values ...T) []list.List[T] {
	return []list.List[T]{
		arraylist.New(values...),
		linkedlist.NewSingleLinked(values...),
		linkedlist.NewDoubleLinked(values...),
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
			lists := getListsForTest(testCase.initial...)
			for _, l := range lists {
				t.Run(fmt.Sprintf("%T", l), func(t *testing.T) {
					assert.False(t, l.Empty())
				})
			}
		})
	}
}

func TestEmptyTrue(t *testing.T) {
	lists := getListsForTest[int]()
	for _, l := range lists {
		assert.True(t, l.Empty())

		l.Append(1)
		assert.False(t, l.Empty())

		l.PopBack()
		assert.True(t, l.Empty())
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
			lists := getListsForTest(testCase.initial...)
			for _, l := range lists {
				t.Run(fmt.Sprintf("%T", l), func(t *testing.T) {
					assert.Equal(t, testCase.expected, l.Size())
					assert.Equal(t, len(testCase.initial), l.Size())
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
			lists := getListsForTest(testCase.initial...)
			for _, l := range lists {
				t.Run(fmt.Sprintf("%T", l), func(t *testing.T) {
					assert.False(t, l.Empty())

					l.Clear()
					assert.True(t, l.Empty())

					l.Clear()
					assert.True(t, l.Empty())

					l.Append(1.2345)
					assert.False(t, l.Empty())

					l.Clear()
					assert.True(t, l.Empty())
				})
			}
		})
	}
}

func TestClearEmpty(t *testing.T) {
	lists := getListsForTest[string]()
	for _, l := range lists {
		assert.True(t, l.Empty())

		l.Clear()
		assert.True(t, l.Empty())

		l.Clear()
		assert.True(t, l.Empty())

		l.Append("hello")
		assert.False(t, l.Empty())

		l.Clear()
		assert.True(t, l.Empty())
	}
}

func TestAppend(t *testing.T) {
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
			lists := getListsForTest(testCase.initial...)
			for _, l := range lists {
				t.Run(fmt.Sprintf("%T", l), func(t *testing.T) {
					prevSize := l.Size()
					l.Append(testCase.pushItem)
					assert.Equal(t, prevSize+1, l.Size())

					actual, ok := l.GetBack()
					assert.True(t, ok)
					assert.Equal(t, testCase.pushItem, actual)
				})
			}
		})
	}
}

func TestAppendAll(t *testing.T) {
	tests := []struct {
		name         string
		initial      []int
		pushItems    []int
		expectedBack int
	}{
		{
			name:         "pushing no items",
			initial:      []int{9, 8, 7},
			pushItems:    []int{},
			expectedBack: 7,
		},
		{
			name:         "pushing one item",
			initial:      []int{9, 8, 7},
			pushItems:    []int{1},
			expectedBack: 1,
		},
		{
			name:         "pushing some items",
			initial:      []int{9, 8, 7},
			pushItems:    []int{500, 1000},
			expectedBack: 1000,
		},
		{
			name:         "pushing one item onto an empty stack",
			initial:      []int{},
			pushItems:    []int{-52},
			expectedBack: -52,
		},
		{
			name:         "pushing some items onto an empty stack",
			initial:      []int{},
			pushItems:    []int{1, 2, 3},
			expectedBack: 3,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			lists := getListsForTest(testCase.initial...)
			for _, l := range lists {
				t.Run(fmt.Sprintf("%T", l), func(t *testing.T) {
					prevSize := l.Size()
					l.AppendAll(testCase.pushItems...)
					assert.Equal(t, prevSize+len(testCase.pushItems), l.Size())

					actual, ok := l.GetBack()
					assert.True(t, ok)
					assert.Equal(t, testCase.expectedBack, actual)
				})
			}
		})
	}
}

func TestPrepend(t *testing.T) {
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
			lists := getListsForTest(testCase.initial...)
			for _, l := range lists {
				t.Run(fmt.Sprintf("%T", l), func(t *testing.T) {
					prevSize := l.Size()
					l.Prepend(testCase.pushItem)
					assert.Equal(t, prevSize+1, l.Size())

					actual, ok := l.GetFront()
					assert.True(t, ok)
					assert.Equal(t, testCase.pushItem, actual)
				})
			}
		})
	}
}

func TestPrependAll(t *testing.T) {
	tests := []struct {
		name          string
		initial       []int
		pushItems     []int
		expectedFront int
	}{
		{
			name:          "pushing no items",
			initial:       []int{9, 8, 7},
			pushItems:     []int{},
			expectedFront: 9,
		},
		{
			name:          "pushing one item",
			initial:       []int{9, 8, 7},
			pushItems:     []int{1},
			expectedFront: 1,
		},
		{
			name:          "pushing some items",
			initial:       []int{9, 8, 7},
			pushItems:     []int{500, 1000},
			expectedFront: 500,
		},
		{
			name:          "pushing one item onto an empty stack",
			initial:       []int{},
			pushItems:     []int{-52},
			expectedFront: -52,
		},
		{
			name:          "pushing some items onto an empty stack",
			initial:       []int{},
			pushItems:     []int{1, 2, 3},
			expectedFront: 1,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			lists := getListsForTest(testCase.initial...)
			for _, l := range lists {
				t.Run(fmt.Sprintf("%T", l), func(t *testing.T) {
					prevSize := l.Size()
					l.PrependAll(testCase.pushItems...)
					assert.Equal(t, prevSize+len(testCase.pushItems), l.Size())

					actual, ok := l.GetFront()
					assert.True(t, ok)
					assert.Equal(t, testCase.expectedFront, actual)
				})
			}
		})
	}
}

func TestInsertOk(t *testing.T) {
	tests := []struct {
		name    string
		initial []string
		item    string
		index   int
	}{
		{
			name:    "insert into an empty list",
			initial: []string{},
			item:    "foo",
			index:   0,
		},
		{
			name:    "insert at the beginning of a 1-item list",
			initial: []string{"bar"},
			item:    "foo",
			index:   0,
		},
		{
			name:    "insert at the end of a 1-item list",
			initial: []string{"bar"},
			item:    "foo",
			index:   1,
		},
		{
			name:    "insert at the beginning of a 3-item list",
			initial: []string{"foo", "bar", "baz"},
			item:    "hello",
			index:   0,
		},
		{
			name:    "insert at index 1 of a 3-item list",
			initial: []string{"foo", "bar", "baz"},
			item:    "hello",
			index:   1,
		},
		{
			name:    "insert at index 2 of a 3-item ArrayList",
			initial: []string{"foo", "bar", "baz"},
			item:    "hello",
			index:   2,
		},
		{
			name:    "insert at the end of a 3-item ArrayList",
			initial: []string{"foo", "bar", "baz"},
			item:    "hello",
			index:   3,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			lists := getListsForTest(testCase.initial...)
			for _, l := range lists {
				t.Run(fmt.Sprintf("%T", l), func(t *testing.T) {
					prevSize := l.Size()
					ok := l.Insert(testCase.index, testCase.item)
					assert.True(t, ok)
					assert.Equal(t, prevSize+1, l.Size())

					item, ok := l.Get(testCase.index)
					assert.True(t, ok)
					assert.Equal(t, testCase.item, item)
				})
			}
		})
	}
}

func TestInsertNotOk(t *testing.T) {
	tests := []struct {
		name    string
		initial []string
		item    string
		index   int
	}{
		{
			name:    "insert before an empty ArrayList",
			initial: []string{},
			item:    "foo",
			index:   -1,
		},
		{
			name:    "insert after an empty ArrayList",
			initial: []string{},
			item:    "foo",
			index:   1,
		},
		{
			name:    "insert way after an empty ArrayList",
			initial: []string{},
			item:    "foo",
			index:   5,
		},
		{
			name:    "insert before a 1-item ArrayList",
			initial: []string{"bar"},
			item:    "foo",
			index:   -1,
		},
		{
			name:    "insert after the end of a 1-item ArrayList",
			initial: []string{"bar"},
			item:    "foo",
			index:   2,
		},
		{
			name:    "insert way after the end of a 1-item ArrayList",
			initial: []string{"bar"},
			item:    "foo",
			index:   4,
		},
		{
			name:    "insert way before a 3-item ArrayList",
			initial: []string{"foo", "bar", "baz"},
			item:    "helloooo",
			index:   -4,
		},
		{
			name:    "insert before a 3-item ArrayList",
			initial: []string{"foo", "bar", "baz"},
			item:    "helloooo",
			index:   -1,
		},
		{
			name:    "insert after the end of a 3-item ArrayList",
			initial: []string{"foo", "bar", "baz"},
			item:    "helloooo",
			index:   4,
		},
		{
			name:    "insert way after the end of a 3-item ArrayList",
			initial: []string{"foo", "bar", "baz"},
			item:    "helloooo",
			index:   12,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			lists := getListsForTest(testCase.initial...)
			for _, l := range lists {
				t.Run(fmt.Sprintf("%T", l), func(t *testing.T) {
					prevSize := l.Size()
					ok := l.Insert(testCase.index, testCase.item)
					assert.False(t, ok)
					assert.Equal(t, prevSize, l.Size())
				})
			}
		})
	}
}

func TestInsertAllOk(t *testing.T) {
	tests := []struct {
		name     string
		initial  []string
		items    []string
		index    int
		expected []string
	}{
		{
			name:     "insert into an empty list",
			initial:  []string{},
			items:    []string{"foo"},
			index:    0,
			expected: []string{"foo"},
		},
		{
			name:     "insert at the beginning of a 1-item list",
			initial:  []string{"bar"},
			items:    []string{"foo"},
			index:    0,
			expected: []string{"foo", "bar"},
		},
		{
			name:     "insert at the end of a 1-item list",
			initial:  []string{"bar"},
			items:    []string{"foo"},
			index:    1,
			expected: []string{"bar", "foo"},
		},
		{
			name:     "insert at the beginning of a 3-item list",
			initial:  []string{"foo", "bar", "baz"},
			items:    []string{"hello", "goodbye"},
			index:    0,
			expected: []string{"hello", "goodbye", "foo", "bar", "baz"},
		},
		{
			name:     "insert at index 1 of a 3-item list",
			initial:  []string{"foo", "bar", "baz"},
			items:    []string{"hello", "goodbye", "no"},
			index:    1,
			expected: []string{"foo", "hello", "goodbye", "no", "bar", "baz"},
		},
		{
			name:     "insert at index 2 of a 3-item list",
			initial:  []string{"foo", "bar", "baz"},
			items:    []string{"hello", "goodbye", "no", "yes"},
			index:    2,
			expected: []string{"foo", "bar", "hello", "goodbye", "no", "yes", "baz"},
		},
		{
			name:     "insert a single item at index 2 of a 3-item ArrayList",
			initial:  []string{"foo", "bar", "baz"},
			items:    []string{"hello"},
			index:    2,
			expected: []string{"foo", "bar", "hello", "baz"},
		},
		{
			name:     "insert at the end of a 3-item list",
			initial:  []string{"foo", "bar", "baz"},
			items:    []string{"hello", "goodbye", "no", "yes"},
			index:    3,
			expected: []string{"foo", "bar", "baz", "hello", "goodbye", "no", "yes"},
		},
		{
			name:     "insert nothing in the middle of a 3-item ArrayList",
			initial:  []string{"foo", "bar", "baz"},
			items:    []string{},
			index:    1,
			expected: []string{"foo", "bar", "baz"},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			lists := getListsForTest(testCase.initial...)
			for _, l := range lists {
				t.Run(fmt.Sprintf("%T", l), func(t *testing.T) {
					prevSize := l.Size()
					ok := l.InsertAll(testCase.index, testCase.items...)
					assert.True(t, ok)
					assert.Equal(t, prevSize+len(testCase.items), l.Size())
					for i := range testCase.expected {
						expected := testCase.expected[i]
						actual, ok := l.Get(i)
						assert.True(t, ok)
						assert.Equal(t, expected, actual)
					}
				})
			}
		})
	}
}

func TestInsertAllNotOk(t *testing.T) {
	tests := []struct {
		name    string
		initial []string
		items   []string
		index   int
	}{
		{
			name:    "insert before an empty ArrayList",
			initial: []string{},
			items:   []string{"foo"},
			index:   -1,
		},
		{
			name:    "insert before an empty ArrayList",
			initial: []string{},
			items:   []string{"foo"},
			index:   1,
		},
		{
			name:    "insert way before the start of a 3-item ArrayList",
			initial: []string{"foo", "bar", "baz"},
			items:   []string{"hello", "goodbye", "no", "yes"},
			index:   -5,
		},
		{
			name:    "insert before the start of a 3-item ArrayList",
			initial: []string{"foo", "bar", "baz"},
			items:   []string{"hello"},
			index:   -1,
		},
		{
			name:    "insert past the end of a 3-item ArrayList",
			initial: []string{"foo", "bar", "baz"},
			items:   []string{"hello", "goodbye", "no", "yes"},
			index:   4,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			lists := getListsForTest(testCase.initial...)
			for _, l := range lists {
				t.Run(fmt.Sprintf("%T", l), func(t *testing.T) {
					prevSize := l.Size()
					ok := l.InsertAll(testCase.index, testCase.items...)
					assert.False(t, ok)
					assert.Equal(t, prevSize, l.Size())
				})
			}
		})
	}
}

func TestPopBackUntilEmpty(t *testing.T) {
	vals := []int{1, 2, 3}
	lists := getListsForTest(vals...)
	for _, l := range lists {
		val, ok := l.PopBack()
		assert.True(t, ok)
		assert.Equal(t, 3, val)

		val, ok = l.PopBack()
		assert.True(t, ok)
		assert.Equal(t, 2, val)

		val, ok = l.PopBack()
		assert.True(t, ok)
		assert.Equal(t, 1, val)

		val, ok = l.PopBack()
		assert.False(t, ok, "Expected not ok but was ok, val was %v", val)

		val, ok = l.PopFront()
		assert.False(t, ok, "Expected not ok but was ok, val was %v", val)
	}
}

func TestPopFrontUntilEmpty(t *testing.T) {
	vals := []int{1, 2, 3}
	lists := getListsForTest(vals...)
	for _, l := range lists {
		val, ok := l.PopFront()
		assert.True(t, ok)
		assert.Equal(t, 1, val)

		val, ok = l.PopFront()
		assert.True(t, ok)
		assert.Equal(t, 2, val)

		val, ok = l.PopFront()
		assert.True(t, ok)
		assert.Equal(t, 3, val)

		val, ok = l.PopFront()
		assert.False(t, ok, "Expected not ok but was ok, val was %v", val)

		val, ok = l.PopBack()
		assert.False(t, ok, "Expected not ok but was ok, val was %v", val)
	}
}

func TestGetFrontNotOk(t *testing.T) {
	lists := getListsForTest[string]()
	for _, l := range lists {
		t.Run(fmt.Sprintf("%T", l), func(t *testing.T) {
			_, ok := l.GetFront()
			assert.False(t, ok)
		})
	}
}

func TestGetBackNotOk(t *testing.T) {
	lists := getListsForTest[string]()
	for _, l := range lists {
		t.Run(fmt.Sprintf("%T", l), func(t *testing.T) {
			_, ok := l.GetBack()
			assert.False(t, ok)
		})
	}
}

func TestGetNotOk(t *testing.T) {
	tests := []struct {
		name         string
		initial      []string
		notOkIndices []int
	}{
		{
			name:         "empty ArrayList",
			initial:      []string{},
			notOkIndices: []int{-77, -1, 0, 1, 12, 20},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			lists := getListsForTest(testCase.initial...)
			for _, l := range lists {
				t.Run(fmt.Sprintf("%T", l), func(t *testing.T) {
					for _, index := range testCase.notOkIndices {
						t.Run(fmt.Sprintf("index %d", index), func(t *testing.T) {
							_, ok := l.Get(index)
							assert.False(t, ok)
						})
					}
				})
			}
		})
	}
}

func TestForEach(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		expected int
	}{
		{
			name:     "sum nothing",
			initial:  []int{},
			expected: 0,
		},
		{
			name:     "sum a single number",
			initial:  []int{12},
			expected: 12,
		},
		{
			name:     "sum a few numbers",
			initial:  []int{-100, 300, 57},
			expected: 257,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			lists := getListsForTest(testCase.initial...)
			for _, l := range lists {
				total := 0
				t.Run(fmt.Sprintf("%T", l), func(t *testing.T) {
					l.ForEach(func(key int, value int) {
						total += value
					})
					assert.Equal(t, testCase.expected, total)
				})
			}
		})
	}
}

func TestAnyAll(t *testing.T) {
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
			lists := getListsForTest(testCase.initial...)
			for _, l := range lists {
				t.Run(fmt.Sprintf("%T", l), func(t *testing.T) {
					isNegative := func(_ int, value int) bool {
						return value < 0
					}
					t.Run("Any", func(t *testing.T) {
						assert.Equal(t, testCase.expectedAny, l.Any(isNegative))
					})
					t.Run("All", func(t *testing.T) {
						assert.Equal(t, testCase.expectedAll, l.All(isNegative))
					})
				})
			}
		})
	}
}

func TestFindOk(t *testing.T) {
	tests := []struct {
		name          string
		initial       []int
		expectedIndex int
		expectedValue int
	}{
		{
			name:          "negative at index 0",
			initial:       []int{-100, 300, 57},
			expectedIndex: 0,
			expectedValue: -100,
		},
		{
			name:          "negative at index 1",
			initial:       []int{100, -300, 57},
			expectedIndex: 1,
			expectedValue: -300,
		},
		{
			name:          "negative at index 2",
			initial:       []int{100, 300, -57},
			expectedIndex: 2,
			expectedValue: -57,
		},
		{
			name:          "negative at index 2 and 3",
			initial:       []int{100, -300, -57},
			expectedIndex: 1,
			expectedValue: -300,
		},
		{
			name:          "negative at index 1 and 3",
			initial:       []int{-100, 300, -100},
			expectedIndex: 0,
			expectedValue: -100,
		},
		{
			name:          "all negatives",
			initial:       []int{-100, -300, -57},
			expectedIndex: 0,
			expectedValue: -100,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			lists := getListsForTest(testCase.initial...)
			for _, l := range lists {
				isNegative := func(_ int, value int) bool {
					return value < 0
				}
				t.Run(fmt.Sprintf("%T", l), func(t *testing.T) {
					idx, val, ok := l.Find(isNegative)
					assert.True(t, ok)
					assert.Equal(t, testCase.expectedIndex, idx)
					assert.Equal(t, testCase.expectedValue, val)
				})
			}
		})
	}
}

func TestFindNotOk(t *testing.T) {
	tests := []struct {
		name    string
		initial []int
	}{
		{
			name:    "no values",
			initial: []int{},
		},
		{
			name:    "no negative values with 1 item",
			initial: []int{12},
		},
		{
			name:    "no negatives with 3 items",
			initial: []int{100, 300, 57},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			lists := getListsForTest(testCase.initial...)
			for _, l := range lists {
				isNegative := func(_ int, value int) bool {
					return value < 0
				}
				t.Run(fmt.Sprintf("%T", l), func(t *testing.T) {
					_, _, ok := l.Find(isNegative)
					assert.False(t, ok)
				})
			}
		})
	}
}

func TestKeys(t *testing.T) {
	tests := []struct {
		name         string
		initial      []int
		expectedKeys []int
	}{
		{
			name:         "no values",
			initial:      []int{},
			expectedKeys: []int{},
		},
		{
			name:         "1 item",
			initial:      []int{12},
			expectedKeys: []int{0},
		},
		{
			name:         "3 items",
			initial:      []int{100, 300, 57},
			expectedKeys: []int{0, 1, 2},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			lists := getListsForTest(testCase.initial...)
			for _, l := range lists {
				t.Run(fmt.Sprintf("%T", l), func(t *testing.T) {
					result := []int{}
					for index := range l.Keys(nil) {
						result = append(result, index)
					}
					assert.Equal(t, testCase.expectedKeys, result)
				})
			}
		})
	}
}

func TestKeysAbort(t *testing.T) {
	tests := []struct {
		name         string
		initial      []int
		abortAtIndex int
		expectedKeys []int
	}{
		{
			name:         "abort at index 0",
			initial:      []int{100, 200, 500},
			abortAtIndex: 0,
			expectedKeys: []int{0},
		},
		{
			name:         "abort at index 1",
			initial:      []int{7, 8, -20, 3, 7},
			abortAtIndex: 1,
			expectedKeys: []int{0, 1},
		},
		{
			name:         "abort at index 2",
			initial:      []int{100, 200, 500, 700},
			abortAtIndex: 2,
			expectedKeys: []int{0, 1, 2},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			lists := getListsForTest(testCase.initial...)
			for _, l := range lists {
				t.Run(fmt.Sprintf("%T", l), func(t *testing.T) {
					result := []int{}
					aborter := make(chan struct{})
					for index := range l.Keys(aborter) {
						result = append(result, index)
						if index == testCase.abortAtIndex {
							close(aborter)
							break
						}
					}
					assert.Equal(t, testCase.expectedKeys, result)
				})
			}
		})
	}
}

func TestValues(t *testing.T) {
	tests := []struct {
		name           string
		initial        []int
		expectedValues []int
	}{
		{
			name:           "no values",
			initial:        []int{},
			expectedValues: []int{},
		},
		{
			name:           "1 item",
			initial:        []int{12},
			expectedValues: []int{12},
		},
		{
			name:           "3 items",
			initial:        []int{100, 300, 57},
			expectedValues: []int{100, 300, 57},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			lists := getListsForTest(testCase.initial...)
			for _, l := range lists {
				t.Run(fmt.Sprintf("%T", l), func(t *testing.T) {
					result := []int{}
					for value := range l.Values(nil) {
						result = append(result, value)
					}
					assert.Equal(t, testCase.expectedValues, result)
				})
			}
		})
	}
}

func TestValuesAbort(t *testing.T) {
	tests := []struct {
		name           string
		initial        []int
		abortAtIndex   int
		expectedValues []int
	}{
		{
			name:           "abort at index 0",
			initial:        []int{100, 200, 500},
			abortAtIndex:   0,
			expectedValues: []int{100},
		},
		{
			name:           "abort at index 1",
			initial:        []int{7, 8, -20, 3, 7},
			abortAtIndex:   1,
			expectedValues: []int{7, 8},
		},
		{
			name:           "abort at index 2",
			initial:        []int{100, 200, 500, 700},
			abortAtIndex:   2,
			expectedValues: []int{100, 200, 500},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			lists := getListsForTest(testCase.initial...)
			for _, l := range lists {
				t.Run(fmt.Sprintf("%T", l), func(t *testing.T) {
					result := []int{}
					aborter := make(chan struct{})
					i := 0
					for value := range l.Values(aborter) {
						result = append(result, value)
						if i == testCase.abortAtIndex {
							close(aborter)
							break
						}
						i += 1
					}
					assert.Equal(t, testCase.expectedValues, result)
				})
			}
		})
	}
}

func TestItems(t *testing.T) {
	tests := []struct {
		name          string
		initial       []int
		expectedItems []enumerable.KeyValue[int, int]
	}{
		{
			name:          "no values",
			initial:       []int{},
			expectedItems: []enumerable.KeyValue[int, int]{},
		},
		{
			name:    "1 item",
			initial: []int{12},
			expectedItems: []enumerable.KeyValue[int, int]{
				{Key: 0, Value: 12},
			},
		},
		{
			name:    "3 items",
			initial: []int{100, 300, 57},
			expectedItems: []enumerable.KeyValue[int, int]{
				{Key: 0, Value: 100},
				{Key: 1, Value: 300},
				{Key: 2, Value: 57},
			},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			lists := getListsForTest(testCase.initial...)
			for _, l := range lists {
				t.Run(fmt.Sprintf("%T", l), func(t *testing.T) {
					result := []enumerable.KeyValue[int, int]{}
					for item := range l.Items(nil) {
						result = append(result, item)
					}
					assert.Equal(t, testCase.expectedItems, result)
				})
			}
		})
	}
}

func TestItemsAbort(t *testing.T) {
	tests := []struct {
		name          string
		initial       []int
		abortAtIndex  int
		expectedItems []enumerable.KeyValue[int, int]
	}{
		{
			name:         "abort at index 0",
			initial:      []int{100, 200, 500},
			abortAtIndex: 0,
			expectedItems: []enumerable.KeyValue[int, int]{
				{Key: 0, Value: 100},
			},
		},
		{
			name:         "abort at index 1",
			initial:      []int{7, 8, -20, 3, 7},
			abortAtIndex: 1,
			expectedItems: []enumerable.KeyValue[int, int]{
				{Key: 0, Value: 7},
				{Key: 1, Value: 8},
			},
		},
		{
			name:         "abort at index 2",
			initial:      []int{100, 200, 500, 700},
			abortAtIndex: 2,
			expectedItems: []enumerable.KeyValue[int, int]{
				{Key: 0, Value: 100},
				{Key: 1, Value: 200},
				{Key: 2, Value: 500},
			},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			lists := getListsForTest(testCase.initial...)
			for _, l := range lists {
				t.Run(fmt.Sprintf("%T", l), func(t *testing.T) {
					result := []enumerable.KeyValue[int, int]{}
					aborter := make(chan struct{})
					i := 0
					for item := range l.Items(aborter) {
						result = append(result, item)
						if i == testCase.abortAtIndex {
							close(aborter)
							break
						}
						i += 1
					}
					assert.Equal(t, testCase.expectedItems, result)
				})
			}
		})
	}
}

func TestIteration(t *testing.T) {
	values := []int{100, 200, 500}
	lists := getListsForTest(values...)
	for _, l := range lists {
		t.Run(fmt.Sprintf("%T", l), func(t *testing.T) {
			// Iterator to element 0
			itr, ok := l.Iterator()
			assert.True(t, ok)
			assert.True(t, itr.HasNext())

			key, ok := itr.Key()
			assert.True(t, ok)
			assert.Equal(t, 0, key)

			val, ok := itr.Value()
			assert.True(t, ok)
			assert.Equal(t, 100, val)

			// Iterator to element 1
			itr, ok = itr.Next()
			assert.True(t, ok)
			assert.True(t, itr.HasNext())

			key, ok = itr.Key()
			assert.True(t, ok)
			assert.Equal(t, 1, key)

			val, ok = itr.Value()
			assert.True(t, ok)
			assert.Equal(t, 200, val)

			// Iterator to element 2
			itr, ok = itr.Next()
			assert.True(t, ok)
			assert.False(t, itr.HasNext())

			key, ok = itr.Key()
			assert.True(t, ok)
			assert.Equal(t, 2, key)

			val, ok = itr.Value()
			assert.True(t, ok)
			assert.Equal(t, 500, val)

			// Invalid iterator
			_, ok = itr.Next()
			assert.False(t, ok)
		})
	}
}

func TestIteration_Empty(t *testing.T) {
	lists := getListsForTest[int]()
	for _, l := range lists {
		t.Run(fmt.Sprintf("%T - Iterator()", l), func(t *testing.T) {
			_, ok := l.Iterator()
			assert.False(t, ok)
		})
	}
}
