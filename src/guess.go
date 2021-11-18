package src

import (
	"errors"
	"math"
)

// SimilarityScore compares a text's distribution to that of a language
// and returns a similarity score for that language.
// TODO: look into better similarity measures
func simliarityScore(dist map[string]float64, lang map[string]float64) (float64, error) {
	var scoreSum float64
	for k := range dist {
		if _, exists := lang[k]; !exists {
			// discount language as possible if text contains a letter that isn't in the language
			return 0, errors.New("A letter in the distribution does not exist in the language.")
		}
		letterScore := math.Abs(dist[k] - lang[k])
		scoreSum += letterScore
	}
	score := scoreSum / float64(len(dist))

	return score, nil
}
