package compare_test

import "github.com/kaschnit/go-ds/pkg/compare"

// Ensure that OrderedComparator implements Comparator
var _ compare.HashKey[int, int] = compare.IdentityHashKey[int]
var _ compare.HashKey[string, string] = compare.IdentityHashKey[string]
