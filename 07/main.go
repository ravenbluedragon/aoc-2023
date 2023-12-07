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
	// for i, hand := range hands {
	// 	log.Printf("% 5d -> %s [%02d %02d %02d %02d %02d] %s (%d)",
	// 		i,
	// 		hand.cards,
	// 		hand.strengths[0],
	// 		hand.strengths[1],
	// 		hand.strengths[2],
	// 		hand.strengths[3],
	// 		hand.strengths[4],
	// 		kindName(hand.kind),
	// 		hand.bid,
	// 	)
	// }
	expectedValue := 0
	for i, hand := range hands {
		expectedValue += hand.bid * (i + 1)
	}
	return expectedValue
}

const (
	HighCard uint8 = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func kindName(kind uint8) string {
	switch kind {
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
	default:
		return "Unknown"
	}
}

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

type Hand struct {
	cards     string
	strengths [5]uint8
	kind      uint8
	bid       int
}

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
	hand.kind = handKind(parts[0], jokers)
	return hand
}

func handKind(hand string, jokers bool) uint8 {
	counts := make(map[rune]int)
	for _, card := range hand {
		counts[card]++
	}

	wild := 0
	if j, ok := counts['J']; ok && jokers {
		delete(counts, 'J')
		wild = j
	}
	if wild == 5 {
		return FiveOfAKind
	}
	var values []int
	for _, count := range counts {
		values = append(values, count)
	}
	slices.Sort(values)
	slices.Reverse(values)
	values[0] += wild

	switch {
	case values[0] == 5:
		return FiveOfAKind
	case values[0] == 4:
		return FourOfAKind
	case values[0] == 3 && values[1] == 2:
		return FullHouse
	case values[0] == 3:
		return ThreeOfAKind
	case values[0] == 2 && values[1] == 2:
		return TwoPairs
	case values[0] == 2:
		return OnePair
	default:
		return HighCard
	}
}

func parseHands(lines []string, jokers bool) []Hand {
	var hands []Hand
	for _, line := range lines {
		hands = append(hands, NewHand(line, jokers))
	}
	return hands
}

func compareHands(h1, h2 Hand) int {
	if h1.kind != h2.kind {
		return cmp.Compare(h1.kind, h2.kind)
	}
	return slices.Compare(h1.strengths[:], h2.strengths[:])
}
