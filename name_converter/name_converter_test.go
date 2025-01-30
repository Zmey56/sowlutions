package main

import (
	"testing"
)

func TestNameConverter(t *testing.T) {
	nc := NameConverter{}

	tests := []struct {
		method   string
		input    string
		expected string
	}{
		{"toCamelCase", "hello_word", "helloWord"},
		{"toPascalCase", "hello_word", "HelloWord"},
		{"toKebabCase", "hello word", "hello-word"},
		{"toSnakeCase", "hello word", "hello_word"},
		{"toCamelCase", "foo_bar_baz", "fooBarBaz"},
		{"toPascalCase", "foo_bar_baz", "FooBarBaz"},
		{"toKebabCase", "foo bar baz", "foo-bar-baz"},
		{"toSnakeCase", "foo bar baz", "foo_bar_baz"},
	}

	for _, test := range tests {
		var result string
		switch test.method {
		case "toCamelCase":
			result = nc.toCamelCase(test.input)
		case "toPascalCase":
			result = nc.toPascalCase(test.input)
		case "toKebabCase":
			result = nc.toKebabCase(test.input)
		case "toSnakeCase":
			result = nc.toSnakeCase(test.input)
		}

		if result != test.expected {
			t.Errorf("%s(%q) = %q; want %q", test.method, test.input, result, test.expected)
		}
	}
}
