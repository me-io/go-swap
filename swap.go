package swap

import (
	"fmt"
	ex "github.com/meabed/go-swap/exchanger"
	"log"
	"reflect"
	"strings"
)

type Swap struct {
	services []ex.Exchanger
}

//
func NewSwap() *Swap {
	return &Swap{}
}

func (b *Swap) AddExchanger(interfaceClass ex.Exchanger) *Swap {
	b.services = append(b.services, interfaceClass)
	return b
}

func (b *Swap) Build() *Swap {
	return b
}

func (b *Swap) latest(currencyPair string) ex.Exchanger {
	// todo
	var currentSrc ex.Exchanger = nil
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
