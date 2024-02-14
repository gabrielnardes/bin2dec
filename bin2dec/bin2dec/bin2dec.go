package bin2dec

import (
    "errors"
    "math"
    "strconv"
)

func ValidateUserInput(input string) (string, error) {
    numberOfDigits := lenLoop(input)

    if numberOfDigits != 8 {
        return "", errors.New("Input must a binary number of 8 digits")
    }
    _, err := strconv.Atoi(input)
    if err != nil {
        return "", errors.New("looks like not a number")
    }

    return "", nil
}

func Calculate(i float64) float64 {
    return math.Log2(i)
}

func lenLoop(i string) int {
    return len(i)
}
