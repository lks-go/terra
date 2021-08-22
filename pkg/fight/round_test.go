package fight_test

import (
	"fmt"
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

func TestFight_Actions(t *testing.T) {
	// TODO must have
}

func TestFight_ShowEnemyTwoFighters(t *testing.T) {
	const (
		minusOne = -1
		zero     = 0
		one      = 1
	)

	r, err := fight.NewRound([]int{zero, one})
	if err != nil {
		t.Errorf("expected no errors, got: %s", err)
	}

	en := r.ShowEnemyNumber(zero)
	if en != one {
		t.Errorf("expected enemy number %d, got: %d", one, en)
	}

	en = r.ShowEnemyNumber(one)
	if en != zero {
		t.Errorf("expected enemy number %d, got: %d", zero, en)
	}

	r.AddAction(zero, one, []int{1}, []int{2})

	en = r.ShowEnemyNumber(zero)
	if en != minusOne {
		t.Errorf("expected enemy number %d, got: %d", minusOne, en)
	}

	r.AddAction(one, zero, []int{1}, []int{2})

	en = r.ShowEnemyNumber(one)
	if en != minusOne {
		t.Errorf("expected enemy number %d, got: %d", minusOne, en)
	}

	if !r.Finished() {
		t.Errorf("expected status finished")
	}
}

func TestFight_ShowEnemyThreeFighters(t *testing.T) {

	const (
		minusOne   = -1
		five       = 5
		twelve     = 12
		ninetyNine = 99
	)

	blocks := []int{1, 3}
	attacks := []int{1}

	r, err := fight.NewRound([]int{five, twelve, ninetyNine})
	if err != nil {
		t.Errorf("expected no errors, got: %s", err)
	}

	tests := []struct {
		fighter       int
		expectedEnemy int
	}{
		{five, twelve},
		{ninetyNine, five},
		{twelve, five},
		{five, ninetyNine},
		{five, minusOne},
		{twelve, ninetyNine},
		{ninetyNine, twelve},
		{ninetyNine, minusOne},
		{twelve, minusOne},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d vs %d", tt.fighter, tt.expectedEnemy)
		t.Run(testName, func(t *testing.T) {
			en := r.ShowEnemyNumber(tt.fighter)
			if en != tt.expectedEnemy {
				t.Errorf("expected enemy number %d, got: %d", tt.expectedEnemy, en)
			}

			if tt.expectedEnemy == minusOne {
				return
			}

			r.AddAction(tt.fighter, tt.expectedEnemy, blocks, attacks)
		})
	}

	if !r.Finished() {
		t.Errorf("expected status finished")
	}
}
