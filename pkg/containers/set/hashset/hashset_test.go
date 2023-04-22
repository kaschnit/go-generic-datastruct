package hashset_test

import (
	"github.com/kaschnit/go-ds/pkg/containers/set"
	"github.com/kaschnit/go-ds/pkg/containers/set/hashset"
)

// Ensure that HashSet implements Set
var _ set.Set[int] = hashset.New(1)
