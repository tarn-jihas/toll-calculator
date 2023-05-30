package main

import (
	"time"

	"github.com/keselj-strahinja/toll-calculator/types"
	"github.com/sirupsen/logrus"
)

type LoggingMiddleware struct {
	next DataProducer
}

func NewLogMiddleware(next DataProducer) *LoggingMiddleware {
	return &LoggingMiddleware{
		next: next,
	}
}

func (l *LoggingMiddleware) ProduceData(data types.OBUData) error {
	defer func(start time.Time) {
		logrus.WithFields(
			logrus.Fields{
				"obuID": data.OBUID,
				"lat":   data.Lat,
				"lon":   data.Lon,
				"took":  time.Since(start),
			},
		).Info("producing to kafka")
	}(time.Now())

	return l.next.ProduceData(data)
}
