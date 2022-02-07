package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func Greeting(name string) (string, error) {
	if len(name) == 0 {
		return "", errors.New("an empty name has been provided")
	}
	numbersRegexp := regexp.MustCompile("[0-9]+")
	if numbersRegexp.Match([]byte(name)) {
		return "", errors.New("name must not contain numbers")
	}
	greetingMsg := fmt.Sprintf(randomFormat(), name)
	return greetingMsg, nil
}

func Greetings(names []string) (map[string]string, error) {
	greetings := make(map[string]string)
	for _, name := range names {
		greetingMsg, error := Greeting(name)
		if error != nil {
			return nil, error
		}
		greetings[name] = greetingMsg
	}
	return greetings, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %s. Welcome!",
		"Great to see you, %s!",
		"Buongiorno, %s!",
		"Buona giornata, caro %s!",
	}
	return formats[rand.Intn(len(formats))]
}
