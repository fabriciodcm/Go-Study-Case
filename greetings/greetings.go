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

func Hellos(names []string) (map[string]string, error) {
    // key->Value
    messages := make(map[string]string)
   
    //foreach
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        
        messages[name] = message
    }
    return messages, nil
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
