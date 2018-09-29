package staticMock

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
)

type mT struct{}

// NewTestMT ... return new RoundTripper for mocking http response
func NewTestMT() http.RoundTripper {
	return &mT{}
}

// Implement http.RoundTripper
func (t *mT) RoundTrip(req *http.Request) (*http.Response, error) {
	// Create mocked http.Response
	response := &http.Response{
		Header:     make(http.Header),
		Request:    req,
		StatusCode: http.StatusOK,
	}
	// url := req.URL.RequestURI()
	host := req.URL.Host

	_, filename, _, _ := runtime.Caller(0)
	tPath := filepath.Dir(filename)

	responseBody := ``
	response.Header.Set("Content-Type", "application/json")

	switch {
	case host == `www.google.com`:
		response.Header.Set("Content-Type", "text/html")
		fc, _ := ioutil.ReadFile(tPath + `/google_html_aed_usd.html`)
		responseBody = string(fc)
		break
	case host == `query1.finance.yahoo.com`:
		fp, _ := filepath.Abs(tPath + `/yahoo_json_aed_usd.json`)
		fc, _ := ioutil.ReadFile(fp)
		responseBody = string(fc)
		break
	case host == `apilayer.net`:
		fp, _ := filepath.Abs(tPath + `/currencylayer_json_aed_usd.json`)
		fc, _ := ioutil.ReadFile(fp)
		responseBody = string(fc)
		break
	case host == `data.fixer.io`:
		fp, _ := filepath.Abs(tPath + `/fixer_json_aed_usd.json`)
		fc, _ := ioutil.ReadFile(fp)
		responseBody = string(fc)
		break
	case host == `forex.1forge.com`:
		fp, _ := filepath.Abs(tPath + `/1forge_json_aed_usd.json`)
		fc, _ := ioutil.ReadFile(fp)
		responseBody = string(fc)
		break
	case host == `themoneyconverter.com`:
		fp, _ := filepath.Abs(tPath + `/themoneyconverter_html_aed_usd.html`)
		fc, _ := ioutil.ReadFile(fp)
		responseBody = string(fc)
		break
	case host == `openexchangerates.org`:
		fp, _ := filepath.Abs(tPath + `/openexchangerates_json_aed_usd.json`)
		fc, _ := ioutil.ReadFile(fp)
		responseBody = string(fc)
		break
	default:

	}

	response.Body = ioutil.NopCloser(strings.NewReader(responseBody))
	return response, nil
}
