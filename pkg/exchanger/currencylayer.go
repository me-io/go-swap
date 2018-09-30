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

type currencyLayerApi struct {
	attributes
}

var (
	currencyLayerApiUrl     = `https://apilayer.net/api/convert?access_key=%s&from=%s&to=%s&amount=1&format=1`
	currencyLayerApiHeaders = map[string][]string{
		`Accept`:          {`text/html,application/xhtml+xml,application/xml,application/json`},
		`Accept-Encoding`: {`text`},
	}
)

func (c *currencyLayerApi) requestRate(from string, to string, opt ...interface{}) (*currencyLayerApi, error) {

	// todo add option opt to add more headers or client configurations
	// free mem-leak
	// optimize for memory leak
	// todo optimize curl connection

	url := fmt.Sprintf(currencyLayerApiUrl, c.apiKey, from, to)
	req, _ := http.NewRequest("GET", url, nil)

	currencyLayerApiHeaders[`User-Agent`] = []string{c.userAgent}
	req.Header = currencyLayerApiHeaders

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

func (c *currencyLayerApi) GetValue() float64 {
	return c.rateValue
}

func (c *currencyLayerApi) GetDateTime() string {
	return c.rateDate.Format(time.RFC3339)
}

func (c *currencyLayerApi) GetExchangerName() string {
	return c.name
}

func (c *currencyLayerApi) Latest(from string, to string, opt ...interface{}) error {

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
	value := json.GetPath(`result`).
		MustFloat64()
	// todo handle error
	c.rateValue = value
	c.rateDate = time.Now()
	return nil
}

func NewCurrencyLayerApi(opt map[string]string) *currencyLayerApi {
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
		name:      `currencylayer`,
		Client:    client,
		apiKey:    opt[`apiKey`],
		userAgent: `Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:21.0) Gecko/20100101 Firefox/21.0`,
	}
	if opt[`userAgent`] != "" {
		attr.userAgent = opt[`userAgent`]
	}

	r := &currencyLayerApi{attr}
	return r
}
