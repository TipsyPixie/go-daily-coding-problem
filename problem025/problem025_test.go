package problem025

import (
	"fmt"
	"regexp"
	"testing"
)

func checkRegex(pattern string, s string) bool {
	result, _ := regexp.MatchString(fmt.Sprintf("^%s$", pattern), s)
	return TestRegex(pattern, s) == result
}

func TestTestRegex(t *testing.T) {
	if !checkRegex(".*at", "chat") {
		t.Fail()
	}
	if !checkRegex("ra.", "ray") {
		t.Fail()
	}
	if !checkRegex("ra.", "raymond") {
		t.Fail()
	}
	if !checkRegex(".*", "quickbrownfox") {
		t.Fail()
	}
	if !checkRegex("a*b*c*", "c") {
		t.Fail()
	}
	if !checkRegex("a*b*c*", "") {
		t.Fail()
	}
}
