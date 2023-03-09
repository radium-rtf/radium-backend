package email

import "regexp"

var emailRegex = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)

func ValidateEmail(email string) bool {
	return emailRegex.MatchString(email)
}
