package exchanger

import (
	"net/http"
	"time"
)

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
	apiKey       string
	userAgent    string
	responseBody string
	rateValue    float64
	rateDate     time.Time
	name         string
	Client       *http.Client // exposed for custom http clients or testing
}
