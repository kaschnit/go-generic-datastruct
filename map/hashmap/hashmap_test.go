package hashmap_test

import (
	mapp "github.com/kaschnit/go-ds/map"
	"github.com/kaschnit/go-ds/map/hashmap"
)

// Ensure that HashMap implements Map
var _ mapp.Map[string, int] = &hashmap.HashMap[string, int]{}
