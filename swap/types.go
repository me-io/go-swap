package swap

import (
	ex "github.com/me-io/go-swap/exchanger"
)

// Swap ... main struct
type Swap struct {
	exchangers []ex.Exchanger
	cache      interface{}
}
