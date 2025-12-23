package utils

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/google/uuid"
)

func ValidationRequired(field, value string) error {
	if value == "" {
		return fmt.Errorf("%s is required", field)
	}
	return nil
}

func ValidationStringLength(field, value string, min, max int) error {
	l := len(value)
	if l < min || l > max {
		return fmt.Errorf("%s must be between %d and %d characters", field, min, max)
	}
	return nil
}

func ValidationRegex(value string, re *regexp.Regexp, errMsg string) error {
	if !re.MatchString(value) {
		return fmt.Errorf("%s", errMsg)
	}
	return nil
}

func ValidationPositiveInt(field, value string) (int, error) {
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be a number", field)
	}

	if v <= 0 {
		return 0, fmt.Errorf("%s must be a positive number", field)
	}

	return v, nil
}

func ValidationUUID(field, value string) (uuid.UUID, error) {
	uid, err := uuid.Parse(value)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s must be a valid UUID", field)
	}
	return uid, nil
}

func ValidationInList(field, value string, allowed map[string]bool) error {
	if !allowed[value] {
		return fmt.Errorf("%s must be one of: %v", field, keys(allowed))
	}
	return nil
}

// convert map to slice
func keys(m map[string]bool) []string {
	var k []string
	for key := range m {
		k = append(k, key)
	}
	return k
}
