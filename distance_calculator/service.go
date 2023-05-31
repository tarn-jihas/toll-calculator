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
	points [][]float64
}

func NewCalcService() CalculatorServicer {
	return &CalculatorService{
		points: make([][]float64, 0),
	}
}

func (s *CalculatorService) CalculateDistance(data types.OBUData) (float64, error) {
	var distance = 0.0

	if len(s.points) > 0 {
		prevPoint := s.points[len(s.points)-1]
		distance = calculateDistance(prevPoint[0], prevPoint[1], data.Lat, data.Lon)
	}
	//distance := calculateDistance(data.Lat, data.Lon)
	s.points = append(s.points, []float64{data.Lat, data.Lon})
	return distance, nil
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {

	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
