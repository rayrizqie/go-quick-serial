package serialize

import (
    "strings"
    "sync"
)

// SerialCodeGenerator generates serial codes based on configuration
type SerialCodeGenerator struct {
    Validator Validator
}

func (g SerialCodeGenerator) Generate(config SerialCodeConfig) (string, error) {
    if err := g.Validator.Validate(config); err != nil {
        return "", err
    }

    var code strings.Builder
    placeholderCount := 0

    for _, char := range config.Pattern {
        if char == '#' {
            placeholderCount++
        }
    }

    randomChars, err := config.Charset.GenerateRandomCharacters(placeholderCount)
    if err != nil {
        return "", err
    }

    randomIndex := 0
    for _, char := range config.Pattern {
        if char == '#' {
            code.WriteByte(randomChars[randomIndex])
            randomIndex++
        } else {
            code.WriteByte(byte(char))
        }
    }

    return config.Prefix + code.String() + config.Suffix, nil
}

func (g SerialCodeGenerator) GenerateBatch(config SerialCodeConfig, count int) ([]string, error) {
    var mu sync.Mutex
    codes := make(map[string]struct{})
    var wg sync.WaitGroup

    // Create a channel to collect results
    resultChan := make(chan string, count)
    errChan := make(chan error, count)

    for i := 0; i < count; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for {
                code, err := g.Generate(config)
                if err != nil {
                    errChan <- err
                    return
                }
                mu.Lock()
                if _, exists := codes[code]; !exists {
                    codes[code] = struct{}{}
                    resultChan <- code
                    mu.Unlock()
                    return
                }
                mu.Unlock()
            }
        }()
    }

    // Wait for all go routines to finish
    go func() {
        wg.Wait()
        close(resultChan)
        close(errChan)
    }()

    // Collect results
    var resultCodes []string
    for code := range resultChan {
        resultCodes = append(resultCodes, code)
    }

    // Check for errors
    if len(errChan) > 0 {
        return nil, <-errChan
    }

    return resultCodes, nil
}

