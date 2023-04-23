package compare_test

import "github.com/kaschnit/go-ds/pkg/compare"

// Ensure that OrderedComparator implements Comparator
var _ compare.Comparator[int] = compare.OrderedComparator[int]
