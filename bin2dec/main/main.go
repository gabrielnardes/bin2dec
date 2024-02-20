package main

import (
    "fmt"
    "bin2dec"
)

func main() {
    var input string
    fmt.Println("Enter a binary number of 8 digits to convert it to decimal")
    fmt.Scan(&input)
    _, err := bin2dec.ValidateUserInput(input)
    if err != nil {
        fmt.Println(err)
        return
    }
    output := bin2dec.Calculate(input)
    fmt.Printf("Binary %s equal %d in decimal\n", input, output)
}
