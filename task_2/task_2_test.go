package main

import (
	"testing"
)

func TestFrequencyCounter(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{"hello", map[string]int{"h": 1, "e": 1, "l": 2, "o": 1}},
		{"test", map[string]int{"t": 2, "e": 1, "s": 1}},
		{"", map[string]int{}},
		{"aabbcc", map[string]int{"a": 2, "b": 2, "c": 2}},
	}

	for _, test := range tests {
		result := frequencyCounter(test.input)
		if !compareMaps(result, test.expected) {
			t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestPalindromeChecker(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"madam", true},
		{"racecar", true},
		{"hello", false},
		{"", true},
		{"a", true},
	}

	for _, test := range tests {
		result := palindromeChecker(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected %t, but got %t", test.input, test.expected, result)
		}
	}
}

func compareMaps(map1, map2 map[string]int) bool {
	if len(map1) != len(map2) {
		return false
	}
	for key, value := range map1 {
		if map2[key] != value {
			return false
		}
	}
	return true
}
