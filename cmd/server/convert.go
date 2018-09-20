package main

import (
	"encoding/json"
	ex "github.com/me-io/go-swap/pkg/exchanger"
	"github.com/me-io/go-swap/pkg/swap"
	"net/http"
)

func (c convertReqObj) Validate() error {
	return nil
}

var Convert = func(w http.ResponseWriter, r *http.Request) {

	convertReq := &convertReqObj{}

	if err := json.
		NewDecoder(r.Body).
		Decode(convertReq); err != nil {
		panic(err)
	}

	if err := convertReq.Validate(); err != nil {
		panic(err)
	}

	//fmt.Println(convertReq)

	Swap := swap.NewSwap()
	for _, v := range convertReq.Exchanger {
		//fmt.Println(k)
		var e ex.Exchanger

		switch v.Name {
		case `google`:
			e = ex.NewGoogleApi(nil)
		case `yahoo`:
			e = ex.NewYahooApi(nil)
		}
		Swap.AddExchanger(e)
	}
	Swap.Build()

	rate := Swap.Latest(convertReq.From + `/` + convertReq.To)

	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	convertRes := convertResObj{
		Value:         rate.GetValue(),
		Date:          rate.GetDate(),
		ExchangerName: rate.GetExchangerName(),
	}
	resJson, err := json.Marshal(convertRes)
	if err != nil {
		// todo handle error
	}
	w.Write(resJson)
}
