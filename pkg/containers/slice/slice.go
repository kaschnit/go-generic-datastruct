package slice

import (
	"fmt"
	"strings"

	"github.com/kaschnit/go-ds/pkg/containers/enumerable"
)

type Slice[V any] []V

func (s Slice[V]) String() string {
	sb := strings.Builder{}
	sb.WriteString("Slice\n")

	strs := []string{}
	for _, value := range s {
		strs = append(strs, fmt.Sprintf("%v", value))
	}

	sb.WriteString(strings.Join(strs, ","))

	return sb.String()
}

func (s Slice[V]) ForEach(op enumerable.Op[int, V]) {
	for i, value := range s {
		op(i, value)
	}
}

func (s Slice[V]) Filter(predicate enumerable.Predicate[int, V]) Slice[V] {
	result := make(Slice[V], 0)

	for i, value := range s {
		if predicate(i, value) {
			result = append(result, value)
		}
	}

	return result
}

func (s Slice[V]) All(predicate enumerable.Predicate[int, V]) bool {
	for i, value := range s {
		if !predicate(i, value) {
			return false
		}
	}

	return true
}

func (s Slice[V]) Any(predicate enumerable.Predicate[int, V]) bool {
	for i, value := range s {
		if predicate(i, value) {
			return true
		}
	}

	return false
}

func (s Slice[V]) Find(predicate enumerable.Predicate[int, V]) (int, V, bool) {
	for i, value := range s {
		if predicate(i, value) {
			return i, value, true
		}
	}

	return -1, *new(V), false
}

func (s Slice[V]) Reverse() {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

func (s Slice[V]) Reversed() Slice[V] {
	newSlice := Slice[V](make([]V, 0, len(s)))

	for i := len(s) - 1; i >= 0; i-- {
		newSlice = append(newSlice, s[i])
	}

	return newSlice
}
