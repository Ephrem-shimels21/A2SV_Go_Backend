package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestAskNameAndSubject(t *testing.T) {
	// Input with extra newlines
	input := "Ephrem\n2\n"
	reader := bufio.NewReader(strings.NewReader(input))

	name, noSubjects, err := askNameAndSubject(reader)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Debug output
	t.Logf("Name received: '%s'", name)
	t.Logf("Number of subjects received: %d", noSubjects)

	if name != "Ephrem" {
		t.Errorf("Expected name to be 'Ephrem', got '%s'", name)
	}
	if noSubjects != 2 {
		t.Errorf("Expected noSubjects to be 2, got %d", noSubjects)
	}
}

func TestCalculateAverageSubjects(t *testing.T) {
	// Input with extra newlines
	input := "Math\n90\nScience\n85\n"
	reader := bufio.NewReader(strings.NewReader(input))

	subjectGradeMap, err := calculateAverageSubjects(reader, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := map[string]int{"Math": 90, "Science": 85}
	for key, val := range expected {
		if subjectGradeMap[key] != val {
			t.Errorf("Expected %d for %s, got %d", val, key, subjectGradeMap[key])
		}
	}
}
