package server

import (
	"context"

	"github.com/teddyking/snowflake/api"
)

//go:generate counterfeiter . Store
type Store interface {
	Create(suiteSummary *api.SuiteSummary) error
	List() ([]*api.SuiteSummary, error)
}

type Server struct {
	store Store
}

func New(store Store) *Server {
	return &Server{
		store: store,
	}
}

func (s *Server) Create(ctx context.Context, req *api.CreateRequest) (*api.CreateResponse, error) {
	if err := s.store.Create(req.Summary); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *Server) List(ctx context.Context, req *api.ListRequest) (*api.ListResponse, error) {
	summaries, err := s.store.List()
	if err != nil {
		return nil, err
	}

	return &api.ListResponse{SuiteSummaries: summaries}, nil
}
