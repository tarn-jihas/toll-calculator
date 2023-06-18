package main

import (
	"context"

	"github.com/keselj-strahinja/toll-calculator/types"
)

type GRPCAggregatorServer struct {
	types.UnimplementedAggregatorServer
	svc Aggregator
}

func NewGRPCAggregatorServer(svc Aggregator) *GRPCAggregatorServer {
	return &GRPCAggregatorServer{
		svc: svc,
	}
}

// transport layer
// JSON -> types.Distance -> all done (same type)
// GRPC -> types.AggregateRequest -> types.Distance
// Webpack transprot -> types.Webpack -> types.Distances

// business layer -> business layer type (main type everyone needs to convert to)
func (s *GRPCAggregatorServer) Aggregate(ctx context.Context, req *types.AggregateRequest) (*types.None, error) {
	distance := types.Distance{
		OBUID: int(req.ObuID),
		Value: req.Value,
		Unix:  req.Unix,
	}

	err := s.svc.AggregateDistance(distance)

	if err != nil {
		return nil, err
	}

	return &types.None{}, nil
}
