package main

import (
	"slices"
	"testing"

	"github.com/ravenbluedragon/aoc-2023/common"
)

const filename = "../test-data/07.txt"

func TestSolve1(t *testing.T) {
	expected := 6440
	actual := solve1(filename)

	if actual != expected {
		t.Errorf("evaluate1(%s) = %d, expected %d", filename, actual, expected)
	}
}

func TestSolve2(t *testing.T) {
	expected := 5905
	actual := solve2(filename)

	if actual != expected {
		t.Errorf("evaluate2(%s) = %d, expected %d", filename, actual, expected)
	}
}

func TestHandTypesSimple(t *testing.T) {
	lines := common.LoadData(filename)
	hands := parseHands(lines, false)
	slices.SortFunc(hands, compareHands)
	type test struct {
		hand string
		kind uint8
		rank int
	}
	expected := []test{
		{"32T3K", OnePair, 1},
		{"KTJJT", TwoPairs, 2},
		{"KK677", TwoPairs, 3},
		{"T55J5", ThreeOfAKind, 4},
		{"QQQJA", ThreeOfAKind, 5},
	}
	for i, hand := range hands {
		if hand.kind != expected[i].kind {
			t.Errorf("hand %s has kind %s, expected %s", hand.cards, kindName(hand.kind), kindName(expected[i].kind))
		}
		if i+1 != expected[i].rank {
			t.Errorf("hand %s has rank %d, expected %d", hand.cards, i, expected[i].rank)
		}
	}
}

func TestHandTypesJokers(t *testing.T) {
	lines := common.LoadData(filename)
	hands := parseHands(lines, true)
	slices.SortFunc(hands, compareHands)
	type test struct {
		hand string
		kind uint8
		rank int
	}
	expected := []test{
		{"32T3K", OnePair, 1},
		{"KK677", TwoPairs, 2},
		{"T55J5", FourOfAKind, 3},
		{"QQQJA", FourOfAKind, 4},
		{"KTJJT", FourOfAKind, 5},
	}
	for i, hand := range hands {
		if hand.kind != expected[i].kind {
			t.Errorf("hand %s has kind %s, expected %s", hand.cards, kindName(hand.kind), kindName(expected[i].kind))
		}
		if i+1 != expected[i].rank {
			t.Errorf("hand %s has rank %d, expected %d", hand.cards, i, expected[i].rank)
		}
	}
}

func TestHandKind(t *testing.T) {
	table := []struct {
		hand   string
		jokers bool
		kind   uint8
	}{
		{"32T3K", false, OnePair},
		{"KK677", false, TwoPairs},
		{"T55J5", false, ThreeOfAKind},
		{"QQQJA", false, ThreeOfAKind},
		{"KTJJT", false, TwoPairs},
		{"KKKKK", false, FiveOfAKind},
		{"KKKKK", true, FiveOfAKind},
		{"KKKJJ", false, FullHouse},
		{"KKKJJ", true, FiveOfAKind},
		{"JJJJJ", false, FiveOfAKind},
		{"JJJJJ", true, FiveOfAKind},
		{"J2783", false, HighCard},
		{"J2783", true, OnePair},
		{"J2345", true, OnePair},
		{"J2234", true, ThreeOfAKind},
		{"J2233", true, FullHouse},
		{"J2223", true, FourOfAKind},
		{"J2222", true, FiveOfAKind},
	}
	for _, test := range table {
		kind := handKind(test.hand, test.jokers)
		if kind != test.kind {
			t.Errorf("hand %s (jokers %v) has kind %s, expected %s", test.hand, test.jokers, kindName(kind), kindName(test.kind))
		}
	}

}
