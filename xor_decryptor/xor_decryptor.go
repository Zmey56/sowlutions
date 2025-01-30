package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func XORDecrypt(ciphertext []int, key string) string {
	plaintext := make([]byte, len(ciphertext))
	keyBytes := []byte(key)
	keyLength := len(keyBytes)

	for i, val := range ciphertext {
		plaintext[i] = byte(val) ^ keyBytes[i%keyLength]
	}
	return string(plaintext)
}

func FindKey(ciphertext []int) string {
	possibleChars := "abcdefghijklmnopqrstuvwxyz"
	bestKey := ""
	bestScore := 0

	for _, a := range possibleChars {
		for _, b := range possibleChars {
			for _, c := range possibleChars {
				key := string([]byte{byte(a), byte(b), byte(c)})
				decrypted := XORDecrypt(ciphertext, key)
				score := countEnglishWords(decrypted)

				if score > bestScore {
					bestScore = score
					bestKey = key
				}
			}
		}
	}
	return bestKey
}

func countEnglishWords(text string) int {
	commonWords := []string{"the", "be", "to", "of", "and", "in", "that", "have", "it", "is", "with", "on"}
	score := 0
	lowerText := strings.ToLower(text)
	for _, word := range commonWords {
		if strings.Contains(lowerText, word) {
			score++
		}
	}
	return score
}

func main() {
	tmp_file := "p059_cipher.txt"
	data, err := os.ReadFile(tmp_file)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	text := strings.TrimSpace(string(data))
	parts := strings.Split(text, ",")
	ciphertext := make([]int, len(parts))
	for i, part := range parts {
		ciphertext[i], err = strconv.Atoi(part)
		if err != nil {
			fmt.Println("Error converting to integer", err)
			return
		}
	}

	// Step 1: Find the key
	key := FindKey(ciphertext)
	fmt.Println("Found key:", key)

	// Step 2: Decrypt the message using the key
	decryptedMessage := XORDecrypt(ciphertext, key)
	fmt.Println("Decrypted Message:", decryptedMessage)
}

// Output:
// Decrypted Message: An extract taken from the introduction of one of Euler's most celebrated papers, "De summis serierum
// reciprocarum" [On the sums of series of reciprocals]: I have recently found, quite unexpectedly, an elegant expression
// for the entire sum of this series 1 + 1/4 + 1/9 + 1/16 + etc., which depends on the quadrature of the circle, so that
// if the true sum of this series is obtained, from it at once the quadrature of the circle follows. Namely, I have found
// that the sum of this series is a sixth part of the square of the perimeter of the circle whose diameter is 1; or by putting
// the sum of this series equal to s, it has the ratio sqrt(6) multiplied by s to 1 of the perimeter to the diameter.
// I will soon show that the sum of this series to be approximately 1.644934066842264364; and from multiplying this number
// by six, and then taking the square root, the number 3.141592653589793238 is indeed produced, which expresses the perimeter
// of a circle whose diameter is 1. Following again the same steps by which I had arrived at this sum, I have discovered
// that the sum of the series 1 + 1/16 + 1/81 + 1/256 + 1/625 + etc. also depends on the quadrature of the circle. Namely,
// the sum of this multiplied by 90 gives the biquadrate (fourth power) of the circumference of the perimeter of a circle
// whose diameter is 1. And by similar reasoning I have likewise been able to determine the sums of the subsequent series
// in which the exponents are even numbers.
