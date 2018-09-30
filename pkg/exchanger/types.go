package exchanger

import (
	"net/http"
	"time"
)

// Rate ... Rate interface
type Rate interface {
	GetRateValue() float64
	GetRateDateTime() string
	GetExchangerName() string
}

// Exchanger ... Exchanger interface
type Exchanger interface {
	Latest(string, string, ...interface{}) error
	Rate
}

// attributes ... Exchanger attributes
type attributes struct {
	apiVersion   string
	apiKey       string
	userAgent    string
	responseBody string
	rateValue    float64
	rateDate     time.Time
	name         string
	Client       *http.Client // exposed for custom http clients or testing
}
