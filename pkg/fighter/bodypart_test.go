package fighter_test

import (
	"testing"

	"github.com/lks-go/terra/pkg/fighter"
)

const (
	testPartName = "test part"
	testDamage   = 7
)

func getSimpleTestPart() fighter.Part {
	leCfg := &fighter.PartConfig{
		Name: testPartName,
	}

	bp := fighter.NewPart(leCfg)

	owner := fighter.New(&fighter.Config{}, nil, nil)
	bp.SetOwner(owner)

	return bp
}

func TestPart_Name(t *testing.T) {
	leg := getSimpleTestPart()
	if leg.Name() != testPartName {
		t.Errorf("body part name \"%s\" is not correspond to config name \"%s\"", leg.Name(), testPartName)
	}
}

func TestPart_TakeDamage(t *testing.T) {
	leg := getSimpleTestPart()
	if leg.CatchDamage(testDamage, false) != testDamage {
		t.Errorf("body part name \"%s\" is not correspond to config name \"%s\"", leg.Name(), testPartName)
	}
}

func TestPart_BlockDamage(t *testing.T) {
	head := getSimpleTestPart()
	head.Block()

	res := head.CatchDamage(10, false)

	if res != 0 {
		t.Errorf("expected damage %d, got %d", 0, res)
	}
}

func TestPart_CriticalDamageToBlock(t *testing.T) {
	damage := 100
	head := getSimpleTestPart()
	head.Block()
	res := head.CatchDamage(damage, true)

	if res != damage {
		t.Errorf("expected damage of %d, got %d", damage, res)
	}
}
