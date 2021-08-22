package fight_test

import (
	"fmt"
	"testing"

	"github.com/lks-go/terra/pkg/bot"

	"github.com/lks-go/terra/pkg/fighter"

	"github.com/lks-go/terra/pkg/fight"
)

func newSimpleFighter() fighter.Fighter {
	head := fighter.NewPart(&fighter.PartConfig{})
	chest := fighter.NewPart(&fighter.PartConfig{})
	groin := fighter.NewPart(&fighter.PartConfig{})
	feet := fighter.NewPart(&fighter.PartConfig{})

	return fighter.New(&fighter.Config{}, []fighter.DamageGetter{head, chest, groin, feet}, nil)
}

func TestFight_Status(t *testing.T) {

	tests := []struct {
		StatusName string
		StatusCode int
		IsError    bool
	}{
		{StatusName: "Unknown", StatusCode: 0, IsError: true},
		{StatusName: "Created", StatusCode: 1, IsError: false},
		{StatusName: "Going", StatusCode: 2, IsError: false},
		{StatusName: "Finished", StatusCode: 3, IsError: false},
		{StatusName: "One more unknown status", StatusCode: 2021, IsError: true},
	}

	f := fight.New(&fight.Config{})

	for _, tt := range tests {
		testName := fmt.Sprintf("Status name: %s, satus code %d", tt.StatusName, tt.StatusCode)
		t.Run(testName, func(t *testing.T) {
			if err := f.SetStatus(fight.Status(tt.StatusCode)); err == nil && tt.IsError {
				t.Errorf("expected error, got nil")
			}

			// test must be failed before than escape the function
			if tt.IsError {
				return
			}

			gotStatus := f.Status()
			if gotStatus != fight.Status(tt.StatusCode) {
				t.Errorf("expected status code %d, got %d", tt.StatusCode, gotStatus)
			}
		})
	}

}

func TestFight_CanJoin(t *testing.T) {
	testFight := fight.New(&fight.Config{FightersLimit: 2})

	if !testFight.CanJoin() {
		t.Errorf("excpected availble acces to join the fight")
	}

	f1 := fighter.New(&fighter.Config{}, nil, nil)
	f2 := fighter.New(&fighter.Config{}, nil, nil)
	testFight.Join(f1, f2)

	if testFight.CanJoin() {
		t.Errorf("excpected not availble acces to join the fight")
	}
}

func TestFight_Join(t *testing.T) {
	testFight := fight.New(&fight.Config{FightersLimit: 2})

	f1 := fighter.New(&fighter.Config{}, nil, nil)
	f2 := fighter.New(&fighter.Config{}, nil, nil)
	f3 := fighter.New(&fighter.Config{}, nil, nil)

	if err := testFight.Join(f1, f2); err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	testFight2 := fight.New(&fight.Config{FightersLimit: 2})

	if err := testFight2.Join(f1, f2, f3); err == nil {
		t.Error("expected error, got no error")
	}

}

func TestFight_StartNewRound(t *testing.T) {
	fightOnTheStreet := fight.New(&fight.Config{FightersLimit: 2})

	if err := fightOnTheStreet.StartNewRound(); err == nil {
		t.Errorf("expected an error: %s, got no error", fight.ErrFightNotGoing)
	}

	fightOnTheStreet.SetStatus(fight.Going)
	if err := fightOnTheStreet.StartNewRound(); err == nil {
		t.Errorf("expected an error: %s, got no erro", fight.ErrFightersCountMustNotBeLessThanTwo)
	}

	fightOnTheStreet.Join(newSimpleFighter())
	fightOnTheStreet.SetStatus(fight.Going)
	if err := fightOnTheStreet.StartNewRound(); err == nil {
		t.Errorf("expected an error: %s, got no erro", fight.ErrFightersCountMustNotBeLessThanTwo)
	}

	fightOnTheStreet.Join(newSimpleFighter())
	if err := fightOnTheStreet.StartNewRound(); err != nil {
		t.Errorf("expected no error, got: %s", err)
	}
}

func TestFight_OneOnOne(t *testing.T) {

	const (
		firstFighterNum  = 0
		secondFighterNum = 1
		blocksCount      = 2
		attacksCount     = 1
	)

	greatBattle := fight.New(&fight.Config{FightersLimit: 2})

	greatBattle.Join(newSimpleFighter())
	greatBattle.Join(newSimpleFighter())

	greatBattle.SetStatus(fight.Going)

	firstFighter := greatBattle.FightersList()[firstFighterNum]
	secondFighter := greatBattle.FightersList()[secondFighterNum]

	for {
		if greatBattle.Status() == fight.Finished {
			break
		}

		if err := greatBattle.StartNewRound(); err != nil {
			t.Errorf("expected no errors, got: %s", err)
			break
		}

		// the fighter acts
		enemyNumForFirst := greatBattle.ShowEnemy(firstFighterNum)
		firstFightersBlocks := bot.RandActions(len(firstFighter.BodyParts()), blocksCount)
		firstsEnemiesAttackedParts := bot.RandActions(len(greatBattle.FightersList()[enemyNumForFirst].BodyParts()), attacksCount)

		if err := greatBattle.Actions(firstFighterNum, enemyNumForFirst, firstFightersBlocks, firstsEnemiesAttackedParts); err != nil {
			t.Errorf("expected no errors, got: %s", err)
		}

		// the fighter acts
		enemyNumFroSecond := greatBattle.ShowEnemy(secondFighterNum)
		secondFightersBlocks := bot.RandActions(len(secondFighter.BodyParts()), blocksCount)
		secondsEnemiesAttackedParts := bot.RandActions(len(greatBattle.FightersList()[enemyNumFroSecond].BodyParts()), attacksCount)

		if err := greatBattle.Actions(secondFighterNum, enemyNumFroSecond, secondFightersBlocks, secondsEnemiesAttackedParts); err != nil {
			t.Errorf("expected no errors, got: %s", err)
		}

	}
}
