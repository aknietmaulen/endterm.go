package main

import "regexp"

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

func isValidSurname (surname string) bool {
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

