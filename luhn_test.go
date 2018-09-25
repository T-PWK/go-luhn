package luhn

import (
	"testing"
)

func BenchmarkChecksum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Checksum("0123456789")
	}
}

func TestIsValid(t *testing.T) {
	data := []struct {
		Input string
		Want  bool
	}{
		{"", false},
		{"1", false},
		{"18", true},
		{"79927398710", false},
		{"79927398712", false},
		{"79927398713", true},
		{"79927398714", false},
		{"2772502047190676862", true},
		{"49927398716", true},
		{"49927398717", false},
		{"1234567812345678", false},
		{"1234567812345670", true},
	}

	for _, d := range data {
		if d.Want != IsValid(d.Input) {
			t.Errorf("Incorrect Luhn checksum validation for %s, got %t, want %t", d.Input, !d.Want, d.Want)
		}
	}
}

func TestChecksum(t *testing.T) {
	data := []struct {
		Input, Want string
	}{
		{"7992739871", "3"},
		{"277250204719067686", "2"},
		{"123456781234567", "0"},
	}

	for _, d := range data {
		if c := Checksum(d.Input); c != d.Want {
			t.Errorf("Invalid Luhn checksum for %s, got %s, want %s", d.Input, c, d.Want)
		}
	}
}
