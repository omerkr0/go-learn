package main

import (
	"fmt"
)

// findMostRepeated takes a slice of strings 'data' and returns the most repeated element
// along with the count of its occurrences in the slice.
func findMostRepeated(data []string) (string, int) {
	// elementCounts is a map to store the count of each element in the data slice.
	elementCounts := make(map[string]int)
	// mostRepeatedElement keeps track of the most repeated element.
	mostRepeatedElement := ""
	// maxCount keeps track of the maximum count of the most repeated element.
	maxCount := 0

	// Iterate over each element in the data slice.
	for _, element := range data {
		// Increment the count of the current element in the map.
		elementCounts[element]++
		// Get the current count of the element.
		count := elementCounts[element]
		// Update mostRepeatedElement and maxCount if the current element has a higher count.
		if count > maxCount {
			mostRepeatedElement = element
			maxCount = count
		}
	}

	// Return the most repeated element and its count.
	return mostRepeatedElement, maxCount
}

func main() {
	// datasets contains multiple slices of strings to test the findMostRepeated function.
	datasets := [][]string{
		{"apple", "pie", "apple", "red", "red", "red"},
		{"a", "b", "c", "a", "a", "d", "d", "d"},
		{"1", "2", "3", "1", "1", "4", "5", "5", "5"},
		{"cat", "dog", "bird", "cat", "fish", "cat", "dog"},
	}

	// Iterate over each dataset and apply the findMostRepeated function.
	for _, dataset := range datasets {
		// Retrieve the most repeated element and its count from the findMostRepeated function.
		mostRepeated, count := findMostRepeated(dataset)
		// Print the results for each dataset.
		fmt.Printf("Dataset: %v\nMost repeated element: %s (repeated %d times)\n---\n", dataset, mostRepeated, count)
	}
}
