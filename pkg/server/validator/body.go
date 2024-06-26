package validator

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/mauriciofsnts/gofast/pkg/server/api"
)

/*
MustGetBody is a helper function that decodes the body of an HTTP request into a struct and validates it. Returns the struct and a bool that is true if it's valid and false otherwise.
The function writes an error response to the response writer if the body is missing or if the body is invalid.
*/
func MustGetBody[T any](w http.ResponseWriter, r *http.Request) (T, bool) {
	var body T

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		slog.Error("Failed to decode body", "err", err)
		if errors.Is(err, io.EOF) {
			api.Err(w, api.BadRequestErr, "Missing body")
		} else {
			api.Err(w, api.BadRequestErr, err.Error())
		}
		return body, false
	}

	validationErrors := Validate(body)

	if validationErrors == nil {
		return body, true
	}

	api.DetailedError(w, api.ValidationErr, validationErrors)

	return body, false
}
