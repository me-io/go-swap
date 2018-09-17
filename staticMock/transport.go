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
	default:

	}

	response.Body = ioutil.NopCloser(strings.NewReader(responseBody))
	return response, nil
}
