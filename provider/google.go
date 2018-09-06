package provider

type GoogleApi struct {
	apiKey  string
	apiData string
}

func (c *GoogleApi) TestMe() string {
	println("TESTEST GoogleApi")
	return ""
}

func (c *GoogleApi) GetValue() float64 {
	println("getValue GoogleApi")
	return 1.1111
}

func (c *GoogleApi) GetDate() string {
	println("getDate GoogleApi")
	return "11"
}

func NewGoogleApi() *GoogleApi {
	r := &GoogleApi{apiKey: "12344"}
	return r
}
