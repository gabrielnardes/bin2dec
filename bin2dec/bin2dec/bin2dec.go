package bin2dec

import (
	"errors"
	"math"
	"regexp"
)

func ValidateUserInput(input string) (string, error) {
    numberOfDigits := lenLoop(input)

    if numberOfDigits != 8 {
        return "", errors.New("Input must a binary number of 8 digits")
    }

    if match, _ := regexp.MatchString("^[01]+$", input); !match {
        return "", errors.New("Input must be only 1 or 0")
    }

    return "", nil
}

func Calculate(i float64) float64 {
    return math.Log2(i)
}

func lenLoop(i string) int {
    return len(i)
}
