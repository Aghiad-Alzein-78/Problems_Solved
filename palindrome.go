func isPalindrome(x int) bool {
    if x < 0 {
		return false
	}
	original := x
	result := 0

	for x > 0 {
		remain := x % 10
		x = x / 10
		result = result*10 + remain
	}
	if original == result {
		return true
	}
	return false
}
