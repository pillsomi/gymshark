package storage

// Mock type of storage to use for unit testing.
type Mock struct {
	GetPackageSizesRes    []int
	GetPackagesSizeErr    error
	UpdatePackageSizesRes []int
	UpdatePackageSizesErr error
}

// GetPackageSizes returns the package sizes in storage.
func (s *Mock) GetPackageSizes() ([]int, error) {
	return s.GetPackageSizesRes, s.GetPackagesSizeErr
}

// UpdatePackageSizes update the package sizes in storage.
func (s *Mock) UpdatePackageSizes(_ []int) ([]int, error) {
	return s.UpdatePackageSizesRes, s.UpdatePackageSizesErr
}
