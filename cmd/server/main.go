package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	//allowedExchanger = map[string]string{
	//	`google`: `googleApi`,
	//	`yahoo`:  `yahooApi`,
	//	`fixer`:  `fixer`,
	//}

	routes = map[string]func(w http.ResponseWriter, r *http.Request){
		`/convert`: Convert,
		`/`:        etc,
	}
)

var etc = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("etc")
}

func main() {

	// handle routers
	for k, v := range routes {
		http.HandleFunc(k, v)
	}

	// todo
	// port and config
	// cache
	go serveHTTP(`0.0.0.0`, 5000)
	select {}
}

func serveHTTP(host string, port int) {

	mux := http.NewServeMux()
	for k, v := range routes {
		mux.HandleFunc(k, v)
	}

	addr := fmt.Sprintf("%v:%d", host, port)
	server := &http.Server{
		Addr:           addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println(`Server Started`)

	err := server.ListenAndServe()
	log.Println(err.Error())
}
