package services

import (
	"math/rand"
	"strings"
	"time"

	"github.com/flosch/pongo2/v6"
)

const (
	digitsChars       = "0123456789"
	digits2Chars      = "123456789" // For FilterGenerateDigits2, excluding '0'
	alphaChars        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphaLowerChars   = "abcdefghijklmnopqrstuvwxyz"
	alphaUpperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	hexChars          = "0123456789abcdef"
	alphaNumericChars = digitsChars + alphaChars
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// generateRandomChars generates a random string of a given length from a given character set.
func generateRandomChars(length int, charset string) string {
	if length <= 0 {
		return ""
	}
	sb := strings.Builder{}
	sb.Grow(length)
	for i := 0; i < length; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

// FilterGenerateString generates a random string of a specified length (default 10).
// Usage: {{ ""|generate_string }} or {{ ""|generate_string:15 }}
func FilterGenerateString(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	length := 10
	if param.IsInteger() {
		length = param.Integer()
	} else if param.IsString() && param.String() != "" { // Allow length to be passed as string
		l, err := pongo2.Value{param.String()}.AsInteger()
		if err == nil {
			length = l
		}
	}
	return pongo2.AsValue(generateRandomChars(length, alphaNumericChars)), nil
}

// FilterGenerateAlphaNumeric generates a random alphanumeric string.
func FilterGenerateAlphaNumeric(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	length := param.Integer()
	if length == 0 { length = 10 } // Default length
	return pongo2.AsValue(generateRandomChars(length, alphaNumericChars)), nil
}

// FilterGenerateHex generates a random hexadecimal string.
func FilterGenerateHex(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	length := param.Integer()
	if length == 0 { length = 10 }
	return pongo2.AsValue(generateRandomChars(length, hexChars)), nil
}

// FilterGenerateDigits generates a random string of digits.
func FilterGenerateDigits(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	length := param.Integer()
	if length == 0 { length = 10 }
	return pongo2.AsValue(generateRandomChars(length, digitsChars)), nil
}

// FilterGenerateDigits2 generates a random string of digits (1-9).
func FilterGenerateDigits2(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	length := param.Integer()
	if length == 0 { length = 10 }
	return pongo2.AsValue(generateRandomChars(length, digits2Chars)), nil
}

// FilterGenerateAlpha generates a random alphabetic string.
func FilterGenerateAlpha(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	length := param.Integer()
	if length == 0 { length = 10 }
	return pongo2.AsValue(generateRandomChars(length, alphaChars)), nil
}

// FilterGenerateAlphaUpper generates a random uppercase alphabetic string.
func FilterGenerateAlphaUpper(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	length := param.Integer()
	if length == 0 { length = 10 }
	return pongo2.AsValue(generateRandomChars(length, alphaUpperChars)), nil
}

// FilterGenerateAlphaLower generates a random lowercase alphabetic string.
func FilterGenerateAlphaLower(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	length := param.Integer()
	if length == 0 { length = 10 }
	return pongo2.AsValue(generateRandomChars(length, alphaLowerChars)), nil
}
