package serialize

// SerialCodeConfig holds the configuration for the serial code generator
type SerialCodeConfig struct {
    Pattern string
    Prefix  string
    Suffix  string
    Charset Charset
}

