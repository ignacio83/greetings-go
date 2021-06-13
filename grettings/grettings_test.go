package grettings

import (
	"github.com/pkg/errors"
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Should return greeting for Golang in English", func(t *testing.T) {
		const name = "Golang"
		want := regexp.MustCompile(`(Hi|Great|Hail).*(\b` + name + `\b)`)
		const language = English
		got, err := Hello(name, language)
		if !want.MatchString(got) || err != nil {
			t.Fatalf(`Hello(%d,%q) = %q, %v, want match for %q, nil`, language, name, got, err, want)
		}
	})

	t.Run("Should return greeting for Golang in Portuguese", func(t *testing.T) {
		const name = "Golang"
		want := regexp.MustCompile(`(Olá|Bom te ver|Oi).*(\b` + name + `\b)`)
		const language = Portuguese
		got, err := Hello("Golang", Portuguese)
		if !want.MatchString(got) || err != nil {
			t.Fatalf(`Hello(%d,%q) = %q, %v, want match for %q, nil`, language, name, got, err, want)
		}
	})

	t.Run("Should return error when name is empty", func(t *testing.T) {
		_, err := Hello("", English)
		assertError(t, err, ErrEmptyName)
	})
}

func TestHellos(t *testing.T) {
	t.Run("Should return greetings for Golang in English", func(t *testing.T) {
		const name1 = "Golang"
		want1 := regexp.MustCompile(`(Hi|Great|Hail).*(\b` + name1 + `\b)`)
		const name2 = "John"
		want2 := regexp.MustCompile(`(Hi|Great|Hail).*(\b` + name2 + `\b)`)
		const language = English
		got, err := Hellos(language, name1, name2)
		if err != nil {
			t.Fatalf(`Hellos(%d,%q,%q) = %q, %v`, language, name1, name2, got, err)
		}

		if !want1.MatchString(got[name1]) {
			t.Fatalf(`map[%q] = %q, nil want match for %q`, name1, got[name1], want1)
		}

		if !want2.MatchString(got[name2]) {
			t.Fatalf(`map[%q] = %q, nil want match for %q`, name2, got[name2], want2)
		}
	})

	t.Run("Should return error when some name is empty", func(t *testing.T) {
		const name1 = "Golang"
		const name2 = ""
		const language = English
		_, err := Hellos(language, name1, name2)
		assertError(t, errors.Cause(err), ErrEmptyName)
	})
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func BenchmarkHellos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Hellos(English, "Golang", "Benchmark")
		if err != nil {
			b.Fatalf("Error in benchmark")
		}
	}
}
