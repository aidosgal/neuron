package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aidosgal/neuron/config"
	"github.com/aidosgal/neuron/pkg/jwt"
	"github.com/aidosgal/neuron/services/clips/engine"
	"github.com/aidosgal/neuron/services/clips/entity"
	deviceUsecase "github.com/aidosgal/neuron/services/device/usecase"
)

type usecase struct {
	cfg    *config.Config
	engine *engine.Engine
	log    *slog.Logger
	device deviceUsecase.Usecase
}

type Usecase interface {
	Execute(ctx context.Context, token string, req *entity.Input) (*entity.Output, error)
}

func New(cfg *config.Config, engine *engine.Engine, log *slog.Logger, device deviceUsecase.Usecase) Usecase {
	return &usecase{
		cfg:    cfg,
		engine: engine,
		log:    log,
		device: device,
	}
}

func (u *usecase) Execute(ctx context.Context, token string, req *entity.Input) (*entity.Output, error) {
	defer u.engine.Destroy()
	deviceID, err := jwt.ParseDeviceID(ctx, token, u.cfg.JwtSecret)
	if err != nil {
		return nil, fmt.Errorf("access denied")
	}

	_, err = u.device.GetDevice(ctx, deviceID)
	if err != nil {
		return nil, fmt.Errorf("access denied")
	}


	u.engine.LoadRules([]string{
		"services/clips/rules/facts.clp",
		"services/clips/rules/congestion.clp",
		"services/clips/rules/incidents.clp",
		"services/clips/rules/weather.clp",
		"services/clips/rules/policy.clp",
		"services/clips/rules/routing.clp",
		"services/clips/rules/coordination.clp",
	})

	output := u.engine.Infer(*req)
	return &output, nil
}
