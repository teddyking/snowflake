package reporter

import (
	"context"

	"github.com/teddyking/snowflake/api"
)

//go:generate counterfeiter . Store
type Store interface {
	CreateReport(report *api.Report) error
}

type Service struct {
	store Store
}

func New(store Store) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) Create(ctx context.Context, req *api.ReporterCreateReq) (*api.ReporterCreateRes, error) {
	if err := s.store.CreateReport(req.Report); err != nil {
		return &api.ReporterCreateRes{}, err
	}

	return &api.ReporterCreateRes{}, nil
}
