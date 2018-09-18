package exchanger

import "net/http"

type attributes struct {
	responseBody string
	rateValue    float64
	rateDate     string
	name         string
	Client       *http.Client // exposed for custom http clients or testing
}
