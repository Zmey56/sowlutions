package main

import (
	"testing"
)

func TestXORDecrypt(t *testing.T) {
	ciphertext := []int{104, 101, 108, 108, 111, 119, 111, 114, 108, 100} // Encrypted "helloworld"
	key := "abc"

	expected := "helloworld"
	result := XORDecrypt(ciphertext, key)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestFindKey(t *testing.T) {
	ciphertext := []int{104, 101, 108, 108, 111, 119, 111, 114, 108, 100}
	expectedKey := "abc"
	result := FindKey(ciphertext)

	if result != expectedKey {
		t.Errorf("Expected key %s, got %s", expectedKey, result)
	}
}

func TestCountEnglishWords(t *testing.T) {
	text := "the quick brown fox jumps over the lazy dog"
	expected := 3 // "the" appears twice, "in" appears once
	result := countEnglishWords(text)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
