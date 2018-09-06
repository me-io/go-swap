package provider

type Rate interface {
	GetValue() float64
	GetDate() string
}

type ExchangeProvide interface {
	TestMe() string
	Rate
}