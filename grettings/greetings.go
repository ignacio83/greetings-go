package grettings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	English    = iota
	Portuguese = iota
)

var formatsByLanguage = map[int][]string{
	English: {
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	},
	Portuguese: {
		"Olá, %v. Bem vindo!",
		"Bom te ver, %v!",
		"Oi, %v, espero que esteja tudo bem com você!",
	},
}

var ErrEmptyName = errors.New("empty name")

func Hello(name string, language int) (string, error) {
	if name == "" {
		return "", ErrEmptyName
	}
	message := fmt.Sprintf(randomGreeting(language), name)
	return message, nil
}

func Hellos(names []string, language int) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name, language)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomGreeting(language int) string {
	formats := formatsByLanguage[language]
	return formats[rand.Intn(len(formats))]
}
