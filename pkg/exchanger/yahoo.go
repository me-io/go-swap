package exchanger

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"math"
	"net"
	"net/http"
	"time"
)

type yahooApi struct {
	attributes
}

var (
	yahooApiUrl     = `https://query1.finance.yahoo.com/v8/finance/chart/%s%s=X?region=US&lang=en-US&includePrePost=false&interval=1d&range=1d&corsDomain=finance.yahoo.com&.tsrc=finance`
	yahooApiHeaders = map[string][]string{
		`Accept`:          {`text/html,application/xhtml+xml,application/xml,application/json`},
		`Accept-Encoding`: {`text`},
	}
)

func (c *yahooApi) requestRate(from string, to string, opt ...interface{}) (*yahooApi, error) {

	// todo add option opt to add more headers or client configurations
	// free mem-leak
	// optimize for memory leak
	// todo optimize curl connection

	url := fmt.Sprintf(yahooApiUrl, from, to)
	req, _ := http.NewRequest("GET", url, nil)

	yahooApiHeaders[`User-Agent`] = []string{c.userAgent}
	req.Header = yahooApiHeaders

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
func (c *yahooApi) GetRateValue() float64 {
	return c.rateValue
}

// GetRateDateTime ... return rate datetime
func (c *yahooApi) GetRateDateTime() string {
	return c.rateDate.Format(time.RFC3339)
}

// GetExchangerName ... return exchanger name
func (c *yahooApi) GetExchangerName() string {
	return c.name
}

// Latest ... populate latest exchange rate
func (c *yahooApi) Latest(from string, to string, opt ...interface{}) error {

	_, err := c.requestRate(from, to, opt)
	if err != nil {
		log.Print(err)
		return err
	}

	json, err := simplejson.NewJson([]byte(c.responseBody))

	if err != nil {
		log.Print(err)
		return err
	}

	// opening price
	value := json.GetPath(`chart`, `result`).
		GetIndex(0).
		//GetPath(`indicators`, `adjclose`).
		//GetIndex(0).
		//GetPath(`adjclose`).
		//GetIndex(0).
		GetPath(`indicators`, `quote`).
		GetIndex(0).
		GetPath(`open`).
		GetIndex(0).
		MustFloat64()
	// todo handle error
	c.rateValue = math.Round(value*1000000) / 1000000
	c.rateDate = time.Now()
	return nil
}

// NewYahooApi ... return new instance of yahooApi
func NewYahooApi(opt map[string]string) *yahooApi {
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
		name:      `yahoo`,
		Client:    client,
		userAgent: `Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:21.0) Gecko/20100101 Firefox/21.0`,
	}
	if opt[`userAgent`] != "" {
		attr.userAgent = opt[`userAgent`]
	}

	r := &yahooApi{attr}
	return r
}
