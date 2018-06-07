package flaker

import (
	"context"

	"github.com/teddyking/snowflake/api"
)

//go:generate counterfeiter . Store
type Store interface {
	ListReports() ([]*api.Report, error)
}

type FlakeAnalyser func(reports []*api.Report) ([]*api.Flake, error)

type Service struct {
	store         Store
	flakeAnalyser FlakeAnalyser
}

func New(store Store, flakeAnalyser FlakeAnalyser) *Service {
	return &Service{
		store:         store,
		flakeAnalyser: flakeAnalyser,
	}
}

func (s *Service) List(context.Context, *api.FlakerListReq) (*api.FlakerListRes, error) {
	reports, err := s.store.ListReports()
	if err != nil {
		return &api.FlakerListRes{}, err
	}

	flakes, err := s.flakeAnalyser(reports)
	if err != nil {
		return &api.FlakerListRes{}, err
	}

	return &api.FlakerListRes{Flakes: flakes}, nil
}
