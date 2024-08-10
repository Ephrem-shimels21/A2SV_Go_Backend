package main

import (
	"testing"
)

func TestFrequencyCounter(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{"hello world", map[string]int{"hello": 1, "world": 1}},
		{"test TEST", map[string]int{"test": 2}},
		{"", map[string]int{}},
		{"aabb cc", map[string]int{"aabb": 1, "cc": 1}},
		{"hello, world!", map[string]int{"hello": 1, "world": 1}},
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
		{"Madam", true},
		{"Racecar", true},
		{"Hello", false},
		{"", true},
		{"A man, a plan, a canal, Panama", true},
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
