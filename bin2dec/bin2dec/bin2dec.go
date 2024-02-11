package bin2dec

import (
    "errors"
)

func ValidateUserInput(input int) (string, error) {
    numberOfDigits := lenLoop(input)

    if numberOfDigits != 8 {
        return "", errors.New("Input must a binary number of 8 digits")
    }

    return "", nil
}

func lenLoop(i int) int {
    if i == 0 {
        return 1
    }
    count := 0
    for i != 0 {
        i /= 10
        count++
    }
    return count
}
