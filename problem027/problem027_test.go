package problem027

import "testing"

func TestValidateBrackets(t *testing.T) {
	if !ValidateBrackets("([])") {
		t.Fail()
	}
	if ValidateBrackets("([)]") {
		t.Fail()
	}
	if ValidateBrackets("((()") {
		t.Fail()
	}
}
