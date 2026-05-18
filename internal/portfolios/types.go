package portfolios

import (
	"context"
	"net/http"

	"github.com/LuisCabantac/portfolyo-go-api/internal/apperrors"
	"github.com/clerk/clerk-sdk-go/v2"
)

type Service interface {
	CreatePortfolioScreenshot(ctx context.Context, portfolioID string, authUser *clerk.User) (StorageResponse, error)
}

type StorageResponse struct {
	Message    string  `json:"message"`
	StatusCode int     `json:"statusCode"`
	Error      *string `json:"error,omitempty"`
	StorageId  *string `json:"storageId,omitempty"`
}

type Portfolio struct {
	ID        string  `json:"_id"`
	StorageId *string `json:"storageId,omitempty"`
	OwnerId   string  `json:"ownerId"`
	URL       string  `json:"url,omitempty"`
}

type User struct {
	ID        string `json:"_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Portfolio string `json:"portfolio,omitempty"`
}

type UploadUrlResponse struct {
	StorageId string `json:"storageId"`
}

var (
	ErrInvalidToken = &apperrors.AppError{
		Message:    "Invalid or expired token.",
		Code:       "invalid_token",
		StatusCode: http.StatusUnauthorized}
	ErrMissingIDParam = &apperrors.AppError{
		Message:    "ID parameter is required.",
		Code:       "missing_id_param",
		StatusCode: http.StatusBadRequest}
	ErrPortfolioNotFound = &apperrors.AppError{
		Message:    "Portfolio not found.",
		Code:       "portfolio_not_found",
		StatusCode: http.StatusNotFound}
	ErrUserNotFound = &apperrors.AppError{
		Message:    "User not found.",
		Code:       "user_not_found",
		StatusCode: http.StatusNotFound}
	ErrUnauthorizedPortfolioAccess = &apperrors.AppError{
		Message:    "You are not authorized to access this portfolio.",
		Code:       "portfolio_access_denied",
		StatusCode: http.StatusForbidden}
	ErrScreenshotCapture = &apperrors.AppError{
		Message:    "Failed to capture portfolio screenshot.",
		Code:       "screenshot_capture_failed",
		StatusCode: http.StatusInternalServerError}
	ErrRemoveExistingScreenshot = &apperrors.AppError{
		Message:    "Failed to remove existing screenshot.",
		Code:       "screenshot_delete_failed",
		StatusCode: http.StatusInternalServerError}
	ErrGenerateUploadURL = &apperrors.AppError{
		Message:    "Failed to generate upload URL.",
		Code:       "upload_url_generation_failed",
		StatusCode: http.StatusBadGateway}
	ErrInvalidUploadURL = &apperrors.AppError{
		Message:    "Invalid upload URL received from storage service.",
		Code:       "invalid_upload_url",
		StatusCode: http.StatusBadGateway}
	ErrUploadRequest = &apperrors.AppError{
		Message:    "Failed to prepare upload request.",
		Code:       "upload_request_creation_failed",
		StatusCode: http.StatusInternalServerError}
	ErrScreenshotUpload = &apperrors.AppError{
		Message:    "Failed to upload screenshot to storage service.",
		Code:       "screenshot_upload_failed",
		StatusCode: http.StatusBadGateway}
	ErrStorageResponse = &apperrors.AppError{
		Message:    "Invalid response from storage service.",
		Code:       "invalid_storage_response",
		StatusCode: http.StatusBadGateway}
	ErrSaveScreenshotMetadata = &apperrors.AppError{
		Message:    "Failed to save screenshot metadata.",
		Code:       "screenshot_save_failed",
		StatusCode: http.StatusInternalServerError}
	ErrParsingUploadResponse = &apperrors.AppError{
		Message:    "Failed to parse upload service response.",
		Code:       "invalid_upload_response",
		StatusCode: http.StatusInternalServerError,
	}
)
