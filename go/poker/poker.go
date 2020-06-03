WIP - just want to have a look at community solutions

// Package poker contains the logic to solve Exercism Go Poker exercise
package poker

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Card is a playing card
type Card struct {
	value int
	suit  string // embedded struct as enum?
}

// Hand is a set of 5 playing cards
type Hand struct {
	cards [5]Card
}

// BestHand returns the best poker hand out of a set of hands
func BestHand(hands []string) ([]string, error) {
	var bh []string
	var handss []Hand

	for _, h := range hands {
		hh, err := parseHand(h)
		if err != nil {
			panic(err)
		}
		handss := append(handss, hh)
	}

	if len(hands) == 1 {
		return hands, nil
	}

	fmt.Println(handss)

	return bh, nil
}

func parseHand(h string) (Hand, error) {
	var hand Hand
	cards := strings.Split(h, " ")

	for _, c := range cards {
		card, err := parseCard(c)
		if err != nil {
			errors.New("invalid card: " + c)
		}
	}
}

func parseCard(c string) (Card, error) {
	return Card{3, "♤"}, nil
}

func checkCard(h string) bool {
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
