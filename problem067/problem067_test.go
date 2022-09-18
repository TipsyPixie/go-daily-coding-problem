package problem067

import (
	"os"
	"testing"
)

var cache *LfuCache = nil

func TestMain(m *testing.M) {
	cache = NewCache(4)
	os.Exit(m.Run())
}

func TestLfuCache_Set(t *testing.T) {
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	cache.Set("key3", "value3")
	cache.Set("key4", "value4")

	cache.Get("key1")
	cache.Get("key1")
	cache.Get("key1")

	cache.Set("key5", "value5")

	if _, exists := cache.dNodes["key1"]; !exists {
		t.Logf("cache key evicted: %s", "key1")
		t.FailNow()
	}
	if _, exists := cache.dNodes["key2"]; exists {
		t.Logf("cache key not evicted: %s", "key2")
		t.FailNow()
	}
}
