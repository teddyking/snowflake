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

func (s *SuiteService) Create(ctx context.Context, req *api.SuiteCreateRequest) (*api.SuiteCreateResponse, error) {
	if err := s.store.Create(req.Summary); err != nil {
		return &api.SuiteCreateResponse{}, err
	}

	return &api.SuiteCreateResponse{}, nil
}

func (s *SuiteService) List(ctx context.Context, req *api.SuiteListRequest) (*api.SuiteListResponse, error) {
	summaries, err := s.store.List()
	if err != nil {
		return &api.SuiteListResponse{}, err
	}

	return &api.SuiteListResponse{SuiteSummaries: summaries}, nil
}

func (s *SuiteService) Get(ctx context.Context, req *api.SuiteGetRequest) (*api.SuiteGetResponse, error) {
	test, err := s.store.Get(req.Codebase, req.Commit, req.Location)
	if err != nil {
		return &api.SuiteGetResponse{}, err
	}

	return &api.SuiteGetResponse{Test: test}, nil
}
