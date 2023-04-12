package greetings

import (
	"errors"
	"fmt"
	"math/rand"
    "time"
)

func Hello(name string) (string, error) {
	if name == "" {
        return "", errors.New("Parâmetro inválido")
    }

    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func init() {
    //DEPRECATED
	//rand.Seed(time.Now().UnixNano())
}


func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }
	//https://stackoverflow.com/questions/75597325/rand-seedseed-is-deprecated-how-to-use-newrandnewseed
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed));
    return formats[r.Intn(len(formats))]
}
