package service

import (
	"errors"
	"strings"

	"github.com/JimmyR7/6-sprint/pkg/morse"
)

func Convert(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", errors.New("empty input")
	}

	if strings.Trim(input, ".- \t\r\n") == "" {
		for _, symbol := range strings.Fields(input) {
			if morse.MorseToRune(symbol) == 0 {
				return "", errors.New("unknown Morse symbol: " + symbol)
			}
		}
		return morse.ToText(input), nil
	}

	for _, symbol := range input {
		if symbol == '\n' || symbol == '\r' || symbol == '\t' || symbol == ' ' {
			continue
		}
		if morse.RuneToMorse(symbol) == "" {
			return "", errors.New("unsupported text character: " + string(symbol))
		}
	}
	return morse.ToMorse(input), nil
}
