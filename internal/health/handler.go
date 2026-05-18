package health

import (
	"net/http"

	"github.com/LuisCabantac/portfolyo-go-api/internal/json"
)

const apiName = "Portfolyo API"

const version = "1.0.0"

type HealthCheckResponse struct {
	Message     string `json:"message"`
	Version     string `json:"version"`
	Status      string `json:"status"`
	Environment string `json:"environment"`
}

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	env, logger := h.service.HealthCheck()

	if err := json.Write(w, http.StatusOK, HealthCheckResponse{
		Message:     apiName,
		Version:     version,
		Status:      "running",
		Environment: env,
	}); err != nil {
		logger.Error("the server encountered a problem and could not process your request", "error", err)
	}
}
