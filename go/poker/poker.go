// Package poker contains the logic to solve Exercism Go Poker exercise
package poker

import (
	"errors"
	"regexp"
	"strings"
)

// BestHand returns the best poker hand out of a set of hands
func BestHand(hands []string) ([]string, error) {

	var bh []string

	for _, h := range hands {
		if !checkHand(h) {
			return bh, errors.New("invalid hand: " + h)
		}
	}

	if len(hands) == 1 {
		return hands, nil
	}

	return bh, nil
}

func checkHand(h string) bool {
	cardRegex := regexp.MustCompile(`^([2-9]|10|J|Q|K|A)(♤|♧|♢|♡)$`)
	cards := strings.Split(h, " ")

	if len(cards) != 5 {
		return false
	}

	for _, c := range cards {
		if !cardRegex.MatchString(c) {
			return false
		}
	}

	return true
}
