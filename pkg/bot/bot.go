package bot

import (
	"math/rand"
	"time"
)

func RandActions(max, need int) []int {
	if need >= max {
		res := make([]int, need)

		for i := 0; i < need; i++ {
			res[i] = i
		}

		return res
	}

	res := make([]int, 0)

	if need <= 0 {
		return res
	}

	rand.Seed(time.Now().UnixNano())

	for {
		if len(res) == need {
			break
		}

		newNum := rand.Intn(max)
		if !checkIfNumberExist(res, newNum) {
			res = append(res, newNum)
		}
	}

	return res
}

func checkIfNumberExist(list []int, n int) bool {
	if len(list) == 0 {
		return false
	}

	for _, ln := range list {
		if ln == n {
			return true
		}
	}

	return false
}

func hasDuplicates(list []int) bool {
	for idx1, n1 := range list {
		for idx2, n2 := range list {
			if idx1 == idx2 {
				continue
			}

			if n1 == n2 {
				return true
			}
		}
	}

	return false
}
