package exchanger

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

type GoogleApi struct {
	apiKey       string
	responseBody string
	rateValue    float64
	rateDate     string
}

// ref @link https://github.com/florianv/exchanger/blob/master/src/Service/Google.php
// example : https://www.google.com/search?q=1+USD+to+USD&ncr=1
// example : https://www.google.com/search?q=1+USD+to+EGP&ncr=1
// example : https://www.google.com/search?q=1+USD+to+AED&ncr=1
var GoogleApiUrl = `https://www.google.com/search?q=1+%s+to+%s&ncr=1`
var GoogleApiHeaders = map[string][]string{
	`Accept`:     {`text/html`},
	`User-Agent`: {`Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:21.0) Gecko/20100101 Firefox/21.0`},
}

func (c *GoogleApi) RequestRate(from string, to string, opt map[string]string) (*GoogleApi, error) {

	// todo add option opt to add more headers or client configurations
	// free mem-leak
	// optimize for memory leak
	// todo optimize curl connection
	keepAliveTimeout := 600 * time.Second
	timeout := 5 * time.Second
	defaultTransport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: keepAliveTimeout,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
	}

	client := &http.Client{
		Transport: defaultTransport,
		Timeout:   timeout,
	}

	url := fmt.Sprintf(GoogleApiUrl, from, to)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header = GoogleApiHeaders
	res, err := client.Do(req)

	if err != nil {
		// todo handle error
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// todo handle error
		return nil, err
	}

	// free mem-leak
	// todo discard data
	c.responseBody = string(body)
	return c, nil
}

func (c *GoogleApi) GetValue() float64 {
	return c.rateValue
}

func (c *GoogleApi) GetDate() string {
	return c.rateDate
}

func (c *GoogleApi) Latest(from string, to string, opt map[string]string) error {

	// todo cache layer
	_, err := c.RequestRate(from, to, opt)
	if err != nil {
		// todo handle error
		return err
	}
	validID := regexp.MustCompile(`knowledge-currency__tgt-input(.*)value="([1-9.]{0,10})" (.*)"`)
	stringMatches := validID.FindStringSubmatch(c.responseBody)
	// todo handle error
	c.rateValue, _ = strconv.ParseFloat(stringMatches[2], 64)
	return nil
}

func NewGoogleApi() *GoogleApi {
	r := &GoogleApi{apiKey: "12344"}
	return r
}
