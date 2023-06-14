package main

import (
	"math"

	"github.com/keselj-strahinja/toll-calculator/types"
)

// we like to end our interfaces with (er)
type CalculatorServicer interface {
	CalculateDistance(types.OBUData) (float64, error)
}

type CalculatorService struct {
	prevPoint []float64
}

func NewCalcService() CalculatorServicer {
	return &CalculatorService{}
}

func (s *CalculatorService) CalculateDistance(data types.OBUData) (float64, error) {
	var distance = 0.0

	if len(s.prevPoint) > 0 {

		distance = calculateDistance(s.prevPoint[0], s.prevPoint[1], data.Lat, data.Lon)
	}
	s.prevPoint = []float64{data.Lat, data.Lon}
	return distance, nil
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {

	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
