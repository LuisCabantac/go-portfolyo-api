package json

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/LuisCabantac/go-portfolyo-api/internal/apperrors"
)

func Write[T comparable](w http.ResponseWriter, status int, data T) error {
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

func WriteJSON[T comparable](w io.Writer, data T) error {
	return json.NewEncoder(w).Encode(data)
}

func Read[T comparable](r *http.Request, data T) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func ReadJSON[T comparable](resp *http.Response, data T) error {
	return json.NewDecoder(resp.Body).Decode(data)
}
