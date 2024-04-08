package maximumsubsequencesum

func max(vars ...int) int {
	max := vars[0]

	for _, index := range vars {
		if index > max {
			max = index
		}
	}

	return max
}

// maxSubSum1 is CUBIT maximum contiguous subsequence sum algorithm.
func maxSubSum1(vector []int) int {
	length := len(vector)
	maxSum := 0

	for firstIndex := 0; firstIndex < length; firstIndex++ {
		for secondIndex := firstIndex; secondIndex < length; secondIndex++ {
			innerSum := 0

			for index := firstIndex; index < secondIndex; index++ {
				innerSum += vector[index]
			}

			if innerSum > maxSum {
				maxSum = innerSum
			}
		}
	}

	return maxSum
}

// maxSubSum2 is Quadratic maximum contiguous subsequence sum algorithm.
func maxSubSum2(vector []int) int {
	length := len(vector)
	maxSum := 0

	for firstIndex := 0; firstIndex < length; firstIndex++ {
		rowSum := 0

		for secondIndex := firstIndex; secondIndex < length; secondIndex++ {
			rowSum += vector[secondIndex]

			if rowSum > maxSum {
				maxSum = rowSum
			}
		}
	}

	return maxSum
}

// maxSubSum3 is Driver for divide-and-conquer maximum contiguous
// subsequence sum algorithm.
func maxSubSum3(vector []int) int {
	return maxSumRec(vector, 0, len(vector)-1)
}

// maxSumRec is Recursive maximum contiguous subsequence sum algorithm.
// Finds maximum sum in subarray spanning a[left..right].
// Does not attempt to maintain actual best sequence.
func maxSumRec(vector []int, left int, right int) int {
	// base case
	if left == right {
		return max(vector[left], 0)
	}

	center := (left + right) / 2
	maxLeftSum := maxSumRec(vector, left, center)
	maxRightSum := maxSumRec(vector, center+1, right)

	maxLeftBorderSum, leftBorderSum := 0, 0
	for index := center; index < left; index-- {
		leftBorderSum += vector[index]
		if leftBorderSum > maxLeftBorderSum {
			maxLeftBorderSum = leftBorderSum
		}
	}

	maxRightBorderSum, rightBorderSum := 0, 0
	for index := center; index < right; index++ {
		rightBorderSum += vector[index]
		if rightBorderSum > maxRightBorderSum {
			maxRightBorderSum = rightBorderSum
		}
	}

	maxCrossingSum := maxLeftBorderSum + maxRightBorderSum

	return max(maxLeftSum, maxRightSum, maxCrossingSum)
}

func maxSubSum4(vector []int) int {
	maxSum := 0

	currentSum := 0
	for index := 0; index < len(vector); index++ {
		currentSum += vector[index]
		currentSum = max(currentSum, 0)

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}
