package handlers

import (
	"api_metering_system/metering"
	"encoding/json"
	"net/http"
	"strconv"
)

// StorageHandler holds the storage metering logic.
type StorageHandler struct {
	storageMetrics *metering.StorageMetrics
}

// NewStorageHandler creates a new StorageHandler.
func NewStorageHandler(storageMetrics *metering.StorageMetrics) *StorageHandler {
	return &StorageHandler{storageMetrics: storageMetrics}
}

// UploadFile handles file uploads and tracks storage usage.
func (h *StorageHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "invalid file upload", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileSize, err := strconv.ParseInt(r.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		http.Error(w, "could not determine file size", http.StatusBadRequest)
		return
	}

	err = h.storageMetrics.AddFileSize(fileSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Upload success"))
}

// GetStorage returns the total storage used as JSON.
func (h *StorageHandler) GetStorage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	totalStorage := h.storageMetrics.GetTotalStorage()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"total_storage": totalStorage})
}
