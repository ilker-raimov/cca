package regex

import (
	"regexp"
)

var email_regex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
var password_length_regex = regexp.MustCompile(`^.{8,}$`)
var password_upper_regex = regexp.MustCompile(`[A-Z]`)
var password_digit_regex = regexp.MustCompile(`\d`)
var password_special_regex = regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`)
var password_regex_list = []*regexp.Regexp{password_length_regex, password_upper_regex, password_digit_regex, password_special_regex}

func Match(value string, regex *regexp.Regexp) bool {
	return regex.MatchString(value)
}

func MatchAll(value string, regex_list []*regexp.Regexp) bool {
	for _, regex := range regex_list {
		match := regex.MatchString(value)

		if !match {
			return false
		}
	}

	return true
}

func Email(email string) bool {
	return Match(email, email_regex)
}

func Password(password string) bool {
	return MatchAll(password, password_regex_list)
}
