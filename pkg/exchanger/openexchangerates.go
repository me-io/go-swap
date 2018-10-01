package exchanger

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

type openExchangeRatesApi struct {
	attributes
}

var (
	openExchangeRatesApiUrl     = `https://openexchangerates.org/api/convert/1/%s/%s?app_id=%s`
	openExchangeRatesApiHeaders = map[string][]string{
		`Accept`:          {`text/html,application/xhtml+xml,application/xml,application/json`},
		`Accept-Encoding`: {`text`},
	}
)

func (c *openExchangeRatesApi) requestRate(from string, to string, opt ...interface{}) (*openExchangeRatesApi, error) {

	// todo add option opt to add more headers or client configurations
	// free mem-leak
	// optimize for memory leak
	// todo optimize curl connection

	url := fmt.Sprintf(openExchangeRatesApiUrl, c.apiKey, from, to)
	req, _ := http.NewRequest("GET", url, nil)

	openExchangeRatesApiHeaders[`User-Agent`] = []string{c.userAgent}
	req.Header = openExchangeRatesApiHeaders

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

// GetRateValue ... get exchange rate value
func (c *openExchangeRatesApi) GetRateValue() float64 {
	return c.rateValue
}

// GetRateDateTime ... return rate datetime
func (c *openExchangeRatesApi) GetRateDateTime() string {
	return c.rateDate.Format(time.RFC3339)
}

// GetExchangerName ... return exchanger name
func (c *openExchangeRatesApi) GetExchangerName() string {
	return c.name
}

// Latest ... populate latest exchange rate
func (c *openExchangeRatesApi) Latest(from string, to string, opt ...interface{}) error {

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

	json, err := simplejson.NewJson([]byte(c.responseBody))

	if err != nil {
		log.Print(err)
		return err
	}

	// opening price
	value := json.GetPath(`response`).
		MustFloat64()
	// todo handle error
	if value <= 0 {
		return fmt.Errorf(`error in retrieving exhcange rate is 0`)
	}
	c.rateValue = value
	c.rateDate = time.Now()
	return nil
}

// NewOpenExchangeRatesApi ... return new instance of openExchangeRatesApi
func NewOpenExchangeRatesApi(opt map[string]string) *openExchangeRatesApi {
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

	attr := attributes{
		name:      `openexchangerates`,
		Client:    client,
		apiKey:    opt[`apiKey`],
		userAgent: `Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:21.0) Gecko/20100101 Firefox/21.0`,
	}
	if opt[`userAgent`] != "" {
		attr.userAgent = opt[`userAgent`]
	}

	r := &openExchangeRatesApi{attr}
	return r
}
