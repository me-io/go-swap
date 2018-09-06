package provider

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

type GoogleApi struct {
	apiKey       string
	responseBody string
	rateValue    float64
	rateDate     string
}

// ref @link https://github.com/florianv/exchanger/blob/master/src/Service/Google.php
var GoogleApiUrl = "https://www.google.com/search?q=1+%s+to+%s&ncr=1"
var GoogleApiHeaders = map[string][]string{
	"Accept":     {"text/html"},
	"User-Agent": {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:21.0) Gecko/20100101 Firefox/21.0"},
}

func (c *GoogleApi) RequestLatest(from string, to string) *GoogleApi {

	client := http.Client{}
	url := fmt.Sprintf(GoogleApiUrl, from, to)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header = GoogleApiHeaders
	res, err := client.Do(req)

	if err != nil {
		// todo handle error
		panic("Body in Null")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// todo handle error
		panic("Body in Null")
	}

	c.responseBody = string(body)
	return c
}

func (c *GoogleApi) GetValue() float64 {
	return c.rateValue
}

func (c *GoogleApi) GetDate() string {
	return c.rateDate
}

func (c *GoogleApi) Latest(from string, to string) {

	// todo cache layer
	c.RequestLatest(from, to)
	var validID = regexp.MustCompile(`knowledge-currency__tgt-input(.*)value="([1-9.]{0,10})" (.*)"`)
	f := validID.FindStringSubmatch(c.responseBody)
	// todo handle error
	c.rateValue, _ = strconv.ParseFloat(f[2], 64)
}

func NewGoogleApi() *GoogleApi {
	r := &GoogleApi{apiKey: "12344"}
	return r
}
