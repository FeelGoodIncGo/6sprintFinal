package service

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Функция для работы с азбукой Морзе
func Text(tM string) (string, error) {

	// Удаляем пробелы
	input := strings.TrimSpace(tM)

	// Проверка на пустую строку
	if input == "" {
		return "", errors.New("empty string")
	}

	// Проверка символов на соответствие азбуке Морзе
	isMorse := func(c rune) bool {
		return c != ' ' && c != '-' && c != '.'
	}
	// Поиск символа, который не является частью азбуки Морзе
	index := strings.IndexFunc(input, isMorse)

	// Проходим по найденным символам, преобразуем в большие буквы, и проверяем на наличие в азбуке Морзе
	if index >= 0 {
		for _, v := range input {
			if v == ' ' {
				continue
			}
			upperR := unicode.ToUpper(v)
			if _, ok := morse.DefaultMorse[upperR]; !ok {
				return "", fmt.Errorf("unknown symbol: %q", v)
			}
		}
		// Преобразуем в азбуку Морзе
		return morse.ToMorse(input), nil
	}
	// Преобразуем в текст
	return morse.ToText(input), nil
}
