package arraystring

import "math"

// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/

func KidsWithCandies(candies []int, extraCandies int) []bool {
	greatest := math.MinInt

	for _, elem := range candies {
		if elem > greatest {
			greatest = elem
		}
	}

	res := make([]bool, len(candies))

	for i, elem := range candies {
		if elem+extraCandies >= greatest {
			res[i] = true
		}
	}
	return res
}
