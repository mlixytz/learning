package shuffle

import (
	"math/rand"
)

/*
	The Knuth shuffle (a.k.a. the Fisher-Yates shuffle) is an algorithm
	for randomly shuffling the elements of an array.
*/

func shuffle(input []int) []int {
	for i := len(input); i > 0; i-- {
		j := rand.Intn(i)
		input[i-1], input[j] = input[j], input[i-1]
	}

	return input
}
