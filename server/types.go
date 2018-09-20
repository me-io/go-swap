package main

type exchangerReqObj struct {
	Name      string `json:"name"`
	UserAgent string `json:"userAgent,omitempty"`
	ApiKey    string `json:"apiKey,omitempty"`
}

type convertReqObj struct {
	Amount    float64 `json:"amount"`
	Exchanger []exchangerReqObj
	From      string `json:"from"`
	To        string `json:"to"`
}

type convertResObj struct {
	Value         float64 `json:"value"`
	Date          string  `json:"date"`
	ExchangerName string  `json:"exchangerName"`
}
