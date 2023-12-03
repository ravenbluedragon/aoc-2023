package parse

// Common functions for parsing runes

// IsDigit returns true if the rune is a digit
func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

// IsLower returns true if the rune is a lowercase letter
func IsLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

// IsUpper returns true if the rune is an uppercase letter
func IsUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

// IsLetter returns true if the rune is a letter
func IsLetter(r rune) bool {
	return IsLower(r) || IsUpper(r)
}

// Common functions to convert runes to integers

// DigitToInt converts a rune to an integer
func DigitToInt(r rune) int {
	return int(r - '0')
}

// LowerToInt converts a lowercase letter to an integer
func LowerToInt(r rune) int {
	return int(r - 'a')
}

// UpperToInt converts an uppercase letter to an integer
func UpperToInt(r rune) int {
	return int(r - 'A')
}

// LetterToInt converts a letter to an integer
func LetterToInt(r rune) int {
	if IsLower(r) {
		return LowerToInt(r)
	}
	return UpperToInt(r)
}
