package main

import (
	"github.com/op/go-logging"
)

// exchangerReqObj ... Exchanger object that form array of Exchangers in the Convert Request Data Object
type exchangerReqObj struct {
	Name       string `json:"name"`
	UserAgent  string `json:"userAgent,omitempty"`
	ApiKey     string `json:"apiKey,omitempty"`
	ApiVersion string `json:"apiVersion,omitempty"`
}

// convertReqObj ... Convert Request Data Object
type convertReqObj struct {
	Amount        float64 `json:"amount"`
	Exchanger     []exchangerReqObj
	From          string `json:"from"`
	To            string `json:"to"`
	CacheTime     string `json:"cacheTime"`
	DecimalPoints int    `json:"decimalPoints"`
}

// convertResObj ... Convert Response Data Object
type convertResObj struct {
	From            string  `json:"from"`
	To              string  `json:"to"`
	ExchangerName   string  `json:"exchangerName"`
	ExchangeValue   float64 `json:"exchangeValue"`
	OriginalAmount  float64 `json:"originalAmount"`
	ConvertedAmount float64 `json:"convertedAmount"`
	ConvertedText   string  `json:"convertedText"`
	RateDateTime    string  `json:"rateDateTime"`
	RateFromCache   bool    `json:"rateFromCache"`
}

// Secret ... Secret Type for logging in the Logger
// Password is just an example type implementing the Redactor interface. Any
// time this is logged, the Redacted() function will be called.
type Secret string

// Redacted ... Secret  Redacted function
func (p Secret) Redacted() interface{} {
	return logging.Redact(string(p))
}
