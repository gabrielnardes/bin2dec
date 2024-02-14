package bin2dec

import (
	"fmt"
	"testing"
)

func TestGiveErrorIfInputIsLessThan8Digits(t *testing.T) {
    userInput := "10"
    _, err := ValidateUserInput(userInput)
    if err == nil {
        t.Fatal("Should warn the user if input is less than 8 digits")
    }
}

func TestGiveErrorIfInputIsGreaterThan8Digits(t *testing.T) {
    userInput := "101010101"
    _, err := ValidateUserInput(userInput)
    if err == nil {
        t.Fatal("Should warn the user if input is greater than 8 digits")
    }
}

func TestCase1OfLogCalculation(t *testing.T) {
    var userInput float64 = 4
    ouput := Calculate(userInput)
    if ouput != 2 {
        t.Fatal("Log2 of 4 must be 2.")
    }
}

func TestNotifyWhenInputIsDifferentOf0Or1(t *testing.T) {
    userInput := []string{
        "10101019",
        "90101010",
        "10109010",
    }
    for _, tt := range userInput {
        t.Run(tt, func(t *testing.T) {
            _, err := ValidateUserInput(tt)
            if err == nil {
                t.Fatal(fmt.Sprintf("Should warn the user if input not exclusively 1 or 0. Input: '%s'", tt))
            }
        })
    }
}

func IfTheResultIsInASingleOuput(t *testing.T) {

}

func MustCountEmojisAsOneCharacter(t *testing.T) {

}
