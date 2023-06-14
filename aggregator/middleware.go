package main

import (
	"time"

	"github.com/keselj-strahinja/toll-calculator/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleWare struct {
	next Aggregator
}

func NewLogMiddleware(next Aggregator) Aggregator {
	return &LogMiddleWare{
		next: next,
	}
}

func (l *LogMiddleWare) AggregateDistance(distance types.Distance) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(
			logrus.Fields{

				"took": time.Since(start),
				"err":  err,
			},
		).Info("aggregating")
	}(time.Now())

	err = l.next.AggregateDistance(distance)

	return
}
