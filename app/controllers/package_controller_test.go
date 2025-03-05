package controllers

import (
	"errors"
	"github.com/pillsomi/gymshark/app/storage"
	tdd "github.com/stretchr/testify/assert"
	"testing"
)

func TestPackageController_GetPackageSizes(t *testing.T) {
	assert := tdd.New(t)

	tests := []struct {
		name        string
		store       storage.Mock
		expectedErr error
		expected    []int
	}{
		{
			name: "storage error",
			store: storage.Mock{
				GetPackageSizesRes: nil,
				GetPackagesSizeErr: errors.New("error"),
			},
			expectedErr: errors.New("error"),
		}, {
			name: "success",
			store: storage.Mock{
				GetPackageSizesRes: []int{1, 2},
				GetPackagesSizeErr: nil,
			},
			expected: []int{1, 2},
		},
	}

	for _, tt := range tests {
		c := New(&tt.store)
		result, err := c.GetPackageSizes()
		if tt.expectedErr != nil {
			assert.Equal(err, tt.expectedErr, tt.name)
			assert.Nil(result, tt.name)
		} else {
			assert.Equal(result, tt.expected, tt.name)
			assert.Nil(err, tt.name)
		}
	}
}

func TestPackageController_UpdatePackageSizes(t *testing.T) {
	assert := tdd.New(t)

	tests := []struct {
		name        string
		input       []int
		store       storage.Mock
		expectedErr error
		expected    []int
	}{
		{
			name:        "invalid input, empty list",
			input:       []int{},
			expectedErr: ErrorEmptyListInput,
		}, {
			name:        "invalid input, non positive integers",
			input:       []int{-1},
			expectedErr: ErrorNonPositiveIntegerListInput,
		}, {
			name:  "storage error",
			input: []int{1, 2},
			store: storage.Mock{
				UpdatePackageSizesRes: nil,
				UpdatePackageSizesErr: errors.New("error"),
			},
			expectedErr: errors.New("error"),
		}, {
			name:  "success",
			input: []int{1, 2},
			store: storage.Mock{
				UpdatePackageSizesRes: []int{1, 2},
				UpdatePackageSizesErr: nil,
			},
			expected: []int{1, 2},
		},
	}

	for _, tt := range tests {
		c := New(&tt.store)
		result, err := c.UpdatePackageSizes(tt.input)
		if tt.expectedErr != nil {
			assert.Equal(err, tt.expectedErr, tt.name)
			assert.Nil(result, tt.name)
		} else {
			assert.Equal(result, tt.expected, tt.name)
			assert.Nil(err, tt.name)
		}
	}
}

func TestPackageController_CalculateNumberOfBoxes(t *testing.T) {
	assert := tdd.New(t)

	tests := []struct {
		name        string
		input       int
		store       storage.Mock
		expectedErr error
		expected    []BoxDetail
	}{
		{
			name:        "invalid input",
			input:       0,
			expectedErr: ErrorInvalidNumberOfItemsInput,
		}, {
			name:  "storage error",
			input: 5,
			store: storage.Mock{
				GetPackageSizesRes: nil,
				GetPackagesSizeErr: errors.New("error"),
			},
			expectedErr: errors.New("error"),
		}, {
			name:  "success",
			input: 5,
			store: storage.Mock{
				GetPackageSizesRes: []int{1, 2, 6},
				GetPackagesSizeErr: nil,
			},
			expected: []BoxDetail{
				{
					Size:   2,
					Number: 2,
				}, {
					Size:   1,
					Number: 1,
				},
			},
		},
	}

	for _, tt := range tests {
		c := New(&tt.store)
		result, err := c.CalculateNumberOfBoxes(tt.input)
		if tt.expectedErr != nil {
			assert.Equal(err, tt.expectedErr, tt.name)
			assert.Nil(result, tt.name)
		} else {
			assert.Equal(result, tt.expected, tt.name)
			assert.Nil(err, tt.name)
		}
	}
}
