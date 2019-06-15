package main

import "testing"

func TestGreetTheAudience(t *testing.T) {
	expected := "Hello World!"
	actual := greetTheAudience()

	if actual != expected {
		t.Error("Expected", expected, ", but got", actual)
	}
}
