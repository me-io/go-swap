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
	"strconv"
	"time"
)

var (
	host        string
	port        int
	cacheDriver string
	redisUrl    string
	Storage     cache.Storage
	Logger      = logging.MustGetLogger("go-swap-server")

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
	host = GetEnv(`h`, `0.0.0.0`)
	port, _ = strconv.Atoi(GetEnv(`p`, `5000`))
	cacheDriver = GetEnv(`cache`, `memory`)
	redisUrl = GetEnv(`REDIS_URL`, ``)

	flag.Parse()

	var err error

	switch cacheDriver {
	case `redis`:
		if Storage, err = redis.NewStorage(redisUrl); err != nil {
			Logger.Panic(err)
		}
		break
	default:
		Storage = memory.NewStorage()
	}

}

func main() {

	Logger.Infof("host %s", host)
	Logger.Infof("port %d", port)
	Logger.Infof("cacheDriver %s", cacheDriver)
	Logger.Infof("REDIS_URL %s", redisUrl)

	// handle routers
	for k, v := range routes {
		http.HandleFunc(k, v)
	}

	go serveHTTP(host, port)
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
