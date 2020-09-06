package problem052

import (
	"testing"
)

func TestCache_Get(t *testing.T) {
	c := NewCache(4)
	c.Set("A", "Avengers").Set("B", "Batman").Set("C", "Cyborg")

	value, correctValue := c.Get("A"), cacheValue("Avengers")
	if value == nil {
		t.Log("A value is nil")
		t.FailNow()
	}
	if *value != cacheValue(correctValue) {
		t.Logf("%s != %s", *value, correctValue)
		t.Fail()
	}
}

func TestCache_Set(t *testing.T) {
	c := NewCache(4)
	c.Set("A", "Avengers").
		Set("B", "Batman").
		Set("C", "Cyborg").
		Set("D", "Dormamu")

	value, correctValue := c.Get("A"), cacheValue("Avengers")
	if value == nil {
		t.Log("A value is nil")
		t.FailNow()
	}
	if *value != cacheValue(correctValue) {
		t.Logf("%s != %s", *value, correctValue)
		t.Fail()
	}

	c.Set("E", "Electra")
	value = c.Get("A")
	if value != nil {
		t.Log("A value not removed")
		t.FailNow()
	}
}
