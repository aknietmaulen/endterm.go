package main

import (
	"regexp"
	"time"
)

func isUpper(s string) bool {
	if len(s) == 0 {
		return false
	}
	return 'A' <= s[0] && s[0] <= 'Z'
}
func isAlphabetic(char rune) bool {
	return ('A' <= char && char <= 'Z') || ('a' <= char && char <= 'z')
}

func isValidName(name string) bool {
	if len(name) < 3 || !isUpper(name) {
		return false
	}

	for _, char := range name {
		if !isAlphabetic(char) {
			return false
		}
	}

	return true
}

func isValidSurname(surname string) bool {
	if len(surname) < 3 || !isUpper(surname) {
		return false
	}

	for _, char := range surname {
		if !isAlphabetic(char) {
			return false
		}
	}

	return true
}

func isValidEmail(email string) bool {
	emailPattern := `^[a-zA-Z0-9!#$%&'*+\-/=?^_{}|~]+(\.[a-zA-Z0-9!#$%&'*+\-/=?^_{}|~]+)*@([a-zA-Z0-9](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`
	return regexp.MustCompile(emailPattern).MatchString(email)
}

func isValidDate(date string) bool {
    t, err := time.Parse("02.01.2006", date)
    if err != nil {
        return false
    }
    currentDate := time.Now()
    return t.After(currentDate) || t.Equal(currentDate)
}

func isValidTime(timeStr string) bool {
	_, err := time.Parse("15:04", timeStr)
	return err == nil
}
