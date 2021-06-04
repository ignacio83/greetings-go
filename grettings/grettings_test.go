package grettings

import (
	"regexp"
	"testing"
)

func TestShouldPrintGreetingForGolangWhenGolangIsInput(t *testing.T) {
	const name = "Golang"
	want := regexp.MustCompile(`\b` + name + `\b`)
	got, err := Hello("Golang")
	if !want.MatchString(got) || err != nil {
		t.Fatalf(`Hello("Golang") = %q, %v, want match for %#q, nil`, got, err, want)
	}
}

func TestShouldReturnErrorWhenNameIsEmpty(t *testing.T) {
	got, err := Hello("")
	if got != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, got, err)
	}
}
