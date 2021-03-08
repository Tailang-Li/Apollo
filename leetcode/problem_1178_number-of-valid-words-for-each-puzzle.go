package main

import "fmt"

/*
With respect to a given puzzle string, a word is valid if both the following conditions are satisfied:
word contains the first letter of puzzle.
For each letter in word, that letter is in puzzle.
For example, if the puzzle is "abcdefg", then valid words are "faced", "cabbage", and "baggage"; while invalid words are "beefed" (doesn't include "a") and "based" (includes "s" which isn't in the puzzle).
Return an array answer, where answer[i] is the number of words in the given word list words that are valid with respect to the puzzle puzzles[i].
*/

func main() {
	fmt.Println(findNumOfValidWords([]string{"aaaa", "asas", "able", "ability", "actt", "actor", "access"}, []string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"}))
	fmt.Println(findNumOfValidWords([]string{"apple", "pleas", "please"}, []string{"aelwxyz", "aelpxyz", "aelpsxy", "saelpxy", "xaelpsy"}))
}

type (
	wordDetail struct {
		beginLetter byte
		letterSet   map[byte]bool
	}
)

func findNumOfValidWords(words []string, puzzles []string) []int {

	wordDetails := make([]wordDetail, len(words))
	for i, word := range words {

		wordDetails[i] = wordDetail{
			letterSet: make(map[byte]bool),
		}

		for j := range word {
			wordDetails[i].letterSet[word[j]] = true
		}
	}

	puzzleDetails := make([]wordDetail, len(puzzles))
	for i, puzzle := range puzzles {

		puzzleDetails[i] = wordDetail{
			beginLetter: puzzle[0],
			letterSet:   make(map[byte]bool),
		}

		for j := range puzzle {
			puzzleDetails[i].letterSet[puzzle[j]] = true
		}
	}

	result := make([]int, len(puzzles))
	for i, puzzleDetail := range puzzleDetails {

	LOOP:
		for _, wordDetail := range wordDetails {

			if !wordDetail.letterSet[puzzleDetail.beginLetter] {
				continue LOOP
			}

			for letter := range wordDetail.letterSet {
				if !puzzleDetail.letterSet[letter] {
					continue LOOP
				}
			}

			result[i]++
		}
	}

	return result
}
