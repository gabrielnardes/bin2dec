package bin2dec

import (
	"fmt"
	"testing"
)

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

func TestBin2DecConversion(t *testing.T) {
    userInput := map[string]int{
        "111111110101010": 32682,
        "10101010": 170,
        "000001": 1,
        "100000": 32,
        "000100": 4,
        "000000": 0,
    }
    for k, v := range userInput {
        t.Run(k, func(t *testing.T) {
            actualV := Calculate(k)
            if actualV != v {
                t.Fatal("Different Results. Actual=", actualV, ". Expected=", v)
            }
        })
    }
}
