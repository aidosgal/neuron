package storage

import (
	"context"
	"log/slog"

	"github.com/aidosgal/neuron/ent"
	"github.com/aidosgal/neuron/ent/admin"
	"github.com/aidosgal/neuron/services/auth/entity"
)

type storage struct {
	log    *slog.Logger
	client *ent.Client
}

type Storage interface {
	Get(ctx context.Context, id int) (*entity.Admin, error)
	GetByLogin(ctx context.Context, login string) (*entity.Admin, error)
}

func New(log *slog.Logger, client *ent.Client) Storage {
	return &storage{
		log:    log,
		client: client,
	}
}

func (s *storage) Get(ctx context.Context, id int) (*entity.Admin, error) {
	admin, err := s.client.Admin.Get(ctx, id)
	if err != nil {
		s.log.Error("failed to get admin", slog.String("error", err.Error()))
		return nil, err
	}

	return entity.MakeStorageAdminToEntity(admin), nil
}

func (s *storage) GetByLogin(ctx context.Context, login string) (*entity.Admin, error) {
	admin, err := s.client.Admin.Query().
		Where(
			admin.LoginEQ(login),
		).First(ctx)
	if err != nil {
		s.log.Error("failed to get admin", slog.String("error", err.Error()))
		return nil, err
	}

	return entity.MakeStorageAdminToEntity(admin), nil
}
