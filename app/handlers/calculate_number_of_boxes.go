package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/pillsomi/gymshark/app/controllers"
)

// CalculateNumberOfBoxes handler contains the logic for parsing the request and calculate the number of boxes.
func CalculateNumberOfBoxes(
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
		var payload CalculateNumberOfBoxesRequest
		if err := json.Unmarshal(body, &payload); err != nil {
			reportError(w, http.StatusBadRequest, "invalid request body")
			return
		}

		// Calculate number of boxes controller.
		boxes, err := controller.CalculateNumberOfBoxes(payload.NumberOfItems)
		if err != nil {
			code := http.StatusInternalServerError
			message := "internal server error"
			if errors.Is(controllers.ErrorInvalidNumberOfItemsInput, err) {
				code = http.StatusBadRequest
				message = err.Error()
			}
			reportError(w, code, message)
			return
		}

		// Transform results.
		result := &CalculateNumberOfBoxesResponse{}
		boxDetails := make([]BoxDetail, 0)
		for _, b := range boxes {
			boxDetails = append(boxDetails, BoxDetail{
				Size:   b.Size,
				Number: b.Number,
			})
		}
		result.BoxDetails = boxDetails

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(result)
	}
	return fn
}

// CalculateNumberOfBoxesRequest request body.
type CalculateNumberOfBoxesRequest struct {
	NumberOfItems int `json:"number_of_items"`
}

// CalculateNumberOfBoxesResponse response body.
type CalculateNumberOfBoxesResponse struct {
	BoxDetails []BoxDetail `json:"number_of_boxes"`
}

// BoxDetail details for the boxes, size and number of items.
type BoxDetail struct {
	Size   int `json:"size"`
	Number int `json:"number"`
}
