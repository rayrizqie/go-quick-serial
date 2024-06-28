# Go-QuickSerial

**Lightweight and efficient Go package for generating unique serial codes**

Go-QuickSerial is designed to help you generate unique serial codes with customizable patterns, prefixes, suffixes, and character sets. Ideal for applications like vouchers, product keys, and access codes, Go-QuickSerial ensures each generated code is unique and secure.

## Features

- **Customizable Patterns**: Define your own patterns with placeholders for random characters.
- **Configurable Prefixes and Suffixes**: Add static text before or after the generated code.
- **Multiple Character Sets**: Choose from alphanumeric, alphabetic, or numeric character sets.
- **Batch Generation**: Generate multiple unique codes concurrently.
- **Concurrency Support**: Efficiently generate codes using Go's concurrency features.

## Installation

Install the package using the following command:

```sh
go get -u github.com/rayrizqie/go-quick-serial
```

## Usage

### Config Reference

- Pattern (string): Defines the pattern for the serial code. Use # as a placeholder for random characters. Example: "##-##-##".

- Prefix (string): A static prefix added before the generated code. Example: "PRE-".

- Suffix (string): A static suffix added after the generated code. Example: "-SUF".

- Charset (Charset): Defines the character set used for generating random characters. Available options:
    - AlphanumericCharset{}: Uses uppercase letters and digits.
    - AlphabetCharset{}: Uses uppercase letters only.
    - NumberCharset{}: Uses digits only.


### Generating a Single Serial Code

``` go
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

    // Generate a single serial code
    serialCode, err := generator.Generate(config)
    if err != nil {
        log.Fatalf("Error generating serial code: %v", err)
    }
    fmt.Printf("Generated Serial Code: %s\n", serialCode)
}

```

### Generating a Batch of Serial Codes

``` go
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
    batchCount := 5
    serialCodes, err := generator.GenerateBatch(config, batchCount)
    if err != nil {
        log.Fatalf("Error generating batch of serial codes: %v", err)
    }
    fmt.Printf("Generated Batch of Serial Codes: %v\n", serialCodes)
}

```

## Example Project

An example project is included in the example directory to demonstrate how to use the package.

To run the examples, navigate to the example directory and use the go run command:

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request with your improvements.


