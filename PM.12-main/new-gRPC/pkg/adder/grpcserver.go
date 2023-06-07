package adder

import (
	"context"
	"log"
	"new-gRPC/pkg/api"
)

// GRPCServer ...
type GRPCServer struct {
	api.UnimplementedAdderServer
}

// Add ...
func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	response := &api.AddResponse{Result: req.GetX() + req.GetY()}
	log.Printf("%v + %v = %v", req.X, req.Y, response.Result)
	return response, nil
}

// NewConnect
func (s *GRPCServer) NewConnect(ctx context.Context, req *api.NewConnRequest) (*api.NewConnResponse, error) {
	response := &api.NewConnResponse{Result: req.Name}
	log.Printf("User '%v' connected", req.Name)
	return response, nil
}
