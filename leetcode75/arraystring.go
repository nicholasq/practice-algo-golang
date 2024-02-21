package leetcode75

import (
	"math"
)

// MergeAlternately https://leetcode.com/problems/merge-strings-alternately
func MergeAlternately(word1 string, word2 string) string {
	l1, l2 := len(word1), len(word2)
	count := 0
	result := ""

	for count < l1 && count < l2 {
		result += string(word1[count]) + string(word2[count])
		count++
	}

	first, second := count, count

	for first < l1 {
		result += string(word1[first])
		first++
	}

	for second < l2 {
		result += string(word2[second])
		second++
	}

	return result
}

// CanPlaceFlowers https://leetcode.com/problems/can-place-flowers
func CanPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	prevIsEmpty := true
	for i, middle := range flowerbed {
		if middle == 0 {
			leftPlotEmpty := prevIsEmpty || i == 0
			rightPlotEmpty := i == len(flowerbed)-1 || flowerbed[i+1] == 0
			if leftPlotEmpty && rightPlotEmpty {
				count++
				prevIsEmpty = false
			} else {
				prevIsEmpty = true
			}

		} else {
			prevIsEmpty = false
		}
		if count == n {
			return true
		}
	}
	return n == 0
}

// KidsWithCandies https://leetcode.com/problems/kids-with-the-greatest-number-of-candies
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

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

// GcdOfStrings https://leetcode.com/problems/greatest-common-divisor-of-strings
func GcdOfStrings(str1 string, str2 string) string {

	if str1+str2 != str2+str1 {
		return ""
	}

	gcdLen := gcd(len(str1), len(str2))

	return str1[:gcdLen]
}

func isVowel(c rune) bool {
	switch c {
	case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}

// ReverseVowels https://leetcode.com/problems/reverse-vowels-of-a-string
func ReverseVowels(s string) string {
	size := len(s)
	start, end := 0, size-1
	runes := []rune(s)

	for start < end {
		left := runes[start]
		right := runes[end]
		leftIsVowel := isVowel(left)
		rightIsVowel := isVowel(right)

		if leftIsVowel && rightIsVowel {
			runes[start], runes[end] = right, left
			start++
			end--
			continue
		}

		if !rightIsVowel {
			end--
		}
		if !leftIsVowel {
			start++
		}
	}
	return string(runes)
}
