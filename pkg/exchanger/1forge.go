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

type oneForgeApi struct {
	attributes
}

var (
	oneForgeApiUrl     = `https://forex.1forge.com/%s/convert?from=%s&to=%s&quantity=1&api_key=%s`
	oneForgeApiHeaders = map[string][]string{
		`Accept`:          {`text/html,application/xhtml+xml,application/xml,application/json`},
		`Accept-Encoding`: {`text`},
	}
)

func (c *oneForgeApi) requestRate(from string, to string, opt ...interface{}) (*oneForgeApi, error) {

	// todo add option opt to add more headers or client configurations
	// free mem-leak
	// optimize for memory leak
	// todo optimize curl connection

	url := fmt.Sprintf(oneForgeApiUrl, c.apiVersion, from, to, c.apiKey)
	req, _ := http.NewRequest("GET", url, nil)

	oneForgeApiHeaders[`User-Agent`] = []string{c.userAgent}
	req.Header = oneForgeApiHeaders

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

func (c *oneForgeApi) GetValue() float64 {
	return c.rateValue
}

func (c *oneForgeApi) GetDateTime() string {
	return c.rateDate.Format(time.RFC3339)
}

func (c *oneForgeApi) GetExchangerName() string {
	return c.name
}

func (c *oneForgeApi) Latest(from string, to string, opt ...interface{}) error {

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
	value := json.GetPath(`value`).
		MustFloat64()
	// todo handle error
	c.rateValue = value
	c.rateDate = time.Now()
	return nil
}

func NewOneForgeApi(opt map[string]string) *oneForgeApi {
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
		name:       `1forge`,
		Client:     client,
		apiVersion: `1.0.3`,
		apiKey:     opt[`apiKey`],
		userAgent:  `Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:21.0) Gecko/20100101 Firefox/21.0`,
	}

	if opt[`userAgent`] != "" {
		attr.userAgent = opt[`userAgent`]
	}

	if opt[`apiVersion`] != "" {
		attr.apiVersion = opt[`apiVersion`]
	}

	r := &oneForgeApi{attr}
	return r
}
