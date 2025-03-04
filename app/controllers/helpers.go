package controllers

// calculateBestNumberOfPackages returns the optimal combination of different package sizes for a given number of order.
func calculateBestNumberOfPackages(
	packagesMap map[int]int,
	sizes []int,
	minExtraItems *int,
	minTotalBoxes *int,
	maximum,
	order int,
	found *bool,
	bestResults map[int]int) {
	if maximum == -1 || *found == true {
		return
	}
	numberOfBoxes := 1

	if sizes[maximum] < order {
		numberOfBoxes = order / sizes[maximum]
		if order%sizes[maximum] != 0 {
			numberOfBoxes += 1
		}
	}
	extraItems := (numberOfBoxes * sizes[maximum]) - order
	packagesMap[sizes[maximum]] = numberOfBoxes

	// If there is no extra items, it means we found the combination of package size that can handle all the items.
	// We can return safely, without checking other combinations at this point.
	if extraItems == 0 {
		*found = true
		for key, value := range packagesMap {
			bestResults[key] = value
		}
		return
	} else {
		// Since we still have extra items
		// check if current combination contains less extra items.
		// In case it does, update the current combination as the best.
		if extraItems < *minExtraItems {
			*minExtraItems = extraItems
			totalBoxes := 0
			for _, boxItems := range packagesMap {
				totalBoxes += boxItems
			}
			*minTotalBoxes = totalBoxes
			for key, value := range packagesMap {
				bestResults[key] = value
			}
		} else if extraItems == *minExtraItems {
			// In case the extra items quantity is the same as the previous best,
			// we check the total number of boxes, and if the number of boxes is lower,
			// update the current combination as best.
			totalBoxes := 0
			for _, boxItems := range packagesMap {
				totalBoxes += boxItems
			}
			if totalBoxes < *minTotalBoxes {
				*minTotalBoxes = totalBoxes
				for key, value := range packagesMap {
					bestResults[key] = value
				}
			}
		}
		// otherwise, we continue checking the other combinations.
	}

	// Check all the possible combinations by removing one of the current maximum size package, in order to check
	// if we can find a combination to match all items, in smaller packages, or find a better combination.
	for {
		// if the number of current maximum size package is less than 1, it means that we tried every combination
		// if at least on package of the current maximum package, and we need to compare combinations of packages with
		// smaller size.
		// If also found the best combination, we don't need to continue searching more.
		if packagesMap[sizes[maximum]]-1 <= 0 || *found == true {
			break
		}
		packagesMap[sizes[maximum]] = packagesMap[sizes[maximum]] - 1
		currentOrder := order - packagesMap[sizes[maximum]]*sizes[maximum]
		// Otherwise, continue the logic of checking combinations with at least on the current maximum size package.
		calculateBestNumberOfPackages(packagesMap, sizes, minExtraItems, minTotalBoxes, maximum-1, currentOrder, found, bestResults)
	}

	// If found the best combination, we dont need to continue searching for more combinations.
	if *found == false {
		for i := 0; i <= maximum; i++ {
			packagesMap[sizes[i]] = 0
		}
		maximum -= 1
		// This step start searching combination, without packages of this current maximum size, or larger than that.
		calculateBestNumberOfPackages(packagesMap, sizes, minExtraItems, minTotalBoxes, maximum, order, found, bestResults)
	}
}
