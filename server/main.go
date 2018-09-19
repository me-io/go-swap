package main

import (
	"fmt"
	"net/http"
)

var (
	allowedExchanger = map[string]string{
		`google`: `googleApi`,
		`yahoo`:  `yahooApi`,
		`fixer`:  `fixer`,
	}

	routes = map[string]func(w http.ResponseWriter, r *http.Request){
		`/convert`: convert,
		`/`:        etc,
	}
)

var etc = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("etc")
}

var convert = func(w http.ResponseWriter, r *http.Request) {
	payload := `{
  "amount": 1,
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
}`
	fmt.Println(payload)

}

func main() {

	// handle routers
	for k, v := range routes {
		http.HandleFunc(k, v)
	}

	fmt.Println(`server started`)
	// todo
	// port and config
	// cache
	http.ListenAndServe(":3003", nil)
}
