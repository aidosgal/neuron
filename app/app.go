package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/aidosgal/neuron/config"
	"github.com/aidosgal/neuron/ent"
	adminServer "github.com/aidosgal/neuron/services/auth/server"
	adminStorage "github.com/aidosgal/neuron/services/auth/storage"
	adminUsecase "github.com/aidosgal/neuron/services/auth/usecase"
	"github.com/aidosgal/neuron/services/clips/engine"
	"github.com/aidosgal/neuron/services/clips/server"
	"github.com/aidosgal/neuron/services/clips/usecase"
	deviceServer "github.com/aidosgal/neuron/services/device/server"
	deviceStorage "github.com/aidosgal/neuron/services/device/storage"
	deviceUsecase "github.com/aidosgal/neuron/services/device/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
)

type app struct {
	log *slog.Logger
	cfg *config.Config
}

type App interface {
	Run() error
}

func New(cfg *config.Config, log *slog.Logger) App {
	return &app{
		cfg: cfg,
		log: log,
	}
}


func (a *app) Run() error {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.URLFormat)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	entPsqlConnect := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		a.cfg.Database.Host,
		a.cfg.Database.Port,
		a.cfg.Database.User,
		a.cfg.Database.Name,
		a.cfg.Database.Password,
		a.cfg.Database.SSLMode,
	)

	client, err := ent.Open("postgres", entPsqlConnect)
	if err != nil {
		a.log.Error("failed opening connection to postgres", slog.String("error", err.Error()))
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		a.log.Error("failed creating schema resources", slog.String("error", err.Error()))
	}

	count, err := client.Admin.Query().Count(context.Background())
	if err != nil {
		return err
	}
	if count == 0 {
		password, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		_, err = client.Admin.Create().
			SetLogin("admin").
			SetPassword(string(password)).
			Save(context.Background())
		if err != nil {
			return err
		}
	}

	adminStorage := adminStorage.New(a.log, client)
	adminUsecase := adminUsecase.New(a.log, adminStorage, a.cfg)
	adminServer := adminServer.New(a.log, adminUsecase)

	deviceStorage := deviceStorage.New(a.log, client)
	deviceUsecase := deviceUsecase.New(a.log, deviceStorage, adminUsecase, a.cfg)
	deviceServer := deviceServer.New(deviceUsecase)

	clipsEngine := engine.NewEngine()
	clipsUsecase := usecase.New(a.cfg, clipsEngine, a.log, deviceUsecase)
	clipsServer := server.New(clipsUsecase)

	router.Route("/api/v1", func(apiRouter chi.Router) {
		apiRouter.Route("/admin", func(authRouter chi.Router) {
			authRouter.Post("/login", adminServer.Login)
		})
		apiRouter.Route("/device", func(authRouter chi.Router) {
			authRouter.Post("/", deviceServer.CreateDevice)
		})
		apiRouter.Route("/execute", func(authRouter chi.Router) {
			authRouter.Post("/", clipsServer.Execute)
		})
	})
	
	a.log.Debug("server running", slog.String("address", fmt.Sprintf("localhost:%d", a.cfg.Port)))
	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", a.cfg.Port), router)
}
