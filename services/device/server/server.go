package server

import (
	"fmt"
	"net/http"

	"github.com/aidosgal/neuron/pkg/json"
	"github.com/aidosgal/neuron/pkg/jwt"
	"github.com/aidosgal/neuron/services/device/entity"
	"github.com/aidosgal/neuron/services/device/usecase"
)

type server struct {
	usecase usecase.Usecase
}

type Server interface {
	CreateDevice(w http.ResponseWriter, r *http.Request)
}

func New(usecase usecase.Usecase) Server {
	return &server{
		usecase: usecase,
	}
}

func (s *server) CreateDevice(w http.ResponseWriter, r *http.Request) {
	adminToken, err := jwt.ParseTokenFromHeader(r)
	if err != nil {
		json.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid credentials"))
		return
	}

	req := &entity.CreateRequest{}
	err = json.ParseJSON(r, req)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req.AdminToken = adminToken

	resp, err := s.usecase.CreateDevice(r.Context(), req)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusCreated, resp)
}
