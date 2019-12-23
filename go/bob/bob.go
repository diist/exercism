WIP - wanna look at community solutions

// Package bob contains the Exercism logic for Bob
package bob

import "regexp"

// Hey replies like an annoying teenager
func Hey(remark string) string {
	var response string

	statement := regexp.MustCompile(`[A-Z][a-z ,-]+\.`)
	shouting := regexp.MustCompile(`[A-Z]{2,}!?`)
	asking := regexp.MustCompile(`[a-z0-9 ,-]+\?`)
	forcing := regexp.MustCompile(`[a-z0-9 ,-]+\!`)

	switch {
	case statement.MatchString(remark):
		response = "Whatever."
	case shouting.MatchString(remark):
		response = "Whoa, chill out!"
	case asking.MatchString(remark):
		response = "Sure."
	case forcing.MatchString(remark):
		response = "Whatever."
	default:
		response = "Shush."
	}

	return response
}
