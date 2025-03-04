package storage

import "sync"

// Storage interface contains functions to get and update packages sizes.
type Storage interface {
	UpdatePackageSizes([]int) ([]int, error)
	GetPackageSizes() ([]int, error)
}

// InMemoryStorage a simple in memory implementation of storage interface.
type InMemoryStorage struct {
	mu          sync.Mutex
	dataStorage []int
}

// New initializes a new InMemoryStorage with default values for packages size.
func New() *InMemoryStorage {
	return &InMemoryStorage{
		dataStorage: []int{250, 500, 1000, 2000, 5000},
	}
}

// GetPackageSizes returns the package sizes in storage.
func (s *InMemoryStorage) GetPackageSizes() ([]int, error) {
	return s.dataStorage, nil
}

// UpdatePackageSizes update the package sizes in storage.
func (s *InMemoryStorage) UpdatePackageSizes(packageSizes []int) ([]int, error) {
	s.mu.Lock()
	s.dataStorage = packageSizes
	s.mu.Unlock()
	return s.dataStorage, nil
}
