package screenshot

import (
	"net/http"

	"github.com/LuisCabantac/go-portfolyo-api/internal/apperrors"
)

var (
	ErrScreenshotCapture = &apperrors.AppError{
		Message:    "Failed to capture portfolio screenshot.",
		Code:       "screenshot_capture_failed",
		StatusCode: http.StatusInternalServerError}
	ErrMissingPortfolioURL = &apperrors.AppError{
		Message:    "Portfolio URL is missing or invalid.",
		Code:       "missing_portfolio_url",
		StatusCode: http.StatusBadRequest}
)
