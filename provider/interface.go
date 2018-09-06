package provider

type Rate interface {
	GetValue() float64
	GetDate() string
}

type ExchangeProvider interface {
	Latest(string, string)
	Rate
}
