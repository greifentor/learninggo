package main

import (
	"testing"
)

func TestFizzBuzz3(t *testing.T) {
	expected := "fizz"
	returned := FizzBuzz(3)
	if expected != returned {
		t.Errorf("expected <%q> but was <%q>", expected, returned)
	}
}

func TestFizzBuzz5(t *testing.T) {
	expected := "buzz"
	returned := FizzBuzz(5)
	if expected != returned {
		t.Errorf("expected <%q> but was <%q>", expected, returned)
	}
}

func TestFizzBuzz7(t *testing.T) {
	expected := "7"
	returned := FizzBuzz(7)
	if expected != returned {
		t.Errorf("expected <%q> but was <%q>", expected, returned)
	}
}

func TestFizzBuzz15(t *testing.T) {
	expected := "fizzbuzz"
	returned := FizzBuzz(15)
	if expected != returned {
		t.Errorf("expected <%q> but was <%q>", expected, returned)
	}
}
