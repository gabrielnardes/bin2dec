package main

import (
	"fmt"
	"os"

	"bin2dec/pkg/bin2dec"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: bin2dec <binary_number>")
		os.Exit(1)
	}
	arg := os.Args[1]

    _, err := bin2dec.ValidateUserInput(arg)
    if err != nil {
        fmt.Println(err)
        return
    }
    output := bin2dec.Calculate(arg)
	fmt.Println("Result:", output)
}
