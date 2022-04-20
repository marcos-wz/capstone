package client

import (
	"github.com/marcos-wz/capstone/proto/filterpb"
	"github.com/marcos-wz/capstone/proto/fruitpb"
	"strings"

	"google.golang.org/grpc"
)

// CONTRACT
type iClient interface {
	Filter(filter, value string)
	Loader()
	FilterCC(filter, value string)
}

type client struct {
	service fruitpb.FruitServiceClient
}

func NewClient(conn *grpc.ClientConn) iClient {
	return &client{
		service: fruitpb.NewFruitServiceClient(conn),
	}
}

// IMPLEMENTATION
func (*client) getAllowedFilter(value string) filterpb.FiltersAllowed {
	switch strings.ToUpper(value) {
	case "ID":
		return filterpb.FiltersAllowed_FILTER_ID
	case "NAME":
		return filterpb.FiltersAllowed_FILTER_NAME
	case "COLOR":
		return filterpb.FiltersAllowed_FILTER_COLOR
	case "COUNTRY":
		return filterpb.FiltersAllowed_FILTER_COUNTRY
	default:
		return filterpb.FiltersAllowed_FILTER_UNDEFINED
	}
}
