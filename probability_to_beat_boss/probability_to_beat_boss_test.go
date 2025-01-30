package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"testing"
)

func createTestCSV(filename string, data [][]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range data {
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}

func TestProbabilityToBeatBoss(t *testing.T) {
	filename := "test_prediction.csv"
	testData := [][]string{
		{"Card Suit", "Animal Name", "Fruit", "Result"},
		{"Hearts", "Lion", "Watermelon", "true"},
		{"Hearts", "Lion", "Watermelon", "false"},
		{"Hearts", "Lion", "Watermelon", "true"},
		{"Diamonds", "Fox", "Bananas", "false"},
	}

	if err := createTestCSV(filename, testData); err != nil {
		t.Fatalf("Failed to create test CSV file: %v", err)
	}
	expected := 66.67 // 2 out of 3 wins -> 66.67%
	result := ProbabilityToBeatBoss("Hearts", "Lion", "Watermelon", filename)

	if strconv.FormatFloat(result, 'f', 2, 64) != strconv.FormatFloat(expected, 'f', 2, 64) {
		t.Errorf("Expected probability %.2f%%, but got %.2f%%", expected, result)
	}
}

func TestProbabilityToBeatBoss_NoMatches(t *testing.T) {
	filename := "test_prediction.csv"
	testData := [][]string{
		{"Card Suit", "Animal Name", "Fruit", "Result"},
		{"Diamonds", "Fox", "Bananas", "false"},
	}

	if err := createTestCSV(filename, testData); err != nil {
		t.Fatalf("Failed to create test CSV file: %v", err)
	}

	expected := 0.0
	result := ProbabilityToBeatBoss("Hearts", "Lion", "Watermelon", filename)

	if result != expected {
		t.Errorf("Expected probability %.2f%%, but got %.2f%%", expected, result)
	}
}
