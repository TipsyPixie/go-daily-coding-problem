package problem022

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	words := []string{"brown", "the", "fox", "quick"}
	sentence := "thequickbrownfox"
	result := Run(words, sentence)
	if !reflect.DeepEqual(result, []string{"the", "quick", "brown", "fox"}) {
		t.Log(result)
		t.FailNow()
	}
}

func TestRun2(t *testing.T) {
	words := []string{"bed", "bath", "bedbath", "and", "beyond"}
	sentence := "bedbathandbeyond"
	result := Run(words, sentence)
	if !reflect.DeepEqual(result, []string{"bed", "bath", "and", "beyond"}) {
		t.Log(result)
		t.FailNow()
	}
}
