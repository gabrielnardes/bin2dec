package bin2dec

import (
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

func UserNotificationWhenInputIsDifferentOf0Or1(t *testing.T) {

}

func IfTheResultIsInASingleOuput(t *testing.T) {

}

func TestMustAllowOnlyNumbers(t *testing.T) {
    userInput := []string{
        "abcdefgh",
        "123defgh",
        "abc45678",
        "abc456gh",
    }
    for _, tt := range userInput {
        t.Run(tt, func(t *testing.T) {
            _, err := ValidateUserInput(tt)
            if err == nil {
                t.Fatal("Should warn the user if input not exclusively numeric.")
            }
        })
    }
}

func MustCountEmojisAsOneCharacter(t *testing.T) {

}
