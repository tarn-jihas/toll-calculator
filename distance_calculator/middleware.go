package main

import (
	"time"

	"github.com/keselj-strahinja/toll-calculator/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleWare struct {
	next CalculatorServicer
}

func NewLogMiddleware(next CalculatorServicer) CalculatorServicer {
	return &LogMiddleWare{
		next: next,
	}
}

func (m *LogMiddleWare) CalculateDistance(data types.OBUData) (dist float64, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err":  err,
			"dist": dist,
		}).Info("calculate distance")
	}(time.Now())
	dist, err = m.next.CalculateDistance(data)
	return 
}
