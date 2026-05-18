package portfolios

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/LuisCabantac/portfolyo-go-api/internal/json"
	"github.com/LuisCabantac/portfolyo-go-api/internal/screenshot"
	"github.com/clerk/clerk-sdk-go/v2"
	convex "github.com/inselfcontroll/convex-go"
	"github.com/inselfcontroll/convex-go/src/codebase"
)

type svc struct {
	client *codebase.Client
}

func NewService(client *codebase.Client) *svc {
	return &svc{
		client: client,
	}
}

func (s *svc) CreatePortfolioScreenshot(ctx context.Context, portfolioID string, authUser *clerk.User) (StorageResponse, error) {
	portfolio, err := convex.Query[Portfolio](ctx, s.client, "queries/portfolios:getPortfolioById", map[string]any{
		"portfolioId": portfolioID,
	})
	if err != nil {
		return StorageResponse{}, fmt.Errorf("%w: %v", ErrMissingIDParam, err)
	}

	usr, err := convex.Query[User](ctx, s.client, "queries/users:getUserByClerkId", map[string]any{
		"clerkId": authUser.ID,
	})
	if err != nil {
		return StorageResponse{}, fmt.Errorf("%w: %v", ErrUserNotFound, err)
	}

	if usr.Portfolio != portfolio.ID {
		return StorageResponse{}, fmt.Errorf("%w: %v", ErrUnauthorizedPortfolioAccess, err)
	}

	buf, err := screenshot.Screenshot(portfolio.URL)
	if err != nil {
		return StorageResponse{}, fmt.Errorf("%w: %v", ErrScreenshotCapture, err)
	}

	_, err = convex.Mutation[Portfolio](ctx, s.client, "mutations/portfolios:deletePortfolioScreenshot", map[string]any{
		"portfolioId": portfolioID,
	})
	if err != nil {
		return StorageResponse{}, fmt.Errorf("%w: %v", ErrRemoveExistingScreenshot, err)
	}

	uploadUrl, err := convex.Mutation[any](ctx, s.client, "mutations/portfolios:generateUploadUrl", map[string]any{})
	if err != nil {
		return StorageResponse{}, fmt.Errorf("%w: %v", ErrGenerateUploadURL, err)
	}

	blobReader := bytes.NewReader(buf)

	url, ok := uploadUrl.(string)
	if !ok {
		return StorageResponse{}, fmt.Errorf("%w: %v", ErrInvalidUploadURL, err)
	}

	req, err := http.NewRequest("POST", url, blobReader)
	if err != nil {
		return StorageResponse{}, fmt.Errorf("%w: %v", ErrUploadRequest, err)
	}
	req.Header.Set("Content-Type", "image/png")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return StorageResponse{}, fmt.Errorf("%w: %v", ErrScreenshotUpload, err)
	}
	defer resp.Body.Close()

	var storageId UploadUrlResponse

	err = json.ReadJSON(resp, &storageId)

	_, err = convex.Mutation[any](ctx, s.client, "mutations/portfolios:savePortfolioScreenshot", map[string]any{
		"portfolioId": portfolioID,
		"storageId":   storageId.StorageId,
	})
	if err != nil {
		return StorageResponse{}, fmt.Errorf("%w: %v", ErrSaveScreenshotMetadata, err)
	}

	return StorageResponse{
		Message:    "Portfolio screenshot taken and stored successfully!",
		StatusCode: http.StatusCreated,
		StorageId:  &storageId.StorageId,
	}, nil
}
