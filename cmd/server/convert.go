package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	ex "github.com/me-io/go-swap/pkg/exchanger"
	"github.com/me-io/go-swap/pkg/swap"
	"math"
	"net/http"
	"time"
)

func (c *convertReqObj) Validate() error {
	// todo implement
	return nil
}

func (c convertReqObj) Hash() string {
	c.Amount = 1
	jsonBytes, _ := json.Marshal(c)
	md5Sum := md5.Sum(jsonBytes)
	return fmt.Sprintf("%x", md5Sum[:])
}

var Convert = func(w http.ResponseWriter, r *http.Request) {

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
	currencyCacheTime, _ := time.ParseDuration(convertReq.CacheTime)

	convertRes := &convertResObj{}
	if string(currencyCachedVal) == "" {
		Swap := swap.NewSwap()
		for _, v := range convertReq.Exchanger {

			var e ex.Exchanger
			opt := map[string]string{`userAgent`: v.UserAgent, `apiKey`: v.ApiKey}

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
			}
			Swap.AddExchanger(e)
		}
		Swap.Build()

		rate := Swap.Latest(convertReq.From + `/` + convertReq.To)

		convertRes.From = convertReq.From
		convertRes.To = convertReq.To
		convertRes.ExchangeValue = rate.GetValue()
		convertRes.DateTime = rate.GetDateTime()
		convertRes.ExchangerName = rate.GetExchangerName()

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
	}

	convertedAmount := math.Round(convertReq.Amount*convertRes.ExchangeValue*math.Pow10(decimalPoint)) / math.Pow10(decimalPoint)
	convertRes.ConvertedAmount = convertedAmount
	convertRes.OriginalAmount = convertReq.Amount

	currencyJsonVal, err := json.Marshal(convertRes)
	if err != nil {
		Logger.Panic(err)
	}

	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(currencyJsonVal)
}
