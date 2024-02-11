package bin2dec

import (
	"testing"
)

func TestGiveErrorIfInputIsLessThan8Digits(t *testing.T) {
    userInput := 10
    _, err := ValidateUserInput(userInput)
    if err == nil {
        t.Fatal("Should warn the user if input is less than 8 digits")
    }
}

func TestGiveErrorIfInputIsGreaterThan8Digits(t *testing.T) {
    userInput := 101010101
    _, err := ValidateUserInput(userInput)
    if err == nil {
        t.Fatal("Should warn the user if input is greater than 8 digits")
    }
}

func UserNotificationWhenInputIsDifferentOf0Or1(t *testing.T) {

}

func IfTheResultIsInASingleOuput(t *testing.T) {

}

func MustAllowOnlyNumbers(t *testing.T) {

}
