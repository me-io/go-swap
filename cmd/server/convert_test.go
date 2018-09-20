package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

type testResponseWriter struct {
}

var testResponse string

func (c testResponseWriter) Header() http.Header {
	return map[string][]string{}
}

func (c testResponseWriter) Write(i []byte) (int, error) {
	testResponse = string(i)
	return 0, nil
}

func (c testResponseWriter) WriteHeader(statusCode int) {
}

func TestConvertObj_Convert(t *testing.T) {
	payloadArr := []string{
		`{
  "amount": 4.5,
  "exchanger": [
    {
      "name": "google",
      "userAgent": "firefox"
    }
  ],
  "from": "USD",
  "to": "AED"
}`,
		`{
  "amount": 4.5,
  "exchanger": [
    {
      "name": "yahoo",
      "userAgent": "firefox"
    }
  ],
  "from": "USD",
  "to": "AED"
}`,
		`{
  "amount": 4.5,
  "exchanger": [
    {
      "name": "google",
      "userAgent": "firefox"
    },
    {
      "name": "yahoo",
      "userAgent": "Chrome"
    },
    {
      "name": "currencyLayer",
      "apiKey": "12312",
      "userAgent": "currencyLayer Chrome"
    },
    {
      "name": "fixer",
      "apiKey": "12312",
      "userAgent": "currencyLayer fixer"
    }
  ],
  "from": "USD",
  "to": "AED"
}`,
	}

	expected := map[int]string{
		0: "google",
		1: "yahoo",
		2: "google",
	}

	for k, payload := range payloadArr {
		// mock the payload
		bytePayload := []byte(payload)
		bytePayloadReader := bytes.NewReader(bytePayload)

		w := testResponseWriter{}

		r := &http.Request{}
		r.Body = ioutil.NopCloser(bytePayloadReader)

		Convert(w, r)
		assert.Contains(t, testResponse, expected[k])
	}
}