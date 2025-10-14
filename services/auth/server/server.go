package server

import (
	"log/slog"
	"net/http"

	"github.com/aidosgal/neuron/pkg/json"
	"github.com/aidosgal/neuron/services/auth/entity"
	"github.com/aidosgal/neuron/services/auth/usecase"
)

type server struct {
	log     *slog.Logger
	usecase usecase.Usecase
}

type Server interface {
	Login(w http.ResponseWriter, r *http.Request)
}

func New(log *slog.Logger, usecase usecase.Usecase) Server {
	return &server{
		log:     log,
		usecase: usecase,
	}
}

func (s *server) Login(w http.ResponseWriter, r *http.Request) {
	req := &entity.LoginRequest{}
	err := json.ParseJSON(r, req)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := s.usecase.Login(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, resp)
}
