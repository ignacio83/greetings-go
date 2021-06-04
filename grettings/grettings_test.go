package grettings

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Should print greeting for Golang", func(t *testing.T) {
		const name = "Golang"
		want := regexp.MustCompile(`\b` + name + `\b`)
		got, err := Hello("Golang")
		if !want.MatchString(got) || err != nil {
			t.Fatalf(`Hello("Golang") = %q, %v, want match for %#q, nil`, got, err, want)
		}
	})

	t.Run("Should return error when name is empty", func(t *testing.T) {
		got, err := Hello("")
		if got != "" || err == nil {
			t.Fatalf(`Hello("") = %q, %v, want "", error`, got, err)
		}
	})
}
