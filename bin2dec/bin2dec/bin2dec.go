package bin2dec

import (
	"errors"
	"math"
	"regexp"
)

func ValidateUserInput(input string) (string, error) {
    if match, _ := regexp.MatchString("^[01]+$", input); !match {
        return "", errors.New("Input must be only 1 or 0")
    }

    return "", nil
}

func Calculate(i float64) float64 {
    return math.Log2(i)
}

