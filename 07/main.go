package main

import (
	"cmp"
	"log"
	"strconv"
	"strings"

	"slices"

	"github.com/ravenbluedragon/aoc-2023/common"
)

// boilerplate to load and solve puzzles
func main() {
	common.ChoosePart(7, solve1, solve2)
}

func solve1(filename string) any {
	lines := common.LoadData(filename)
	hands := parseHands(lines, false)
	slices.SortFunc(hands, compareHands)
	expectedValue := 0
	for i, hand := range hands {
		expectedValue += hand.bid * (i + 1)
	}
	return expectedValue
}

func solve2(filename string) any {
	lines := common.LoadData(filename)
	hands := parseHands(lines, true)
	slices.SortFunc(hands, compareHands)
	expectedValue := 0
	for i, hand := range hands {
		expectedValue += hand.bid * (i + 1)
	}
	return expectedValue
}

// HandType represents the type of a hand
type HandType uint8

const (
	HighCard HandType = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

// String converts a hand type to a string (for debugging)
func (h HandType) String() string {
	switch h {
	case HighCard:
		return "High Card"
	case OnePair:
		return "One Pair"
	case TwoPairs:
		return "Two Pairs"
	case ThreeOfAKind:
		return "Three of a Kind"
	case FullHouse:
		return "Full House"
	case FourOfAKind:
		return "Four of a Kind"
	case FiveOfAKind:
		return "Five of a Kind"
	}

	return "Unknown"
}

// cardStrength maps a card to its strength
// NOTE: Jokers are mapped to 1 for part 2
var cardStrength map[rune]uint8 = map[rune]uint8{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

// Hand represents a hand of cards
type Hand struct {
	cards     string
	strengths [5]uint8
	handType  HandType
	bid       int
}

// NewHand creates a new hand from a string
// NOTE: Jokers are mapped to 1 for part 2
func NewHand(line string, jokers bool) Hand {
	var hand Hand
	var err error
	parts := strings.Split(line, " ")
	hand.cards = parts[0]
	hand.bid, err = strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("Failed to parse bid %s: %v", parts[1], err)
	}
	for i, card := range hand.cards {
		if jokers && card == 'J' {
			hand.strengths[i] = 1
		} else {
			hand.strengths[i] = cardStrength[card]
		}
	}
	hand.handType = handType(parts[0], jokers)
	return hand
}

// handType determines the kind of a hand
func handType(hand string, jokers bool) HandType {
	counts := make(map[rune]int)
	for _, card := range hand {
		counts[card]++
	}

	wild := 0
	if j, ok := counts['J']; ok && jokers {
		delete(counts, 'J')
		wild = j
	}
	var values []int
	for _, count := range counts {
		values = append(values, count)
	}
	if len(values) < 2 {
		return FiveOfAKind
	}
	slices.Sort(values)
	slices.Reverse(values)

	return determineHandKind(values[0]+wild, values[1])
}

// determineHandKind determines the kind of a hand from the counts of the cards
// NOTE: The counts are assumed to be sorted in descending order
// This function is only called from handType, which sorts the counts
func determineHandKind(c1, c2 int) HandType {
	switch {
	case c1 == 5:
		return FiveOfAKind
	case c1 == 4:
		return FourOfAKind
	case c1 == 3 && c2 == 2:
		return FullHouse
	case c1 == 3:
		return ThreeOfAKind
	case c1 == 2 && c2 == 2:
		return TwoPairs
	case c1 == 2:
		return OnePair
	default:
		return HighCard
	}
}

// parseHands parses a slice of strings into a slice of hands
func parseHands(lines []string, jokers bool) []Hand {
	var hands []Hand
	for _, line := range lines {
		hands = append(hands, NewHand(line, jokers))
	}
	return hands
}

// compareHands compares two hands based on their kind and strengths
func compareHands(h1, h2 Hand) int {
	if h1.handType != h2.handType {
		return cmp.Compare(h1.handType, h2.handType)
	}
	return slices.Compare(h1.strengths[:], h2.strengths[:])
}
