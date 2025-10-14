package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aidosgal/neuron/config"
	"github.com/aidosgal/neuron/pkg/jwt"
	authUsecase "github.com/aidosgal/neuron/services/auth/usecase"
	"github.com/aidosgal/neuron/services/device/entity"
	"github.com/aidosgal/neuron/services/device/storage"
)

type usecase struct {
	log     *slog.Logger
	storage storage.Storage
	auth    authUsecase.Usecase
	cfg     *config.Config
}

type Usecase interface {
	CreateDevice(ctx context.Context, req *entity.CreateRequest) (*entity.CreateResponse, error)
	GetDevice(ctx context.Context, deviceID string) (*entity.Device, error)
}

func New(log *slog.Logger, storage storage.Storage, auth authUsecase.Usecase, cfg *config.Config) Usecase {
	return &usecase{
		log:     log,
		storage: storage,
		auth:    auth,
		cfg:     cfg,
	}
}

func (u *usecase) CreateDevice(ctx context.Context, req *entity.CreateRequest) (*entity.CreateResponse, error) {
	adminID, err := jwt.ParseUserID(ctx, req.AdminToken, u.cfg.JwtSecret)
	if err != nil {
		return nil, err
	}

	_, err = u.auth.Get(ctx, adminID)
	if err != nil {
		return nil, fmt.Errorf("access denied")
	}

	device, err := u.storage.CreateDevice(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	token, err := jwt.GenerateDeviceToken(ctx, device.ID, u.cfg.JwtSecret)
	if err != nil {
		return nil, err
	}

	device.DeviceToken = token
	
	return &entity.CreateResponse{
		Device: *device,
	}, nil
}

func (u *usecase) GetDevice(ctx context.Context, deviceID string) (*entity.Device, error) {
	device, err := u.storage.CreateDevice(ctx, deviceID)
	if err != nil {
		return nil, err
	}
	
	return device, nil
}
