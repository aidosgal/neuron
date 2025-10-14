package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aidosgal/neuron/config"
	"github.com/aidosgal/neuron/pkg/jwt"
	"github.com/aidosgal/neuron/services/auth/entity"
	"github.com/aidosgal/neuron/services/auth/storage"
	"golang.org/x/crypto/bcrypt"
)

type usecase struct {
	log     *slog.Logger
	storage storage.Storage
	cfg     *config.Config
}

type Usecase interface {
	Login(ctx context.Context, req *entity.LoginRequest) (*entity.LoginResponse, error)
	Get(ctx context.Context, id int) (*entity.Admin, error)
}

func New(log *slog.Logger, storage storage.Storage, cfg *config.Config) Usecase {
	return &usecase{
		log:     log,
		cfg:     cfg,
		storage: storage,
	}
}

func (u *usecase) Login(ctx context.Context, req *entity.LoginRequest) (*entity.LoginResponse, error) {
	admin, err := u.storage.GetByLogin(ctx, req.Login)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password)); err != nil {
		return nil, fmt.Errorf("invlid password or login")
	}

	token, err := jwt.Generate(ctx, admin.ID, u.cfg.JwtSecret)
	if err != nil {
		return nil, err
	}

	return &entity.LoginResponse{
		Token: token,
		Admin: *admin,
	}, nil
}

func (u *usecase) Get(ctx context.Context, id int) (*entity.Admin, error) {
	admin, err := u.storage.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return admin, nil
}
