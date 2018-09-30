package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/go-ozzo/ozzo-validation"
	ex "github.com/me-io/go-swap/pkg/exchanger"
	"github.com/me-io/go-swap/pkg/swap"
	"io/ioutil"
	"math"
	"net/http"
	"time"
)

// Validate ... Validation function for convertReqObj
func (c *convertReqObj) Validate() error {

	return validation.ValidateStruct(c,
		validation.Field(&c.Amount, validation.Required),
		validation.Field(&c.From, validation.Required, validation.In(ex.CurrencyListArr...)),
		validation.Field(&c.To, validation.Required, validation.In(ex.CurrencyListArr...)),
		validation.Field(&c.Exchanger, validation.Required),
	)

	//if ex.CurrencyList[c.To] == "" || ex.CurrencyList[c.From] == "" {
	//	return fmt.Errorf("currency %s or %s is not supported", c.From, c.To)
	//}
}

// Hash ... return md5 string hash of the convertReqObj with 1 Unit Amount to cache the rate only for 1 Unit Amount
func (c convertReqObj) Hash() string {
	// hash exchange key only with 1 Unit value
	c.Amount = 1
	jsonBytes, _ := json.Marshal(c)
	md5Sum := md5.Sum(jsonBytes)
	return fmt.Sprintf("%x", md5Sum[:])
}

// Convert ... Main convert function attached to the router handler
var Convert = func(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		ConvertPost(w, r)
	}
	if r.Method == "GET" {
		ConvertGet(w, r)
	}
}

// ConvertGet ... handle GET request and simulate payload from get query params to ConvertPost function
var ConvertGet = func(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	apiKey := query.Get("apiKey")
	exchanger := query.Get("exchanger")
	amount := query.Get("amount")
	from := query.Get("from")
	to := query.Get("to")
	cacheTime := query.Get("cacheTime")

	payload := fmt.Sprintf(`{
  "amount": %s,
  "exchanger": [
    {
      "name": "%s",
      "apiKey": "%s"
    }
  ],
  "from": "%s",
  "to": "%s",
  "cacheTime":"%s"
}`, amount, exchanger, apiKey, from, to, cacheTime)

	bytePayload := []byte(payload)
	bytePayloadReader := bytes.NewReader(bytePayload)

	r.Body = ioutil.NopCloser(bytePayloadReader)
	ConvertPost(w, r)
}

// ConvertPost ... handle POST request, build Swap object, get and cache the currency exchange rate and amount
var ConvertPost = func(w http.ResponseWriter, r *http.Request) {

	convertReq := &convertReqObj{}

	if err := json.
		NewDecoder(r.Body).
		Decode(convertReq); err != nil {
		Logger.Panic(err)
	}

	if err := convertReq.Validate(); err != nil {
		Logger.Panic(err)
	}

	decimalPoint := convertReq.DecimalPoints
	if decimalPoint == 0 {
		decimalPoint = 4
	}

	currencyCacheKey := convertReq.Hash()

	currencyCachedVal := Storage.Get(currencyCacheKey)
	// default cache time
	if convertReq.CacheTime == "" {
		convertReq.CacheTime = "120s"
	}
	currencyCacheTime, _ := time.ParseDuration(convertReq.CacheTime)

	convertRes := &convertResObj{}
	if string(currencyCachedVal) == "" {
		Swap := swap.NewSwap()
		for _, v := range convertReq.Exchanger {

			var e ex.Exchanger
			opt := map[string]string{`userAgent`: v.UserAgent, `apiKey`: v.ApiKey, `apiVersion`: v.ApiVersion}

			switch v.Name {
			case `google`:
				e = ex.NewGoogleApi(opt)
				break
			case `yahoo`:
				e = ex.NewYahooApi(opt)
				break
			case `currencylayer`:
				e = ex.NewCurrencyLayerApi(opt)
				break
			case `fixer`:
				e = ex.NewFixerApi(opt)
				break
			case `1forge`:
				e = ex.NewOneForgeApi(opt)
				break
			case `themoneyconverter`:
				e = ex.NewTheMoneyConverterApi(opt)
				break
			case `openexchangerates`:
				e = ex.NewOpenExchangeRatesApi(opt)
				break
			}
			Swap.AddExchanger(e)
		}
		Swap.Build()

		rate := Swap.Latest(convertReq.From + `/` + convertReq.To)

		convertRes.From = convertReq.From
		convertRes.To = convertReq.To
		convertRes.ExchangeValue = rate.GetRateValue()
		convertRes.RateDateTime = rate.GetRateDateTime()
		convertRes.ExchangerName = rate.GetExchangerName()
		convertRes.RateFromCache = false

		var err error
		if currencyCachedVal, err = json.Marshal(convertRes); err != nil {
			Logger.Panic(err)
		}
		Storage.Set(currencyCacheKey, currencyCachedVal, currencyCacheTime)
		w.Header().Set("X-Cache", "Miss")
	} else {
		// get from cache
		w.Header().Set("X-Cache", "Hit")
		json.Unmarshal(currencyCachedVal, &convertRes)
		convertRes.RateFromCache = true
	}

	convertedAmount := math.Round(convertReq.Amount*convertRes.ExchangeValue*math.Pow10(decimalPoint)) / math.Pow10(decimalPoint)
	convertRes.ConvertedAmount = convertedAmount
	convertRes.OriginalAmount = convertReq.Amount

	// formatted message like "1 USD is worth 3.675 AED"
	convertRes.ConvertedText = fmt.Sprintf("%g %s is worth %g %s", convertRes.OriginalAmount, convertRes.From, convertRes.ConvertedAmount, convertRes.To)

	currencyJsonVal, err := json.Marshal(convertRes)
	if err != nil {
		Logger.Panic(err)
	}

	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(currencyJsonVal)
}
