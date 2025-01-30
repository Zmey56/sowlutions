package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

type NameConverter struct{}

func (nc NameConverter) toCamelCase(input string) string {
	words := strings.Split(input, "_")
	for i := 1; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}

func (nc NameConverter) toPascalCase(input string) string {
	words := strings.Split(input, "_")
	for i := 0; i < len(words); i++ {
		words[i] = toTitle(words[i])
	}
	return strings.Join(words, "")
}

func (nc NameConverter) toKebabCase(input string) string {
	re := regexp.MustCompile(`\s+`)
	words := re.Split(input, -1)
	return strings.ToLower(strings.Join(words, "-"))
}

func (nc NameConverter) toSnakeCase(input string) string {
	re := regexp.MustCompile(`\s+`)
	words := re.Split(input, -1)
	return strings.ToLower(strings.Join(words, "_"))
}

func toTitle(word string) string {
	if len(word) == 0 {
		return word
	}
	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}

func main() {
	nc := NameConverter{}

	input := "hello word"
	fmt.Println("Pascal Case:", nc.toPascalCase(input)) // Output: "HelloWord"
	fmt.Println("Kebab Case:", nc.toKebabCase(input))   // Output: "hello-word"
	fmt.Println("Snake Case:", nc.toSnakeCase(input))   // Output: "hello_word"
	fmt.Println("Camel Case:", nc.toCamelCase(input))   // Output: "helloWord"

	input2 := "hello_word"
	fmt.Println("Pascal Case:", nc.toPascalCase(input2)) // Output: "HelloWord"
	fmt.Println("Kebab Case:", nc.toKebabCase(input2))   // Output: "hello-word"
	fmt.Println("Snake Case:", nc.toSnakeCase(input2))   // Output: "hello_word"
	fmt.Println("Camel Case:", nc.toCamelCase(input2))   // Output: "helloWord"
}
