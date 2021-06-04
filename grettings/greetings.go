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

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomGreeting(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

const welcome = "Hi, %v. Welcome!"
const greatSeeYou = "Great to see you, %v!"
const hail = "Hail, %v! Well met!"

func randomGreeting() string {
	formats := []string{
		welcome,
		greatSeeYou,
		hail,
	}
	return formats[rand.Intn(len(formats))]
}
