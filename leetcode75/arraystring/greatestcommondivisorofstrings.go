package arraystring

// https://leetcode.com/problems/greatest-common-divisor-of-strings/description/?envType=study-plan-v2&envId=leetcode-75

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func GcdOfStrings(str1 string, str2 string) string {

	if str1+str2 != str2+str1 {
		return ""
	}

	gcdLen := gcd(len(str1), len(str2))

	return str1[:gcdLen]
}
