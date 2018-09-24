package main

import (
	"flag"
	"fmt"
	"github.com/me-io/go-swap/pkg/cache"
	"github.com/me-io/go-swap/pkg/cache/memory"
	"github.com/me-io/go-swap/pkg/cache/redis"
	"github.com/op/go-logging"
	"net/http"
	"os"
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
	host        *string
	port        *int
	cacheDriver *string
	Storage     cache.Storage

	Logger = logging.MustGetLogger("go-swap-server")

	format = logging.MustStringFormatter(
		`%{color}%{time:2006-01-02T15:04:05.999999} %{shortfunc} ▶ %{level:.8s} %{id:03x}%{color:reset} %{message}`,
	)

	routes = map[string]func(w http.ResponseWriter, r *http.Request){
		`/favicon.ico`: favIcon,
		`/convert`:     Convert,
		`/`:            home,
	}

	_, filename, _, _ = runtime.Caller(0)
	sPath             = filepath.Dir(filename) + `/static/`
)

var favIcon = func(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, sPath+`favicon.ico`)
}

var home = func(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, sPath+`index.html`)
}
// Password is just an example type implementing the Redactor interface. Any
// time this is logged, the Redacted() function will be called.
type Password string

func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}

func init() {
	// Logging
	backendStderr := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatted := logging.NewBackendFormatter(backendStderr, format)
	// Only DEBUG and more severe messages should be sent to backend1
	backendLevelFormatted := logging.AddModuleLevel(backendFormatted)
	backendLevelFormatted.SetLevel(logging.DEBUG, "")
	// Set the backend to be used.
	logging.SetBackend(backendLevelFormatted)

	// Caching
	host = flag.String(`h`, `0.0.0.0`, `HTTP Server Hostname or IP`)
	port = flag.Int(`p`, 5000, `HTTP Server Port`)
	cacheDriver = flag.String(`s`, `memory`, `Cache strategy (memory or redis)`)

	flag.Parse()

	var err error

	switch *cacheDriver {
	case `redis`:
		if Storage, err = redis.NewStorage(os.Getenv(`REDIS_URL`)); err != nil {
			panic(err)
		}
		break
	default:
		Storage = memory.NewStorage()
	}

	//Logger.Debugf("debug %s", Password(`secret`))
	//Logger.Info("info")
	//Logger.Notice("notice")
	//Logger.Warning("warning")
	//Logger.Error("err")
	//Logger.Critical("crit")

}

func main() {

	Logger.Debugf("host %s", *host)
	Logger.Debugf("port %d", *port)
	Logger.Debugf("cacheDriver %s", *cacheDriver)
	Logger.Warningf("cacheDriver %s", *cacheDriver)

	// handle routers
	for k, v := range routes {
		http.HandleFunc(k, v)
	}

	// todo
	// port and config
	go serveHTTP(*host, *port)
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

	Logger.Infof("Server Started @ %v:%d", host, port)

	err := server.ListenAndServe()
	Logger.Error(err.Error())
}
