package fight_test

import (
	"fmt"
	"testing"

	"github.com/lks-go/terra/pkg/fighter"

	"github.com/lks-go/terra/pkg/fight"
)

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
			if err := f.SetStatus(tt.StatusCode); err == nil && tt.IsError {
				t.Errorf("expected error, got nil")
			}

			// test must be failed before than escape the function
			if tt.IsError {
				return
			}

			gotStatus := f.Status()
			if gotStatus != tt.StatusCode {
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
