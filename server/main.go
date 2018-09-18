package main

import (
	"fmt"
	"net/http"
)

var allowedExchanger = map[string]string{
	`google`: `googleApi`,
	`yahoo`:  `yahooApi`,
	`fixer`:  `fixer`,
}

func main() {

	http.HandleFunc(`/convert`, convertOne)
	http.HandleFunc(`/convert-multi`, convertMultiWithSwap)

	fmt.Println(`server started`)
	// todo
	// port and config
	// cache
	http.ListenAndServe(":3003", nil)
}

var convertOne = func(w http.ResponseWriter, r *http.Request) {

}

var convertMultiWithSwap = func(w http.ResponseWriter, r *http.Request) {

}
