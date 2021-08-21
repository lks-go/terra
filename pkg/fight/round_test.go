package fight_test

import (
	"testing"

	"github.com/lks-go/terra/pkg/fight"
)

func TestRound_LessThanTwoFighters(t *testing.T) {
	if _, err := fight.NewRound([]int{0}); err == nil {
		t.Error("expected an error because at least two fighters must participate in the round")
	}
}

func TestRound_OneOnOne(t *testing.T) {
	r, err := fight.NewRound([]int{0, 1})
	if err != nil {
		t.Errorf("expected no errors, got: %s", err)
	}

	r.AddAction(0, 1, []int{0, 1}, []int{0, 1})
	if r.Finished() {
		t.Error("expected not finished round, got finished")
	}

	r.AddAction(1, 0, []int{0, 1}, []int{0, 1})
	if !r.Finished() {
		t.Error("expected finished round, got not finished")
	}
}

func TestRound_FiveFighters(t *testing.T) {
	totalFightersNumber := 5
	maxFighterNumber := totalFightersNumber - 1

	fightersList := make([]int, 0, totalFightersNumber)
	for i := 0; i < totalFightersNumber; i++ {
		fightersList = append(fightersList, i)
	}

	r, err := fight.NewRound(fightersList)
	if err != nil {
		t.Errorf("expected no errors, got: %s", err)
	}

	for _, fn := range fightersList {
		for en := 0; en < totalFightersNumber; en++ {
			if fn == en {
				continue
			}

			r.AddAction(fn, en, []int{0, 1}, []int{0, 1})
			if fn < maxFighterNumber && en <= maxFighterNumber && r.Finished() {
				t.Error("expected not finished round, got finished")
			}
		}
	}

	if !r.Finished() {
		t.Error("expected finished round, got not finished")
	}
}

func TestRound_ActionsList(t *testing.T) {

}
