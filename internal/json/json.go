package json

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/LuisCabantac/portfolyo-go-api/internal/apperrors"
)

func Write(w http.ResponseWriter, status int, data any) error {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(buf.Bytes())

	return nil
}

func WriteErrorResponse(w http.ResponseWriter, err error) error {
	appErr, ok := errors.AsType[*apperrors.AppError](err)
	if !ok {
		return Write(w, apperrors.ErrInternalServerError.StatusCode, apperrors.ErrInternalServerError)
	}

	return Write(w, appErr.StatusCode, appErr)
}

func WriteJSON(w io.Writer, data any) error {
	return json.NewEncoder(w).Encode(data)
}

func Read(r *http.Request, data any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func ReadJSON(w *http.Response, data any) error {
	return json.NewDecoder(w.Body).Decode(data)
}
