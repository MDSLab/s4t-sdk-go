package utils

import (
	"math/rand"
	"time"
)

func randStringGenerator(n int) string {
    
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	
	for i := range b {
        b[i] = letters[seededRand.Intn(len(letters))]
    }
    return string(b)
}

func CompareFields(data map[string] interface{}, key_array []string) bool {
	for key := range data {
		found := false
		for _,board_key := range key_array{
			if key == board_key {
				found = true
				break
			}
		}
		if !found {
			return false
		}

	}
	return true
}
