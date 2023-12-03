package parse

import (
	"testing"
)

func TestIsDigit(t *testing.T) {
	if !IsDigit('5') {
		t.Errorf("IsDigit('5') = false, expected true")
	}
	if IsDigit('a') {
		t.Errorf("IsDigit('a') = true, expected false")
	}
}

func TestIsLower(t *testing.T) {
	if !IsLower('a') {
		t.Errorf("IsLower('a') = false, expected true")
	}
	if IsLower('A') {
		t.Errorf("IsLower('A') = true, expected false")
	}
}

func TestIsUpper(t *testing.T) {
	if !IsUpper('A') {
		t.Errorf("IsUpper('A') = false, expected true")
	}
	if IsUpper('a') {
		t.Errorf("IsUpper('a') = true, expected false")
	}
}

func TestIsLetter(t *testing.T) {
	if !IsLetter('a') {
		t.Errorf("IsLetter('a') = false, expected true")
	}
	if !IsLetter('A') {
		t.Errorf("IsLetter('A') = false, expected true")
	}
	if IsLetter('1') {
		t.Errorf("IsLetter('1') = true, expected false")
	}
}

func TestDigitToInt(t *testing.T) {
	if DigitToInt('5') != 5 {
		t.Errorf("DigitToInt('5') = %d, expected 5", DigitToInt('5'))
	}
}

func TestLowerToInt(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want int
	}{
		{"lowercase a", 'a', 0},
		{"lowercase z", 'z', 25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LowerToInt(tt.r); got != tt.want {
				t.Errorf("LowerToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperToInt(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want int
	}{
		{"uppercase A", 'A', 0},
		{"uppercase Z", 'Z', 25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpperToInt(tt.r); got != tt.want {
				t.Errorf("UpperToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLetterToInt(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want int
	}{
		{"lowercase a", 'a', 0},
		{"uppercase A", 'A', 0},
		{"lowercase z", 'z', 25},
		{"uppercase Z", 'Z', 25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LetterToInt(tt.r); got != tt.want {
				t.Errorf("LetterToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
