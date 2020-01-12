package problem029

import "testing"

func TestEncode(t *testing.T) {
	plainText := "AAAABBBCCDAA"
	encodedText := Encode(plainText)
	if encodedText != "4A3B2C1D2A" {
		t.Fail()
	}
}
