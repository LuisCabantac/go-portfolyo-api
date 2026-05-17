package json

import (
	"bytes"
	"encoding/json"
	"net/http"
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
