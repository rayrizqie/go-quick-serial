package serialize

import (
    "errors"
    "strings"
)

// Validator defines an interface for validation
type Validator interface {
    Validate(config SerialCodeConfig) error
}

// PatternValidator validates the pattern of the serial code
type PatternValidator struct{}

func (v PatternValidator) Validate(config SerialCodeConfig) error {
    if len(config.Pattern) < 6 {
        return errors.New("pattern must be at least 6 characters long")
    }

    containsPlaceholder := strings.Contains(config.Pattern, "#")
    if !containsPlaceholder {
        return errors.New("pattern must contain at least one '#' placeholder")
    }

    return nil
}

