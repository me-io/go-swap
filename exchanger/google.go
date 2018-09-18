package exchanger

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

type googleApi struct {
	attributes
}

// ref @link https://github.com/florianv/exchanger/blob/master/src/Service/Google.php
// example : https://www.google.com/search?q=1+USD+to+USD&ncr=1
// example : https://www.google.com/search?q=1+USD+to+EGP&ncr=1
// example : https://www.google.com/search?q=1+USD+to+AED&ncr=1
var (
	googleApiUrl     = `https://www.google.com/search?q=1+%s+to+%s&ncr=1`
	googleApiHeaders = map[string][]string{
		`Accept`:     {`text/html`},
		`User-Agent`: {`Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:21.0) Gecko/20100101 Firefox/21.0`},
	}
)

func (c *googleApi) requestRate(from string, to string, opt ...interface{}) (*googleApi, error) {

	// todo add option opt to add more headers or client configurations
	// free mem-leak
	// optimize for memory leak
	// todo optimize curl connection

	// format the url and replace currency
	url := fmt.Sprintf(googleApiUrl, from, to)
	// prepare the request
	req, _ := http.NewRequest("GET", url, nil)
	// assign the request headers
	req.Header = googleApiHeaders
	// execute the request
	res, err := c.Client.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// free mem-leak
	// todo discard data
	c.responseBody = string(body)
	return c, nil
}

func (c *googleApi) GetValue() float64 {
	return c.rateValue
}

func (c *googleApi) GetDate() string {
	return c.rateDate
}

func (c *googleApi) GetExchangerName() string {
	return c.name
}

func (c *googleApi) Latest(from string, to string, opt ...interface{}) error {

	// todo cache layer
	_, err := c.requestRate(from, to, opt)
	if err != nil {
		log.Print(err)
		return err
	}

	// if from currency is same as converted currency return value of 1
	if from == to {
		c.rateValue = 1
		return nil
	}

	validID := regexp.MustCompile(`(?s)knowledge-currency__tgt-input(.*)value="([0-9.]{0,12})"(.*)"`)
	stringMatches := validID.FindStringSubmatch(c.responseBody)

	c.rateValue, err = strconv.ParseFloat(stringMatches[2], 64)

	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func NewGoogleApi(opt map[string]string) *googleApi {

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

	r := &googleApi{attributes{name: `googleApi`, Client: client,}}

	return r
}
