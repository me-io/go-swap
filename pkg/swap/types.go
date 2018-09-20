package swap

import (
	ex "github.com/me-io/go-swap/pkg/exchanger"
)

// Swap ... main struct
type Swap struct {
	exchangers []ex.Exchanger
	cache      interface{}
}
