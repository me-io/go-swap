package exchanger

import "net/http"

type Rate interface {
	GetValue() float64
	GetDate() string
	GetExchangerName() string
}

type Exchanger interface {
	Latest(string, string, ...interface{}) error
	Rate
}

type attributes struct {
	responseBody string
	rateValue    float64
	rateDate     string
	name         string
	Client       *http.Client // exposed for custom http clients or testing
}
