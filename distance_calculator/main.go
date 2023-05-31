package main

import "log"

// type DistanceCalculator struct {
// 	consumer DataConsumer

const kafkaTopic = "obudata"

func main() {
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic)

	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()
}
