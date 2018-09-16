package swap

import (
	"fmt"
	ex "github.com/me-io/go-swap/exchanger"
	"log"
	"reflect"
	"strings"
)

// Swap ... main struct
type Swap struct {
	services []ex.Exchanger
}

// NewSwap ... configure new swap instance
func NewSwap(opt ...string) *Swap {
	return &Swap{}
}

// AddExchanger ... add service to the swap stack
func (b *Swap) AddExchanger(interfaceClass ex.Exchanger) *Swap {
	b.services = append(b.services, interfaceClass)
	return b
}

// Build ... build and init swap object
func (b *Swap) Build() *Swap {
	return b
}

// Latest ... get latest rate exchange from the first api that respond from the swap stack
func (b *Swap) Latest(currencyPair string) ex.Exchanger {
	if len(b.services) < 1 {
		// configure at least one service
		log.Panic(400)
	}

	// todo
	var currentSrc ex.Exchanger
	errArr := map[string]string{}

	args := strings.Split(currencyPair, "/")
	for _, srv := range b.services {
		err := srv.Latest(args[0], args[1], nil)

		if err != nil {
			// add errors to array so we can log them
			errArr[reflect.TypeOf(srv).String()] = fmt.Sprint(err)
			continue
		}
		// assign the service after first working service and break the loop
		currentSrc = srv
		break
	}

	if currentSrc == nil {
		log.Panic(errArr)
	}
	return currentSrc
}
