package main

import (
	"testing"
)

func TestHappyRun(t *testing.T) {
	li, err := ReadLatest("proxy.golang.org", "golang.org/x/text")

	if err != nil {
		t.Errorf("no error expected but caught one: %#v", err)
	}

	var expected = "v0.3.7"
	if li.Version != expected {
		t.Errorf("expected <%s>, but was <%s>", li.Version, expected)
	}
}
