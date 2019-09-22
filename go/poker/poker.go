// Package poker contains the logic to solve Exercism Go Poker exercise
package poker

import (
	"errors"
	"regexp"
)

// BestHand returns the best poker hand out of a set of hands
func BestHand(hands []string) ([]string, error) {

	var bh []string

	for _, h := range hands {
		if !checkHand(h) {
			return bh, errors.New("invalid hand: " + h)
		}
	}

	return bh, nil
}

func checkHand(h string) bool {
	vh := regexp.MustCompile(`^((([2-9]|10|J|Q|K|A)(♤|♧|♢|♡) )){4}(([2-9]|10|J|Q|K|A)(♤|♧|♢|♡))$`)

	return vh.MatchString(h)
}
