package main

import (
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

	cities := cities.NewService()
	cities.AddCity("Moscow", "Russia", 14_000_000, 1147)
	cities.AddCity("Kiev", "Ukraine", 4_000_000, 882)
	cities.AddCity("Moscow", "Belarus", 2_000_000, 1067)
	cities.AddCity("Vilnius", "Lietuva", 500_000, 1323)

	rmux := remux.New()
	catcherMd := ErrCatcherMd.ErrCatcher
	loggerMd := logger.Logger
	rmux.RegisterPlain(remux.GET, "/cities", http.HandlerFunc(citiesHandler), loggerMd, catcherMd)
	rmux.RegisterPlain(remux.GET, "/fail", http.HandlerFunc(citiesHandler2), loggerMd, catcherMd)

	log.Fatal(http.ListenAndServe(addr, rmux))

	return nil
}

func citiesHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
	return
}

func citiesHandler2(w http.ResponseWriter, r *http.Request) {
	remux.ExecPanic()
	return
}

//func bandsHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodGet {
//		http.Error(w, http.ErrNotSupported.Error(), http.StatusBadRequest)
//		return
//	}
//	bands := core.Bands()
//	data, err := json.Marshal(bands)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	w.Write(data)
//}
