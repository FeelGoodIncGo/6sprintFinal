package service

import (
	"fmt"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

type Converter struct {
	runeToMorse       map[rune]string
	morseToRune       map[string]rune
	charSeparator     string
	wordSeparator     string
	convertToUpper    bool
	trailingSeparator bool

	Handling ErrorHandler
}
type ErrorHandler func(error) string

// AutoConvert автоматически определяет и конвертирует текст в Морзе или Морзе в текст
func (c *Converter) AutoConvert(input string) (string, error) {
	// Определяем текст или код Морзе
	if c.isText(input) {
		// Конвертируем текст в Морзе
		return morse.ToMorse(input), nil
	} else if c.isMorse(input) {
		// Конвертируем Морзе в текст
		return morse.ToText(input), nil
	}
	return "", fmt.Errorf("Error format")
}

// isText проверка на строку текста
func (c *Converter) isText(input string) bool {
	for _, char := range input {
		if !unicode.IsLetter(char) && !unicode.IsSpace(char) {
			return false
		}
	}
	return true
}

// isMorse провека на код Морзе
func (c *Converter) isMorse(input string) bool {
	morseChars := map[rune]bool{
		'.': true, '-': true, ' ': true,
	}
	for _, char := range input {
		if !morseChars[char] {
			return false
		}
	}
	return true
}
