package main

import (
	"flag"
	"fmt"
	"github.com/me-io/go-swap/pkg/cache"
	"github.com/me-io/go-swap/pkg/cache/memory"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
	"time"
)

var (
	//allowedExchanger = map[string]string{
	//	`google`: `googleApi`,
	//	`yahoo`:  `yahooApi`,
	//	`fixer`:  `fixer`,
	//}

	routes = map[string]func(w http.ResponseWriter, r *http.Request){
		`/favicon.ico`: favIcon,
		`/convert`:     Convert,
		`/`:            home,
	}
	_, filename, _, _ = runtime.Caller(0)
	sPath             = filepath.Dir(filename) + `/static/`
	Storage           cache.Storage
)

var favIcon = func(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, sPath+`favicon.ico`)
}

var home = func(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, sPath+`index.html`)
}

func init() {
	cacheDriver := flag.String("s", "memory", "Cache strategy (memory or redis)")
	flag.Parse()
	switch *cacheDriver {
	case `redis`:
		break
	default:
		Storage = memory.NewStorage()
	}
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

	fmt.Printf("Server Started @ %v:%d", host, port)

	err := server.ListenAndServe()
	log.Println(err.Error())
}
