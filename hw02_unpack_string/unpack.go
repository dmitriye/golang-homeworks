package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	// fmt.Printf(" Start to unpacking string: %s\n", str)

	if str == "" {
		return "", nil
	}

	runes := []rune(str)
	count := len(runes)
	var strBuilder strings.Builder

	getNext := func(index int, slice []rune) rune {
		value := '1'
		if index < len(slice)-1 {
			value = slice[index+1]
		}
		return value
	}

	for i := 0; i < count; i++ {
		cur := runes[i]
		next := getNext(i, runes)

		// fmt.Printf("1. i = %d; cur = %s; next = %s;\n", i, string(cur), string(next))

		if unicode.IsDigit(cur) {
			return "", ErrInvalidString
		}

		if cur == '\\' {
			if !(next == '\\' || unicode.IsDigit(next)) || (next == '1' && i == count-1) {
				return "", ErrInvalidString
			}
			i++

			cur = next
			next = getNext(i, runes)
		}

		if unicode.IsDigit(next) {
			i++
		} else {
			next = '1'
		}

		// fmt.Printf("2. i = %d; cur = %s; next = %s;\n", i, string(cur), string(next))

		repeat, _ := strconv.Atoi(string(next))
		strBuilder.WriteString(strings.Repeat(string(cur), repeat))
	}

	return strBuilder.String(), nil
}
