package main

import (
	"fmt"
)

func findMostRepeated(data []string) (string, int) {

	elementCounts := make(map[string]int)
	mostRepeatedElement := ""
	maxCount := 0

	for _, element := range data {
		elementCounts[element]++
		count := elementCounts[element]
		if count > maxCount {
			mostRepeatedElement = element
			maxCount = count
		}
	}

	return mostRepeatedElement, maxCount
}

func main() {
	datasets := [][]string{
		{"apple", "pie", "apple", "red", "red", "red"},
		{"a", "b", "c", "a", "a", "d", "d", "d"},
		{"1", "2", "3", "1", "1", "4", "5", "5", "5"},
		{"cat", "dog", "bird", "cat", "fish", "cat", "dog"},
	}

	for _, dataset := range datasets {
		mostRepeated, count := findMostRepeated(dataset)
		fmt.Printf("Dataset: %v\nMost repeated element: %s (repeated %d times)\n---\n", dataset, mostRepeated, count)
	}
}
