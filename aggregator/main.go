package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/keselj-strahinja/toll-calculator/types"
	"google.golang.org/grpc"
)

func main() {
	httpListenAddr := flag.String("httpAddr", ":3000", "the listen address of the service of the HTTP Transport handler server")
	grpcListenAddr := flag.String("grpcAddr", ":3001", "the listen address of the service of the GRPC Transport handler server")
	flag.Parse()

	var (
		store = NewMemoryStore()
		svc   = NewInvoiceAggregator(store)
	)

	svc = NewLogMiddleware(svc)
	go makeGRPCTransport(*grpcListenAddr, svc)
	makeHTTPTransport(*httpListenAddr, svc)
}

func makeGRPCTransport(listenAddr string, svc Aggregator) error {
	fmt.Println("grpc transport running on port", listenAddr)
	// Make a TCP Listener
	conn, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	defer func() {
		conn.Close()
	}()
	// Make a new GRPC native server with options
	server := grpc.NewServer([]grpc.ServerOption{}...)
	// Register our GRPC server implememtation to the GRPC package
	types.RegisterAggregatorServer(server, NewGRPCAggregatorServer(svc))

	if err := server.Serve(conn); err != nil {
		return err
	}

	return nil
}

func makeHTTPTransport(listenAddr string, svc Aggregator) {
	fmt.Println("Http transport running on port", listenAddr)
	http.HandleFunc("/aggregate", handleAggregate(svc))
	http.HandleFunc("/invoice", handleGetInvoice(svc))
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatalf("Http server failed to start: %v", err)
	}
}

func handleAggregate(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		if err := svc.AggregateDistance(distance); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
	}

}

func handleGetInvoice(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		values, ok := r.URL.Query()["obu"]
		if !ok || values[0] == "" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"err": "Missing OBU id"})
			return
		}

		obuID, err := strconv.Atoi(values[0])
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"err": "invalid OBU id"})
			return
		}

		invoice, err := svc.CalculateInvoice(obuID)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"err": err.Error()})
			return

		}

		writeJSON(w, http.StatusOK, invoice)
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
