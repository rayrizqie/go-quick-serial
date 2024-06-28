package main

import (
    "fmt"
    "log"
    "github.com/rayrizqie/go-quick-serial/serialize"
)

func main() {
    config := serialize.SerialCodeConfig{
        Pattern: "##-##-##",
        Prefix:  "PRE-",
        Suffix:  "-SUF",
        Charset: serialize.AlphanumericCharset{},
    }

    generator := serialize.SerialCodeGenerator{
        Validator: serialize.PatternValidator{},
    }

    // Generate a single serial code
    serialCode, err := generator.Generate(config)
    if err != nil {
        log.Fatalf("Error generating serial code: %v", err)
    }
    fmt.Printf("Generated Serial Code: %s\n", serialCode)

    // Generate a batch of serial codes
    batchCount := 5
    serialCodes, err := generator.GenerateBatch(config, batchCount)
    if err != nil {
        log.Fatalf("Error generating batch of serial codes: %v", err)
    }
    fmt.Printf("Generated Batch of Serial Codes: %v\n", serialCodes)
}

