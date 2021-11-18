package src

import (
	"strings"
	"unicode"
)

// Contains functions used by both guess and learn modes.

// Distribution splits a string into individual letters and calculates a distribution of prevalence
func Distribution(str string) map[string]float64 {

	var letters []string
	characters := strings.Split(str, "")
	for _, char := range characters {
		if unicode.IsLetter([]rune(char)[0]) {
			char = strings.ToLower(char)
			letters = append(letters, char)
		}
	}
	total := len(letters)
	unique := unique(letters)

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

// unique returns a slice of all unique characters in a slice.
func unique(slice []string) []string {
	unique := make([]string, 0, 30)
	for _, letter := range slice {
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
