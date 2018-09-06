package swap

import (
	p "github.com/me-io/go-swap/provider"
)

type Builder struct {
	services []p.ExchangeProvide
}

//
func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Add(interfaceClass p.ExchangeProvide, opt map[string]string) *Builder {
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

func (b *Builder) Build() *Builder {
	// println(fmt.Sprintf("%+v", b.services))
	//class.Test()

	return b
}

func (b *Builder) latest(currencyPair string) p.ExchangeProvide {
	// provider
	// loop on services
	// call
	// on success
	// set ApiData
	//for key, val := range b.services {
	//	println(key, val)
	//}
	v := b.services[0]
	//println(currencyPair)
	//println(r)
	//println(reflect.TypeOf(r).String())
	return v
}
