package utils

// Converter is the interface that wraps basic Convert method
type Converter interface {
	Convert() ([]string, error)
}

// Convert convert chinese from simplicity to traditional or simplicity to traditional
func Convert(c Converter) ([]string, error) {
	return c.Convert()
}
