package main

import "github.com/keselj-strahinja/toll-calculator/types"

type LoggingMiddleware struct {
	next DataProducer
}

func (l *LoggingMiddleware) ProduceData(data types.OBUData) error {

}
