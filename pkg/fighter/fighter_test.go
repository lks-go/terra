package fighter_test

import (
	"fmt"
	"testing"

	"github.com/lks-go/terra/pkg/fighter"
)

const (
	testGamerName     = "Tester"
	testHeadMockName  = "Big head"
	testChestMockName = "Wide chest"
	testGroinMockName = "Steel balls"
	testFeetMockName  = "Short feet"
)

func newDefaultFighterMock(cfg *fighter.Config) fighter.Fighter {
	bodyParts := make([]fighter.DamageGetter, 0)
	bodyParts = append(bodyParts, fighter.NewPart(&fighter.PartConfig{Name: testHeadMockName}))
	bodyParts = append(bodyParts, fighter.NewPart(&fighter.PartConfig{Name: testChestMockName}))
	bodyParts = append(bodyParts, fighter.NewPart(&fighter.PartConfig{Name: testGroinMockName}))
	bodyParts = append(bodyParts, fighter.NewPart(&fighter.PartConfig{Name: testFeetMockName}))

	return fighter.New(cfg, bodyParts, nil)
}

func TestFighter_CurrentHealth(t *testing.T) {

	var tests = []struct {
		damage               int
		health               int
		expectedRestOfHealth int
	}{
		{10, 50, 40},
		{3, 10, 7},
		{100, 50, -50},
		{20, 20, 0},
	}

	for _, tt := range tests {

		testName := fmt.Sprintf("baseHealth: %d; damage: %d, expected rest of: %d", tt.health, tt.damage, tt.expectedRestOfHealth)
		t.Run(testName, func(t *testing.T) {
			cfg := &fighter.Config{
				Name:       testGamerName,
				BaseHealth: tt.health,
			}

			fighter := newDefaultFighterMock(cfg)

			bpList := fighter.BodyParts()

			bpList[0].CatchDamage(tt.damage, false)

			if fighter.CurrentHealth() != tt.expectedRestOfHealth {
				t.Errorf("current baseHealth %d expected: %d", fighter.CurrentHealth(), tt.expectedRestOfHealth)
			}

		})
	}

}

func TestFighter_SimpleKilling(t *testing.T) {
	cfg := &fighter.Config{
		Name:       testGamerName,
		BaseHealth: 50,
	}

	fighter := newDefaultFighterMock(cfg)

	hits := []struct {
		damage int
	}{
		{10},
		{7},
		{8},
		{15},
		{6},
		{4},
	}

	for _, hit := range hits {
		fighter.BodyParts()[0].CatchDamage(hit.damage, false)
	}

	if fighter.CurrentHealth() != 0 {
		t.Errorf("Expected baseHealth 0, got: %d", fighter.CurrentHealth())
	}
}

func TestFighter_SimpleBlock(t *testing.T) {
	var health, damage int = 50, 10

	cfg := &fighter.Config{
		Name:       testGamerName,
		BaseHealth: health,
	}

	fighter := newDefaultFighterMock(cfg)

	fighter.BodyParts()[0].Block()
	fighter.BodyParts()[0].CatchDamage(damage, false)

	if fighter.CurrentHealth() != health {
		t.Errorf("Expected baseHealth %d, got: %d", health, fighter.CurrentHealth())
	}
}

func TestUnit_BodyPartsOrderNumbers(t *testing.T) {

	partsNames := []string{testHeadMockName, testChestMockName, testGroinMockName, testFeetMockName}

	bodyParts := make([]fighter.DamageGetter, len(partsNames))
	for i := 0; i < len(partsNames); i++ {
		bodyParts[i] = fighter.NewPart(&fighter.PartConfig{Name: partsNames[i]})
	}

	cfg := &fighter.Config{
		Name: testGamerName,
	}

	terminator := fighter.New(cfg, bodyParts, nil)

	bodyPartsNumbers := terminator.OrderedBodyPartsNumbers()

	terminatorsParts := terminator.BodyParts()

	for _, n := range bodyPartsNumbers {
		bp := terminatorsParts[n]
		if bp.Name() != partsNames[n] {
			t.Errorf("expected body part: %s, got: %s", partsNames[n], bp.Name())
		}
	}

}
