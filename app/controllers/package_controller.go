package controllers

import (
	"errors"
	"sort"

	"github.com/pillsomi/gymshark/app/storage"
)

// ErrorEmptyListInput list is empty validation.
var ErrorEmptyListInput = errors.New("please provide a non empty list of positive integers")

// ErrorNonPositiveIntegerListInput list contains non positive items validation.
var ErrorNonPositiveIntegerListInput = errors.New("please provide a valid list of positive integers")

// ErrorInvalidNumberOfItemsInput invalid number of items validation.
var ErrorInvalidNumberOfItemsInput = errors.New("invalid number of items")

// Controller contains a collection of functions to handle business logic.
type Controller interface {
	GetPackageSizes() ([]int, error)
	UpdatePackageSizes([]int) ([]int, error)
	CalculateNumberOfBoxes(int) ([]BoxDetail, error)
}

// BoxDetail contains detail of the box, size and number of items.
type BoxDetail struct {
	Size   int `json:"size"`
	Number int `json:"number"`
}

// PackageController implementation of Controller interface.
type PackageController struct {
	store storage.Storage
}

// New initiates a new instance of PackageController.
func New(store storage.Storage) *PackageController {
	return &PackageController{
		store: store,
	}
}

// GetPackageSizes returns the package sizes.
func (p *PackageController) GetPackageSizes() ([]int, error) {
	// Retrieve package sizes from storage.
	packageSizes, err := p.store.GetPackageSizes()
	if err != nil {
		return nil, err
	}
	return packageSizes, nil
}

// UpdatePackageSizes updates the package sizes.
func (p *PackageController) UpdatePackageSizes(packageSizes []int) ([]int, error) {
	// Validate input.
	if err := validatePackageSizes(packageSizes); err != nil {
		return nil, err
	}

	// Normalize input.
	normalizedInput := normalizePackageSizes(packageSizes)

	// Update package size in store.
	updated, err := p.store.UpdatePackageSizes(normalizedInput)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// Validate package sizes input.
func validatePackageSizes(packageSizes []int) error {
	if len(packageSizes) <= 0 {
		return ErrorEmptyListInput
	}
	for _, p := range packageSizes {
		if p <= 0 {
			return ErrorNonPositiveIntegerListInput
		}
	}
	return nil
}

// Normalize package size input.
func normalizePackageSizes(packageSizes []int) []int {
	// Deduplicate input.
	duplicateMap := make(map[int]struct{})
	for _, p := range packageSizes {
		duplicateMap[p] = struct{}{}
	}

	deduplicated := make([]int, 0)
	for packageSize, _ := range duplicateMap {
		deduplicated = append(deduplicated, packageSize)
	}

	sort.Ints(deduplicated)

	return deduplicated
}

// CalculateNumberOfBoxes calculates number of boxes.
func (p *PackageController) CalculateNumberOfBoxes(numberOfItems int) ([]BoxDetail, error) {
	// Validate input.
	if err := validateNumberOfItems(numberOfItems); err != nil {
		return nil, err
	}
	// Retrieve package size from store.
	packageSizes, err := p.store.GetPackageSizes()
	if err != nil {
		return nil, err
	}
	// Set initial extra items as number of items.
	initialExtra := numberOfItems
	extraItems := &initialExtra
	maximum := 0

	for i, ps := range packageSizes {
		if ps > numberOfItems {
			maximum = i
			break
		}
		maximum = i
	}

	// If maximum is 0, it means there are less items that fits in the package of the smallest size,
	// so we return the package of the smallest size, and just one package.
	if maximum == 0 {
		return []BoxDetail{
			{
				Size:   packageSizes[maximum],
				Number: 1,
			},
		}, nil
	}

	// contains different combinations of package sizes.
	packagesMap := make(map[int]int)
	for i := 0; i <= maximum; i++ {
		packagesMap[packageSizes[i]] = 0
	}

	// Calculate initial number of boxes taking in consideration that we will the largest possible package.
	numberOfBoxes := 1

	if packageSizes[maximum] < numberOfItems {
		numberOfBoxes = numberOfBoxes / packageSizes[maximum]
		if numberOfItems%packageSizes[maximum] != 0 {
			numberOfBoxes += 1
		}
	}
	minTotalBoxes := &numberOfBoxes

	// Found indicates if an optimal solution is found, meaning no extra items.
	var found *bool
	falseValue := false
	found = &falseValue
	bestResults := make(map[int]int)
	// Calculate the best choice for package number.
	calculateBestNumberOfPackages(
		packagesMap,
		packageSizes,
		extraItems,
		minTotalBoxes,
		maximum,
		numberOfItems,
		found,
		bestResults,
	)

	// Parse result.
	packages := make([]BoxDetail, 0)
	for packageSize, items := range bestResults {
		if items > 0 {
			packages = append(packages, BoxDetail{
				Size:   packageSize,
				Number: items,
			})
		}
	}

	sort.Slice(packages, func(i, j int) bool {
		return packages[i].Size > packages[j].Size
	})

	return packages, nil
}

// Validate number of items input.
func validateNumberOfItems(numberOfItems int) error {
	if numberOfItems <= 0 {
		return ErrorInvalidNumberOfItemsInput
	}
	return nil
}
