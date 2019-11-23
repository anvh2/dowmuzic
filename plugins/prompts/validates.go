package prompts

import (
	"errors"
	"strconv"
	"strings"
)

func validateEmptyStringInput(imput string) error {
	if len(strings.TrimSpace(imput)) < 1 {
		return errors.New("Input must not be empty")
	}
	return nil
}

func validateIntegerNumberInput(input string) error {
	_, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return errors.New("Invalid number")
	}
	return nil
}
