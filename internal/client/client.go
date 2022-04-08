package client

import (
	"strings"

	pb "github.com/marcos-wz/capstone/internal/fruitpb"
	"google.golang.org/grpc"
)

// CONTRACT
type iClient interface {
	Filter(filter, value string)
	Loader()
	FilterCC(filter, value string)
}

type client struct {
	service pb.FruitServiceClient
}

func NewClient(conn *grpc.ClientConn) iClient {
	return &client{
		service: pb.NewFruitServiceClient(conn),
	}
}

// IMPLEMENTATION
func (*client) getAllowedFilter(value string) pb.FiltersAllowed {
	switch strings.ToUpper(value) {
	case "ID":
		return pb.FiltersAllowed_ID
	case "NAME":
		return pb.FiltersAllowed_NAME
	case "COLOR":
		return pb.FiltersAllowed_COLOR
	case "COUNTRY":
		return pb.FiltersAllowed_COUNTRY
	default:
		return pb.FiltersAllowed_UNDEFINED
	}
}
