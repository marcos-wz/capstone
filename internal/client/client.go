package client

import (
	"github.com/marcos-wz/capstone/proto/fruitpb"
	"google.golang.org/grpc"
)

var DebugLevel uint32

type FruitClient interface {
	Filter(filter, value string)
	Loader()
	FilterCC(filter, value string)
}

type fruitClient struct {
	service fruitpb.FruitServiceClient
}

func NewFruitClient(conn *grpc.ClientConn) FruitClient {
	return &fruitClient{
		service: fruitpb.NewFruitServiceClient(conn),
	}
}
