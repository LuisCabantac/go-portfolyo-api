package portfolios

import (
	"io"
	"net/http"

	"github.com/LuisCabantac/portfolyo-go-api/internal/json"
	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/go-chi/chi/v5"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) CreatePortfolioScreenshot(w http.ResponseWriter, r *http.Request) {
	portfolioID := chi.URLParam(r, "portfolioID")
	if portfolioID == "" {
		_ = json.WriteErrorResponse(w, ErrMissingIDParam)
		return
	}

	claims, ok := clerk.SessionClaimsFromContext(r.Context())
	if !ok {
		_ = json.WriteErrorResponse(w, ErrInvalidToken)
		return
	}

	usr, err := user.Get(r.Context(), claims.Subject)
	if err != nil {
		_ = json.WriteErrorResponse(w, ErrInvalidToken)
		return
	}

	theme := "light"
	portfolioTheme := PortfolioRequest{
		Theme: &theme,
	}

	err = json.Read(r, &portfolioTheme)
	if err != nil && err != io.EOF {
		return
	}

	resp, err := h.service.CreatePortfolioScreenshot(r.Context(), portfolioID, usr, *portfolioTheme.Theme)
	if err != nil {
		_ = json.WriteErrorResponse(w, err)
		return
	}

	json.Write(w, http.StatusCreated, resp)
}
