package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/pillsomi/gymshark/app/controllers"
)

func UpdatePackageSizes(
	controller controllers.Controller,
) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Add default headers
		w.Header().Set("Content-Type", "application/json")

		// Requests MUST be POST
		if r.Method != http.MethodPost {
			reportError(w, http.StatusBadRequest, "invalid http method")
			return
		}

		// Parse request
		const MaxBodyBytes = int64(65536)
		r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			reportError(w, http.StatusInternalServerError, "server cannot parse request body")
			return
		}

		// Parse request body in the request struct.
		var payload UpdatePackageSizesRequest
		if err := json.Unmarshal(body, &payload); err != nil {
			reportError(w, http.StatusBadRequest, "invalid request body")
			return
		}

		// Update package size controller.
		updatedPackageSizes, err := controller.UpdatePackageSizes(payload.Packages)
		if err != nil {
			code := http.StatusInternalServerError
			message := "internal server error"
			if errors.Is(controllers.ErrorEmptyListInput, err) ||
				errors.Is(controllers.ErrorNonPositiveIntegerListInput, err) {
				code = http.StatusBadRequest
				message = err.Error()
			}
			reportError(w, code, message)
			return
		}

		// Parse result.
		result := &UpdatePackageSizesResponse{Packages: updatedPackageSizes}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(result)
	}
	return fn
}

// UpdatePackageSizesRequest request body.
type UpdatePackageSizesRequest struct {
	Packages []int `json:"packages"`
}

// UpdatePackageSizesResponse response body.
type UpdatePackageSizesResponse struct {
	Packages []int `json:"packages"`
}
