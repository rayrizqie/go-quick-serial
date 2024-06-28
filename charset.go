package serialize

import (
    "crypto/rand"
)

// Charset defines an interface for character set
type Charset interface {
    GenerateRandomCharacters(length int) (string, error)
}

// AlphanumericCharset implements the Charset interface for alphanumeric characters
type AlphanumericCharset struct{}

func (a AlphanumericCharset) GenerateRandomCharacters(length int) (string, error) {
    const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    return generateRandomCharacters(length, charset)
}

// AlphabetCharset implements the Charset interface for alphabetic characters
type AlphabetCharset struct{}

func (a AlphabetCharset) GenerateRandomCharacters(length int) (string, error) {
    const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    return generateRandomCharacters(length, charset)
}

// NumberCharset implements the Charset interface for numeric characters
type NumberCharset struct{}

func (n NumberCharset) GenerateRandomCharacters(length int) (string, error) {
    const charset = "0123456789"
    return generateRandomCharacters(length, charset)
}

// generateRandomCharacters generates random characters from a given charset
func generateRandomCharacters(length int, charset string) (string, error) {
    bytes := make([]byte, length)
    _, err := rand.Read(bytes)
    if err != nil {
        return "", err
    }
    for i, b := range bytes {
        bytes[i] = charset[b%byte(len(charset))]
    }
    return string(bytes), nil
}

