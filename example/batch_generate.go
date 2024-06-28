package main

import (
    "fmt"
    "log"
    serialize "github.com/rayrizqie/go-quick-serial"
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

    // Generate a batch of serial codes
    batchCount := 10
    serialCodes, err := generator.GenerateBatch(config, batchCount)
    if err != nil {
        log.Fatalf("Error generating batch of serial codes: %v", err)
    }
    fmt.Printf("Generated Batch of Serial Codes: %v\n", serialCodes)
}

