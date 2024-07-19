package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"racecar", true},
		{"hello", false},
		{"csdd", false},
		{"", true},
	}

	for _, test := range tests {
		result := IsPalindrome(test.input)
		if result != test.expected {
			t.Errorf("IsPalindrome(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestCountVowels(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello", 2},
		{"world", 1},
		{"aeiou", 5},
		{"HELLO", 2},
		{"", 0},
	}

	for _, test := range tests {
		result := CountVowels(test.input)
		if result != test.expected {
			t.Errorf("CountVowels(%q) = %d; want %d", test.input, result, test.expected)
		}
	}
}
