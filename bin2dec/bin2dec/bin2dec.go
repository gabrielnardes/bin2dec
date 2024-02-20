package bin2dec

import (
	"errors"
	"math"
	"regexp"
	"strconv"
)

func ValidateUserInput(input string) (string, error) {
    if match, _ := regexp.MatchString("^[01]+$", input); !match {
        return "", errors.New("Input must be only 1 or 0")
    }

    return "", nil
}

func Calculate(input string) int {
    pow := len(input) - 1
    result := 0
    for _, digit := range input {
        integer, _ := strconv.Atoi(string(digit))
        result += integer * int(math.Pow(2, float64(pow)))
        pow--
    }
    return result
}

