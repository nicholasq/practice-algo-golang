package arraystring

// https://leetcode.com/problems/merge-strings-alternately
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
