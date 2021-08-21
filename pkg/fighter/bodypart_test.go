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
	if head.CatchDamage(10, false) != 0 {
		t.Errorf("expected ")
	}
}

// TODO test critical damage on blocked body part
// TODO test critical damage on unblocked body part
// TODO test passing damage to owner
