package problem055

import "testing"

func TestShorten(t *testing.T) {
	url := "https://google.com"
	shortened := Shorten(url)
	shortened2 := Shorten(url)
	if shortened != shortened2 {
		t.Logf("%s != %s", shortened, shortened2)
		t.FailNow()
	}
}

func TestResolve(t *testing.T) {
	url := "https://google.com"
	shortened := Shorten(url)
	if Resolve(shortened) != url {
		t.Logf("%s != %s", shortened, url)
		t.FailNow()
	}
}
