package flaker

import (
	"context"

	"github.com/teddyking/snowflake/api"
)

type FlakeAnalyser func(summaries []*api.SuiteSummary) ([]*api.Flake, error)

type Service struct {
	flakeAnalyser FlakeAnalyser
}

func New(flakeAnalyser FlakeAnalyser) *Service {
	return &Service{
		flakeAnalyser: flakeAnalyser,
	}
}

func (s *Service) List(ctx context.Context, req *api.FlakeListRequest) (*api.FlakeListResponse, error) {
	flakes, err := s.flakeAnalyser(req.Suites)
	if err != nil {
		return &api.FlakeListResponse{}, err
	}

	return &api.FlakeListResponse{Flakes: flakes}, nil
}
