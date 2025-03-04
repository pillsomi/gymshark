package controllers

import (
	tdd "github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateBestNumberOfPackages(t *testing.T) {
	assert := tdd.New(t)

	minExtra1 := 750
	minTotalBoxes1 := 1
	found1 := false

	minExtra2 := 500000
	minTotalBoxes2 := 9434
	found2 := false

	tests := []struct {
		boxMap        map[int]int
		sizes         []int
		minExtraItems *int
		minTotalBoxes *int
		maximun       int
		order         int
		found         *bool
		bestResults   map[int]int
		expected      map[int]int
	}{
		{
			boxMap: map[int]int{
				250:  0,
				500:  0,
				1000: 1,
			},
			sizes:         []int{250, 500, 1000, 2000, 5000},
			minExtraItems: &minExtra1,
			minTotalBoxes: &minTotalBoxes1,
			maximun:       2,
			order:         750,
			found:         &found1,
			bestResults:   map[int]int{},
			expected: map[int]int{
				250:  1,
				500:  1,
				1000: 0,
			},
		}, {
			boxMap: map[int]int{
				23: 0,
				31: 0,
				53: 9434,
			},
			sizes:         []int{23, 31, 53},
			minExtraItems: &minExtra2,
			minTotalBoxes: &minTotalBoxes2,
			maximun:       2,
			order:         500000,
			found:         &found2,
			bestResults:   map[int]int{},
			expected: map[int]int{
				23: 2,
				31: 7,
				53: 9429,
			},
		},
	}

	for _, tt := range tests {
		calculateBestNumberOfPackages(
			tt.boxMap,
			tt.sizes,
			tt.minExtraItems,
			tt.minTotalBoxes,
			tt.maximun,
			tt.order,
			tt.found,
			tt.bestResults,
		)
		assert.Equal(tt.expected, tt.bestResults)
	}
}
