package problem034

import "testing"

func TestClosestPalindrome(t *testing.T) {
	result := ClosestPalindrome("race")
	if result != "ecarace" {
		t.Fail()
	}

	result = ClosestPalindrome("google")
	if result != "elgoogle" {
		t.Fail()
	}
}
