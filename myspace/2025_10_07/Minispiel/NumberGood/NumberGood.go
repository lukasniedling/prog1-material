package NumberGood

import "math/rand"

func NumberGood(guess int) bool {
	n := rand.Intn(3) + 1
	if guess == n {
		return true
		} else {
			return false
		}
	}