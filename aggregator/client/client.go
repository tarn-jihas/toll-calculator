package client

import (
	"context"

	"github.com/keselj-strahinja/toll-calculator/types"
)

type Client interface {
	Aggregate(context.Context, *types.AggregateRequest) error
}
