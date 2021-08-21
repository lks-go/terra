package fighter

import (
	"fmt"
	"testing"
)

const (
	testGamerName     = "Tester"
	testHeadMockName  = "Big head"
	testChestMockName = "Wide chest"
	testGroinMockName = "Steel balls"
	testFeetMockName  = "Short feet"
)

func newDefaultFighterMock(cfg *Config) Fighter {
	bodyParts := make([]DamageGetter, 0)
	bodyParts = append(bodyParts, NewPart(&PartConfig{Name: testHeadMockName}))
	bodyParts = append(bodyParts, NewPart(&PartConfig{Name: testChestMockName}))
	bodyParts = append(bodyParts, NewPart(&PartConfig{Name: testGroinMockName}))
	bodyParts = append(bodyParts, NewPart(&PartConfig{Name: testFeetMockName}))

	return New(cfg, bodyParts, nil)
}

func TestFighter_CurrentHealth(t *testing.T) {

	var tests = []struct {
		damage               int32
		health               int32
		expectedRestOfHealth int32
	}{
		{10, 50, 40},
		{3, 10, 7},
		{100, 50, -50},
		{20, 20, 0},
	}

	for _, tt := range tests {

		testName := fmt.Sprintf("health: %d; damage: %d, expected rest of: %d", tt.health, tt.damage, tt.expectedRestOfHealth)
		t.Run(testName, func(t *testing.T) {
			cfg := &Config{
				Name:   testGamerName,
				Health: tt.health,
			}

			fighter := newDefaultFighterMock(cfg)

			bpList := fighter.BodyParts()

			bpList[0].CatchDamage(tt.damage, false)

			if fighter.CurrentHealth() != tt.expectedRestOfHealth {
				t.Errorf("current health %d expected: %d", fighter.CurrentHealth(), tt.expectedRestOfHealth)
			}

		})
	}

}

func TestFighter_SimpleKilling(t *testing.T) {
	cfg := &Config{
		Name:   testGamerName,
		Health: 50,
	}

	fighter := newDefaultFighterMock(cfg)

	hits := []struct {
		damage int32
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
		t.Errorf("Expected health 0, got: %d", fighter.CurrentHealth())
	}
}

func TestFighter_SimpleBlock(t *testing.T) {
	var health, damage int32 = 50, 10

	cfg := &Config{
		Name:   testGamerName,
		Health: health,
	}

	fighter := newDefaultFighterMock(cfg)

	fighter.BodyParts()[0].Block()
	fighter.BodyParts()[0].CatchDamage(damage, false)

	if fighter.CurrentHealth() != health {
		t.Errorf("Expected health %d, got: %d", health, fighter.CurrentHealth())
	}
}
