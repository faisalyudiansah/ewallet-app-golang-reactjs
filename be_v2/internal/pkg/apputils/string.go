package apputils

import (
	"net/url"
	"regexp"
)

func GetStringValueOrDefault(newValue, defaultValue string) string {
	if newValue != "" {
		return newValue
	}
	return defaultValue
}

func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func IsValidURL(input string) bool {
	_, err := url.ParseRequestURI(input)

	return err == nil
}
