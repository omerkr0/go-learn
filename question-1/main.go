package main

import (
	"fmt"
	"sort"
	"strings"
)

// sortByCharacterCount is a function that sorts a list of strings based on the count
// of the character 'a' in each string. If the counts are equal, it sorts them based
// on the length of the strings.
func sortByCharacterCount(words []string) []string {
	// sort.Slice function is used for sorting a slice.
	sort.Slice(words, func(x, y int) bool {
		// countA is a helper function that calculates the number of occurrences of 'a'
		// in a given word.
		countA := func(word string) int {
			return strings.Count(word, "a")
		}
		// Sorting criteria:
		// 1. Sort in descending order based on the count of 'a' in the words,
		// 2. If the counts are equal, sort in descending order based on the length of the words.
		return countA(words[x]) > countA(words[y]) || (countA(words[x]) == countA(words[y]) && len(words[x]) > len(words[y]))
	})
	return words
}

func main() {
	// inputWords contains the list of strings to be sorted using the sortByCharacterCount function.
	inputWords := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	// Sort the string slice using the sortByCharacterCount function.
	sortedWords := sortByCharacterCount(inputWords)
	// Print the sorted string slice.
	fmt.Println(sortedWords)
}
