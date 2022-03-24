package main

import (
	"regexp"
)

//ValidaFormatoEmail ::
func ValidaFormatoEmail(email string) (ok bool) {
	emailRegexp := regexp.MustCompile(`^[a-zA-Z0-9\.%+-_]+@[a-zA-Z0-9-_]*\.[a-zA-Z]{3}[(a-zA-z\.)?]*$`)
	if !emailRegexp.MatchString(email) {
		return
	}
	return true
}
