package main

import (
	"log"

	"github.com/keselj-strahinja/toll-calculator/aggregator/client"
)

const (
	kafkaTopic         = "obudata"
	aggregatorEndpoint = "http://127.0.0.1:3000/aggregate"
)

// Transport (HTTP, GRPC, Kafka) -> attach business logic to this transport
func main() {
	var (
		err error
		svc CalculatorServicer
	)

	svc = NewCalcService()
	svc = NewLogMiddleware(svc)

	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, client.NewClient(aggregatorEndpoint))

	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()
}
