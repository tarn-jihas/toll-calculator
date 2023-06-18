package main

import (
	"log"

	"github.com/keselj-strahinja/toll-calculator/aggregator/client"
)

const (
	kafkaTopic             = "obudata"
	httpAggregatorEndpoint = "http://127.0.0.1:3000/aggregate"
	grpcAggregatorEndpoint = "127.0.0.1:3001"
)

// Transport (HTTP, GRPC, Kafka) -> attach business logic to this transport
func main() {
	var (
		err error
		svc CalculatorServicer
	)

	svc = NewCalcService()
	svc = NewLogMiddleware(svc)
	// httpClient := client.NewHTTPClient(httpAggregatorEndpoint)
	grpcClient, err := client.NewGRPCClient(grpcAggregatorEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, grpcClient)

	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()
}
