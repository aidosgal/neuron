package storage

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aidosgal/neuron/ent"
	"github.com/aidosgal/neuron/services/device/entity"
	"github.com/google/uuid"
)

type storage struct {
	log    *slog.Logger
	client *ent.Client
}

type Storage interface {
	CreateDevice(ctx context.Context, name string) (*entity.Device, error)
	GetDevice(ctx context.Context, id uuid.UUID) (*entity.Device, error)
}

func New(log *slog.Logger, client *ent.Client) Storage {
	return &storage{
		log:    log,
		client: client,
	}
}

func (s *storage) CreateDevice(ctx context.Context, name string) (*entity.Device, error) {
	device, err := s.client.Device.Create().
		SetDeviceName(name).
		Save(ctx)
	if err != nil {
		s.log.Error("failed to create Device", slog.String("Error", err.Error()))
		return nil, fmt.Errorf("failed to create Device: %w", err)
	}
	
	return entity.MakeStorageDeviceToEntity(device), nil
}

func (s *storage) GetDevice(ctx context.Context, id uuid.UUID) (*entity.Device, error) {
	device, err := s.client.Device.Get(ctx, id)
	if err != nil {
		s.log.Error("failed to get Device", slog.String("Error", err.Error()))
		return nil, fmt.Errorf("failed to get Device: %w", err)
	}
	
	return entity.MakeStorageDeviceToEntity(device), nil
}

