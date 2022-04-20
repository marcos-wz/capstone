package server

import (
	"context"
	"github.com/marcos-wz/capstone/proto/loaderpb"
)

func (*server) Loader(ctx context.Context, in *loaderpb.LoaderRequest) (*loaderpb.LoaderResponse, error) {
	return nil, nil
}
