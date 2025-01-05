package casefmt

import (
	"strings"
	"unicode"
)

func CamelCaseToSnakeCase(s string) string {
	var result []rune
	for i, char := range s {
		if unicode.IsUpper(char) && i > 0 {
			result = append(result, '_', unicode.ToLower(char))
		} else {
			result = append(result, char)
		}
	}
	return strings.ToLower(string(result))
}
