package main

import (
	"testing"
)

func TestHappyRunReadLatest(t *testing.T) {
	li, err := ReadLatest("proxy.golang.org", "golang.org/x/text")

	if err != nil {
		t.Errorf("no error expected but caught one: %#v", err)
	}

	assertEquals("v0.3.7", li.Version, t)
}

func TestHappyRunReadInfo(t *testing.T) {
	li, err := ReadInfo("proxy.golang.org", "golang.org/x/text", "v0.3.7")

	if err != nil {
		t.Errorf("no error expected but caught one: %#v", err)
	}

	assertEquals("v0.3.7", li.Version, t)
}

func TestHappyRunReadList(t *testing.T) {
	lip, err := ReadList("proxy.golang.org", "golang.org/x/text")

	if err != nil {
		t.Errorf("no error expected but caught one: %#v", err)
	}

	var li = *lip

	assertEquals("v0.3.0", li[0], t)

}

func assertEquals(expected string, current string, t *testing.T) {

	if expected != current {
		t.Errorf("expected <%s>, but was <%s>", expected, current)
	}
}
