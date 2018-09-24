package main

import (
	"encoding/json"
	ex "github.com/me-io/go-swap/pkg/exchanger"
	"github.com/me-io/go-swap/pkg/swap"
	"math"
	"net/http"
	"time"
)

func (c convertReqObj) Validate() error {
	// todo implement
	return nil
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

	currencyKey := convertReq.From + `/` + convertReq.To
	currencyCachedVal := Storage.Get(currencyKey)
	currencyCacheTime, _ := time.ParseDuration(convertReq.CacheTime)

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

		rate := Swap.Latest(currencyKey)
		amount := math.Round(convertReq.Amount*rate.GetValue()*math.Pow10(decimalPoint)) / math.Pow10(decimalPoint)

		convertRes := convertResObj{
			Amount:        amount,
			Value:         rate.GetValue(),
			Date:          rate.GetDate(),
			ExchangerName: rate.GetExchangerName(),
		}

		var err error
		if currencyCachedVal, err = json.Marshal(convertRes); err != nil {
			Logger.Panic(err)
		}
		Storage.Set(currencyKey, currencyCachedVal, currencyCacheTime)
		w.Header().Set("X-Cache", "Miss")
	} else {
		// get from cache
		w.Header().Set("X-Cache", "Hit")
	}

	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(currencyCachedVal)
}
