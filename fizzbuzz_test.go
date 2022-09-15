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

func TestFizzBuzzBatch(t *testing.T) {
	expected := []string{"1", "fizz", "buzz", "7", "fizzbuzz"}
	passed := []int{1, 3, 5, 7, 15}
	for i := 0; i < len(passed); i++ {
		returned := FizzBuzz(passed[i])
		if expected[i] != returned {
			t.Errorf("test position %q: expected <%q> but was <%q>", i, expected[i], returned)
		}
		t.Logf("passed step %d", i)
	}
}
