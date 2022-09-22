package main

import (
	"testing"
)

func TestHappyRunReadLatest(t *testing.T) {
	li, err := ReadLatest("proxy.golang.org", "golang.org/x/text")

	if err != nil {
		t.Errorf("no error expected but caught one: %#v", err)
	}

	var expected = "v0.3.7"
	if li.Version != expected {
		t.Errorf("expected <%s>, but was <%s>", li.Version, expected)
	}
}

func TestHappyRunReadInfo(t *testing.T) {
	li, err := ReadInfo("proxy.golang.org", "golang.org/x/text", "v0.3.7")

	if err != nil {
		t.Errorf("no error expected but caught one: %#v", err)
	}

	var expected = "v0.3.7"
	if li.Version != expected {
		t.Errorf("expected <%s>, but was <%s>", li.Version, expected)
	}
}

func TestHappyRunReadList(t *testing.T) {
	lip, err := ReadList("proxy.golang.org", "golang.org/x/text")

	if err != nil {
		t.Errorf("no error expected but caught one: %#v", err)
	}

	var expected = "v0.3.0"

	var li = *lip

	if li[0] != expected {
		t.Errorf("expected <%s>, but was <%s>", li[0], expected)
	}

}
