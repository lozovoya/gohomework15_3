package main

import (
	"encoding/json"
	"github.com/lozovoya/gohomework15_3/pkg/cities"
	"github.com/lozovoya/gohomework15_3/pkg/middleware/ErrCatcherMd"
	"github.com/lozovoya/gohomework15_3/pkg/middleware/logger"
	"github.com/lozovoya/gohomework15_3/pkg/remux"
	"log"
	"net"
	"net/http"
	"os"
)

const defaultPort = "9999"
const defaultHost = "0.0.0.0"

func main() {

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = defaultPort
	}

	host, ok := os.LookupEnv("HOST")
	if !ok {
		host = defaultHost
	}

	if err := execute(net.JoinHostPort(host, port)); err != nil {
		log.Println(err)
		os.Exit(1)
	}

}

func execute(addr string) error {

	rmux := remux.New()
	catcherMd := ErrCatcherMd.ErrCatcher
	loggerMd := logger.Logger
	rmux.RegisterPlain(remux.GET, "/cities", http.HandlerFunc(citiesHandler), loggerMd, catcherMd)
	rmux.RegisterPlain(remux.GET, "/fail", http.HandlerFunc(errHandler), loggerMd, catcherMd)

	log.Fatal(http.ListenAndServe(addr, rmux))

	return nil
}

func citiesHandler(w http.ResponseWriter, r *http.Request) {

	cities := cities.AllCities()

	data, err := json.Marshal(cities)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(data)
	return
}

func errHandler(w http.ResponseWriter, r *http.Request) {
	remux.ExecPanic()
	return
}
