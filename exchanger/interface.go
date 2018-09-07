package exchanger

type Rate interface {
	GetValue() float64
	GetDate() string
}

type Exchanger interface {
	Latest(string, string, map[string]string) error
	Rate
}
