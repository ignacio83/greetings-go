package grettings

import (
	"regexp"
	"testing"
)

func TestShouldPrintGreetingForGolangWhenGolangIsInput(t *testing.T) {
	name := "Golang"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Golang")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Golang") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestShouldReturnErrorWhenNameIsEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
