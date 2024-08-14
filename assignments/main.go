package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	fmt.Println("hello world")

	fmt.Println(arraySign([]int{2, 1}))                    // 1
	fmt.Println(arraySign([]int{-2, 1}))                   // -1
	fmt.Println(arraySign([]int{-1, -2, -3, -4, 3, 2, 1})) // 1

	fmt.Println("anak kana", isAnagram("anak", "kana"))             // true
	fmt.Println("anak mana", isAnagram("anak", "mana"))             // false
	fmt.Println("anagram managra", isAnagram("anagram", "managra")) // true

	fmt.Println(string(findTheDifference("abcd", "abcde"))) // 'e'
	fmt.Println(string(findTheDifference("abcd", "abced"))) // 'e'
	fmt.Println(string(findTheDifference("", "y")))         // 'y'

	fmt.Println(canMakeArithmeticProgression([]int{1, 5, 3}))    // true; 1, 3, 5 adalah baris aritmatik +2
	fmt.Println(canMakeArithmeticProgression([]int{5, 1, 9}))    // true; 9, 5, 1 adalah baris aritmatik -4
	fmt.Println(canMakeArithmeticProgression([]int{1, 2, 4, 8})) // false; 1, 2, 4, 8 bukan baris aritmatik, melainkan geometrik x2

	tesDeck()
}

// https://leetcode.com/problems/sign-of-the-product-of-an-array
func arraySign(nums []int) int {
	// write code here
	signs := 1

	for _, v := range nums {
		if v == 0 {
			return 0
		}

		if v < 0 {
			signs = signs * -1
		}
	}

	return signs // if positive
	// return -1 // if negative
}

// https://leetcode.com/problems/valid-anagram
func isAnagram(s string, t string) bool {
	runeS := []rune(s)
	runeT := []rune(t)

	if len(s) != len(t) {
		return false
	}

	sort.Slice(runeS, func(i, j int) bool {
		return runeS[i] < runeS[j]
	})
	sort.Slice(runeT, func(i, j int) bool {
		return runeT[i] < runeT[j]
	})

	if string(runeS) != string(runeT) {
		return false
	}

	return true
}

// lenovo [23,50,22,12, 76,90]

// https://leetcode.com/problems/find-the-difference
func findTheDifference(s string, t string) byte {
	runesS := []rune(s)
	runesT := []rune(t)

	// b := byte(0)

	elements := make(map[rune]int)

	for _, v := range runesS {
		elements[v]++
	}

	for _, v := range runesT {
		elements[v]++
	}

	for i, v := range elements {
		if v == 1 {
			return byte(i)
		}
	}

	// write code here

	return byte(0)
}

// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence
func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) < 2 {
		return false // Minimal 2 elemen diperlukan untuk sebuah deret aritmetika
	}

	// Urutkan array untuk mempermudah pemeriksaan
	sort.Ints(arr)
	var arr2 []int
	// write code here
	difference := arr[0] - arr[1]
	for i := 0; i < len(arr)-1; i++ {
		arr2 = append(arr2, arr[i]-arr[i+1])
	}

	for i := 0; i < len(arr2); i++ {
		if arr2[i] != difference {
			return false
		}
	}
	return true
}

// Deck represents a "standard" deck consisting of 52 cards
type Deck struct {
	cards []Card
}

// Card represents a card in a "standard" deck
type Card struct {
	symbol int // 0: spade, 1: heart, 2: club, 3: diamond
	number int // Ace: 1, Jack: 11, Queen: 12, King: 13
}

// New inserts 52 cards into deck d, sorted by symbol & then number.
func (d *Deck) New() {
	d.cards = make([]Card, 0, 52)
	for symbol := 0; symbol < 4; symbol++ {
		for number := 1; number <= 13; number++ {
			d.cards = append(d.cards, Card{symbol: symbol, number: number})
		}
	}
}

// PeekTop returns n cards from the top
func (d Deck) PeekTop(n int) []Card {
	if n > len(d.cards) {
		n = len(d.cards)
	}
	return d.cards[:n]
}

// PeekBottom returns n cards from the bottom
func (d Deck) PeekBottom(n int) []Card {
	if n > len(d.cards) {
		n = len(d.cards)
	}
	return d.cards[len(d.cards)-n:]
}

// PeekCardAtIndex returns a card at the specified index
func (d Deck) PeekCardAtIndex(idx int) Card {
	if idx < 0 || idx >= len(d.cards) {
		return Card{} // Return a default Card if the index is out of range
	}
	return d.cards[idx]
}

// Shuffle randomly shuffles the deck
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

// Cut performs a single "Cut" technique. Moves n top cards to bottom.
func (d *Deck) Cut(n int) {
	if n > len(d.cards) || n <= 0 {
		return
	}
	d.cards = append(d.cards[n:], d.cards[:n]...)
}

func (c Card) ToString() string {
	textNum := ""
	switch c.number {
	case 1:
		textNum = "Ace"
	case 11:
		textNum = "Jack"
	case 12:
		textNum = "Queen"
	case 13:
		textNum = "King"
	default:
		textNum = fmt.Sprintf("%d", c.number)
	}
	texts := []string{"Spade", "Heart", "Club", "Diamond"}
	return fmt.Sprintf("%s of %s", textNum, texts[c.symbol])
}

func tesDeck() {
	deck := Deck{}
	deck.New()

	fmt.Println("Top 3 Cards:")
	top3Cards := deck.PeekTop(3)
	for _, c := range top3Cards {
		fmt.Println(c.ToString())
	}
	fmt.Println("---\n")

	fmt.Println("Specific Cards by Index:")
	fmt.Println(deck.PeekCardAtIndex(12).ToString()) // Queen Spade
	fmt.Println(deck.PeekCardAtIndex(13).ToString()) // King Spade
	fmt.Println(deck.PeekCardAtIndex(14).ToString()) // Ace Heart
	fmt.Println(deck.PeekCardAtIndex(15).ToString()) // 2 Heart
	fmt.Println("---\n")

	deck.Shuffle()
	fmt.Println("Top 10 Cards after Shuffle:")
	top10Cards := deck.PeekTop(10)
	for _, c := range top10Cards {
		fmt.Println(c.ToString())
	}

	fmt.Println("---\n")
	deck.New()
	deck.Cut(5)
	fmt.Println("Bottom 10 Cards after Cut:")
	bottom10Cards := deck.PeekBottom(10)
	for _, c := range bottom10Cards {
		fmt.Println(c.ToString())
	}
}
