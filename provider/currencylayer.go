package provider

type CurrencyLayerApi struct {
	apiKey  string
	apiData string
}

func (c *CurrencyLayerApi) Request(from string, to string) string {
	println("TESTEST CurrencyLayerApi")
	return ""
}

func (c *CurrencyLayerApi) GetValue() float64 {
	println("TESTEST GoogleApi")
	return 1.1111
}

func (c *CurrencyLayerApi) GetDate() string {
	println("TESTEST GoogleApi")
	return "11"
}

func NewCurrencyLayerApi() *CurrencyLayerApi {
	r := &CurrencyLayerApi{apiKey: "111"}
	return r
}
