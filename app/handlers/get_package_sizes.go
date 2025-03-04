package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pillsomi/gymshark/app/controllers"
)

// GetPackageSizes handler contains the logic for retrieving the different sizes of packages.
func GetPackageSizes(
	controller controllers.Controller,
) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Add default headers
		w.Header().Set("Content-Type", "application/json")

		// Requests MUST be POST
		if r.Method != http.MethodGet {
			reportError(w, http.StatusBadRequest, "invalid http method")
			return
		}

		// Get package size controller.
		packageSizes, err := controller.GetPackageSizes()
		if err != nil {
			reportError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		// Parse result.
		result := &GetPackageSizesResponse{PackageSizes: packageSizes}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(result)
	}
	return fn
}

// GetPackageSizesResponse response body.
type GetPackageSizesResponse struct {
	PackageSizes []int `json:"package_sizes"`
}
