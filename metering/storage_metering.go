package metering

import "errors"

// MaxStorage is the maximum storage allowed (1GB).
const MaxStorage = 1 << 30 

// StorageMetrics tracks the total storage used.
type StorageMetrics struct {
	totalStorage int64 // Total storage used in bytes.
}

// NewStorageMetrics creates a new StorageMetrics instance.
func NewStorageMetrics() *StorageMetrics {
	return &StorageMetrics{}
}

// AddFileSize adds the file size to the total storage.
// Returns an error if the storage limit is exceeded.
func (sm *StorageMetrics) AddFileSize(size int64) error {
	if sm.totalStorage+size > MaxStorage {
		return errors.New("storage limit exceeded")
	}
	sm.totalStorage += size
	return nil
}

// GetTotalStorage returns the total storage used.
func (sm *StorageMetrics) GetTotalStorage() int64 {
	return sm.totalStorage
}
