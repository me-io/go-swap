package swap

import (
	ex "github.com/meabed/go-swap/exchanger"
	"strings"
)

type Swap struct {
	services []ex.Exchanger
}

//
func NewSwap() *Swap {
	return &Swap{}
}

func (b *Swap) AddExchanger(interfaceClass ex.Exchanger, opt map[string]string) *Swap {
	b.services = append(b.services, interfaceClass)
	return b
}

func (b *Swap) Build() *Swap {
	return b
}

func (b *Swap) latest(currencyPair string) ex.Exchanger {
	// todo
	var currentSrc ex.Exchanger = nil
	args := strings.Split(currencyPair, "/")
	for _, srv := range b.services {
		srv.Latest(args[0], args[1], nil)
		// todo handle error and go to second in stack
		currentSrc = srv
	}

	return currentSrc
}
