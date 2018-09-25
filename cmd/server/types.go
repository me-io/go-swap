package main

import (
	"github.com/op/go-logging"
)

type exchangerReqObj struct {
	Name      string `json:"name"`
	UserAgent string `json:"userAgent,omitempty"`
	ApiKey    string `json:"apiKey,omitempty"`
}

type convertReqObj struct {
	Amount        float64 `json:"amount"`
	Exchanger     []exchangerReqObj
	From          string `json:"from"`
	To            string `json:"to"`
	CacheTime     string `json:"cacheTime"`
	DecimalPoints int    `json:"decimalPoints"`
}

type convertResObj struct {
	Value         float64 `json:"value"`
	Amount        float64 `json:"amount"`
	DateTime      string  `json:"dateTime"`
	ExchangerName string  `json:"exchangerName"`
}

// Password is just an example type implementing the Redactor interface. Any
// time this is logged, the Redacted() function will be called.
type Secret string

func (p Secret) Redacted() interface{} {
	return logging.Redact(string(p))
}
