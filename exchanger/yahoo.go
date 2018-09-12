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

type YahooApi struct {
	apiKey       string
	responseBody string
	rateValue    float64
	rateDate     string
}

// ref @link https://github.com/florianv/exchanger/blob/master/src/Service/Yahoo.php
// var YahooApiUrl = `https://query.yahooapis.com/v1/public/yql?q=%s&env=store://datatables.org/alltableswithkeys&format=json`
var YahooApiUrl = `https://quote.yahoo.com/d/quotes.csv?s=%s%s=X`

var YahooApiHeaders = map[string][]string{
	`Accept`:     {`text/html`},
	`User-Agent`: {`Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:21.0) Gecko/20100101 Firefox/21.0`},
}

func (c *YahooApi) RequestRate(from string, to string, opt map[string]string) (*YahooApi, error) {

	// todo add option opt to add more headers or client configurations
	// free mem-leak
	// optimize for memory leak
	// todo optimize curl connection
	keepAliveTimeout := 600 * time.Second
	timeout := 2 * time.Second
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

	// query := fmt.Sprintf(`select+*+from+yahoo.finance.xchange+where+pair+in+("%s%s")`, from, to)
	url := fmt.Sprintf(YahooApiUrl, from, to)
	println(url)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header = YahooApiHeaders
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

func (c *YahooApi) GetValue() float64 {
	return c.rateValue
}

func (c *YahooApi) GetDate() string {
	return c.rateDate
}

func (c *YahooApi) Latest(from string, to string, opt map[string]string) error {

	// todo cache layer
	_, err := c.RequestRate(from, to, opt)
	if err != nil {
		// todo handle error
		return err
	}
	println(c.responseBody)
	println(c.responseBody)
	println(c.responseBody)
	println(c.responseBody)
	validID := regexp.MustCompile(`knowledge-currency__tgt-input(.*)value="([1-9.]{0,10})" (.*)"`)
	stringMatches := validID.FindStringSubmatch(c.responseBody)
	// todo handle error
	c.rateValue, _ = strconv.ParseFloat(stringMatches[2], 64)
	return nil
}

func NewYahooApi() *YahooApi {
	r := &YahooApi{apiKey: "12344"}
	return r
}
