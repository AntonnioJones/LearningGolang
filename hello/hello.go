package main

import ( 
    "fmt"
    "../greetings/greetings.go"
)


func main() {
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}