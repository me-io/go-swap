package swap

import (
	ex "github.com/me-io/go-swap/exchanger"
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
	//fmt.Println(fmt.Sprintf("%+v", class))
	//fmt.Println(fmt.Sprintf("%+v", b.services))
	// b.services = append(b.services, reflect.TypeOf(interfaceClass).String())
	b.services = append(b.services, interfaceClass)
	//interfaceClass.TestMe()
	//fmt.Println(fmt.Sprintf("%+v", class))
	//fmt.Println(fmt.Sprintf("%+v", b.services))
	//Fin.testMe(class)
	return b
}

func (b *Swap) Build() *Swap {
	// println(fmt.Sprintf("%+v", b.services))
	//class.Test()
	return b
}

func (b *Swap) latest(currencyPair string) ex.Exchanger {
	// todo
	// provider
	// loop on services
	// call
	// on success
	// set ApiData
	//for key, val := range b.services {
	//	println(key, val)
	//}
	v := b.services[0]
	args := strings.Split(currencyPair, "/")
	v.Latest(args[0], args[1])

	//print(res)
	//println(currencyPair)
	//println(r)
	//println(reflect.TypeOf(r).String())
	return v
}
