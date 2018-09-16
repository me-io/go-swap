package exchanger

type Rate interface {
	GetValue() float64
	GetDate() string
	GetExchangerName() string
}

type Exchanger interface {
	Latest(string, string, ...interface{}) error
	Rate
}
