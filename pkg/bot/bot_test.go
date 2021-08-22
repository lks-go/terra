package bot

import (
	"testing"
)

func Test_checkIfNumberExist(t *testing.T) {

	list := []int{1, 2, 9, 55}
	needle := 3

	if checkIfNumberExist(list, needle) {
		t.Error("expected response: false, got: true")
	}

	list = []int{1, 8, 2, 16, 9, 55, 333}
	needle = 16

	if !checkIfNumberExist(list, needle) {
		t.Error("expected response: true, got: false")
	}

	list = []int{}
	needle = 1024

	if checkIfNumberExist(list, needle) {
		t.Error("expected response: false, got: true")
	}

}

func Test_RandActions(t *testing.T) {

	bodyParts := []int{0, 1, 2, 4}

	tests := []struct {
		need   int
		expect int
	}{
		{1, 1},
		{0, 0},
		{-1, 0},
		{-500, 0},
		{100, 100},
		{2, 2},
	}

	for _, tt := range tests {
		res := RandActions(len(bodyParts), tt.need)
		if len(res) != tt.expect {
			t.Errorf("expected len: %d, got: %d", tt.expect, res)
		}

		if hasDuplicates(res) {
			t.Error("expected no duplicates")
		}
	}

}

func Test_hasDuplicates(t *testing.T) {

	if hasDuplicates([]int{1, 2, 3}) {
		t.Error("expected no duplicates")
	}

	if hasDuplicates([]int{22}) {
		t.Error("expected no duplicates")
	}

	if !hasDuplicates([]int{1, 1, 2, 3}) {
		t.Error("expected duplicates")
	}

	if !hasDuplicates([]int{99, 99}) {
		t.Error("expected duplicates")
	}
}
