package grettings

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Should return greeting for Golang in English", func(t *testing.T) {
		const name = "Golang"
		want := regexp.MustCompile(`(Hi|Great|Hail).*(\b` + name + `\b)`)
		got, err := Hello("Golang", English)
		if !want.MatchString(got) || err != nil {
			t.Fatalf(`Hello(%q) = %q, %v, want match for %#q, nil`, want, got, err, want)
		}
	})

	t.Run("Should return greeting for Golang in Portuguese", func(t *testing.T) {
		const name = "Golang"
		want := regexp.MustCompile(`(Olá|Bom te ver|Oi).*(\b` + name + `\b)`)
		got, err := Hello("Golang", Portuguese)
		if !want.MatchString(got) || err != nil {
			t.Fatalf(`Hello(%q) = %q, %v, want match for %#q, nil`, want, got, err, want)
		}
	})

	t.Run("Should return error when name is empty", func(t *testing.T) {
		got, err := Hello("", English)
		if got != "" || err == nil {
			t.Fatalf(`Hello("") = %q, %v, want "", error`, got, err)
		}
	})
}

func BenchmarkHellos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Hellos([]string{"Golang", "Benchmark"}, English)
		if err != nil {
			b.Fatalf("Error in benchmark")
		}
	}
}
