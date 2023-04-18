package main

import (
    "fmt"
    "log"
    "greetings"
)

func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Array de nome
    names := []string{"Fabricio", "Jessyca", "Doris"}
    
    message, err := greetings.Hello("Fabricio")
    
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(message)

    messages, err := greetings.Hellos(names)

    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(messages)
}