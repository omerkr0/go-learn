package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortByCharacterCount(words []string) []string {
	sort.Slice(words, func(x, y int) bool {
		countA := func(word string) int {
			return strings.Count(word, "a")
		}
		return countA(words[x]) > countA(words[y]) || (countA(words[x]) == countA(words[y]) && len(words[x]) > len(words[y]))
	})
	return words
}

func main() {
	inputWords := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	sortedWords := sortByCharacterCount(inputWords)
	fmt.Println(sortedWords)
}
