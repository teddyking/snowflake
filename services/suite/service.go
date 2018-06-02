package suite

import (
	"context"

	"github.com/teddyking/snowflake/api"
)

//go:generate counterfeiter . Store
type Store interface {
	Create(suiteSummary *api.SuiteSummary) error
	List() ([]*api.SuiteSummary, error)
	Get(codebase, commit, location string) (*api.Test, error)
}

type SuiteService struct {
	store Store
}

func New(store Store) *SuiteService {
	return &SuiteService{
		store: store,
	}
}

func (s *SuiteService) Create(ctx context.Context, req *api.CreateRequest) (*api.CreateResponse, error) {
	if err := s.store.Create(req.Summary); err != nil {
		return &api.CreateResponse{}, err
	}

	return &api.CreateResponse{}, nil
}

func (s *SuiteService) List(ctx context.Context, req *api.ListRequest) (*api.ListResponse, error) {
	summaries, err := s.store.List()
	if err != nil {
		return &api.ListResponse{}, err
	}

	return &api.ListResponse{SuiteSummaries: summaries}, nil
}

func (s *SuiteService) Get(ctx context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	test, err := s.store.Get(req.Codebase, req.Commit, req.Location)
	if err != nil {
		return &api.GetResponse{}, err
	}

	return &api.GetResponse{Test: test}, nil
}
