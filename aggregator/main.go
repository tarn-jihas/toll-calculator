package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/keselj-strahinja/toll-calculator/types"
)

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "the listen address of the service of the HTTP Transport handler server")
	flag.Parse()

	store := NewMemoryStore()
	var (
		svc = NewInvoiceAggregator(store)
	)

	makeHTTPTransport(*listenAddr, svc)
}

func makeHTTPTransport(listenaddr string, svc Aggregator) {
	fmt.Println("Http transport running on port", listenaddr)
	http.HandleFunc("/aggregate", handleAggregate(svc))
	http.ListenAndServe(listenaddr, nil)
}

func handleAggregate(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
