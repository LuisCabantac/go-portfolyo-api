package portfolios

import (
	"bytes"
	"context"
	"net/http"

	"github.com/LuisCabantac/portfolyo-go-api/internal/json"
	"github.com/LuisCabantac/portfolyo-go-api/internal/screenshot"
	"github.com/clerk/clerk-sdk-go/v2"
	convex "github.com/inselfcontroll/convex-go"
	"github.com/inselfcontroll/convex-go/src/codebase"
)

type svc struct {
	client     *codebase.Client
	httpClient *http.Client
}

func NewService(client *codebase.Client, httpClient *http.Client) *svc {
	return &svc{
		client:     client,
		httpClient: httpClient,
	}
}

func (s *svc) CreatePortfolioScreenshot(ctx context.Context, portfolioID string, authUser *clerk.User, theme string) (StorageResponse, error) {
	portfolio, err := convex.Query[Portfolio](ctx, s.client, "queries/portfolios:getPortfolioById", map[string]any{
		"portfolioId": portfolioID,
	})
	if err != nil {
		return StorageResponse{}, ErrMissingIDParam
	}

	usr, err := convex.Query[User](ctx, s.client, "queries/users:getUserByClerkId", map[string]any{
		"clerkId": authUser.ID,
	})
	if err != nil {
		return StorageResponse{}, ErrUserNotFound
	}

	if usr.Portfolio != portfolio.ID {
		return StorageResponse{}, ErrUnauthorizedPortfolioAccess
	}

	buf, err := screenshot.Screenshot(portfolio.URL, theme)
	if err != nil {
		return StorageResponse{}, ErrScreenshotCapture
	}

	_, err = convex.Mutation[Portfolio](ctx, s.client, "mutations/portfolios:deletePortfolioScreenshot", map[string]any{
		"portfolioId": portfolioID,
	})
	if err != nil {
		return StorageResponse{}, ErrRemoveExistingScreenshot
	}

	uploadUrl, err := convex.Mutation[any](ctx, s.client, "mutations/portfolios:generateUploadUrl", map[string]any{})
	if err != nil {
		return StorageResponse{}, ErrGenerateUploadURL
	}

	blobReader := bytes.NewReader(buf)

	url, ok := uploadUrl.(string)
	if !ok {
		return StorageResponse{}, ErrInvalidUploadURL
	}

	req, err := http.NewRequest("POST", url, blobReader)
	if err != nil {
		return StorageResponse{}, ErrUploadRequest
	}
	req.Header.Set("Content-Type", "image/png")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return StorageResponse{}, ErrScreenshotUpload
	}
	defer resp.Body.Close()

	var storageId UploadUrlResponse

	err = json.ReadJSON(resp, &storageId)
	if err != nil {
		return StorageResponse{}, ErrParsingUploadResponse
	}

	_, err = convex.Mutation[any](ctx, s.client, "mutations/portfolios:savePortfolioScreenshot", map[string]any{
		"portfolioId": portfolioID,
		"storageId":   storageId.StorageId,
	})
	if err != nil {
		return StorageResponse{}, ErrSaveScreenshotMetadata
	}

	return StorageResponse{
		Message:    "Portfolio screenshot taken and stored successfully!",
		StatusCode: http.StatusCreated,
		StorageId:  &storageId.StorageId,
	}, nil
}
