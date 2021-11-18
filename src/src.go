package src

import "strings"

// Contains functions used by both guess and learn modes.

// Distribution splits a string into individual letters and calculates a distribution of prevalence
func Distribution(str string) map[string]float64 {
	letters := strings.Split(str, "")
	total := len(letters)
	unique := uniqueLetters(str)

	distribution := make(map[string]float64)

	// get prevalence percentage of each unique letter
	// and create a map of distribution
	for _, letter := range unique {
		count := strings.Count(str, letter)
		percentage := float64(count) / float64(total)

		distribution[letter] = percentage
	}
	return distribution
}

// uniqueLetters returns a slice of all unique letters in a string.
// TODO: make this remove punctuation
func uniqueLetters(str string) []string {
	letters := strings.Split(str, "")
	unique := make([]string, 0, 30)
	for _, letter := range letters {
		found := findInSlice(unique, letter)
		if !found {
			unique = append(unique, letter)
		}
	}
	return unique
}

// findInSlice takes a slice and a string and returns true
// if the string is in the slice.
func findInSlice(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
