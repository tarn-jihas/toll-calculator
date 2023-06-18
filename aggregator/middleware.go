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
		).Info("aggregating distance")
	}(time.Now())

	err = l.next.AggregateDistance(distance)

	return
}

func (l *LogMiddleWare) CalculateInvoice(id int) (invoice *types.Invoice, err error) {
	defer func(start time.Time) {
		var (
			distance float64
			amount   float64
		)
		if invoice != nil {
			distance = invoice.TotalDistance
			amount = invoice.TotalAmount
		}
		logrus.WithFields(
			logrus.Fields{
				"took":     time.Since(start),
				"err":      err,
				"OBUID":    id,
				"amount":   amount,
				"distance": distance,
			},
		).Info("Fetching Invoice")
	}(time.Now())

	return l.next.CalculateInvoice(id)

}
