package hashmap_test

import (
	mapp "github.com/kaschnit/go-ds/pkg/containers/map"
	"github.com/kaschnit/go-ds/pkg/containers/map/hashmap"
)

// Ensure that HashMap implements Map
var _ mapp.Map[string, int] = &hashmap.HashMap[string, int]{}
