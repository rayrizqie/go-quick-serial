package main

import (
    "fmt"
    "log"
    serialize "github.com/rayrizqie/go-quick-serial"
)

func main() {
    config := serialize.SerialCodeConfig{
        Pattern: "######-######",
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
}

