package fight

import (
	"github.com/lks-go/terra/pkg/fighter"
)

type Fight interface {
	// SetStatus starts the Created fight
	SetStatus(int32) error

	// CanJoin tells to fighter is it possible to join the fight
	CanJoin() bool

	// Join join's a gamer to the fight
	Join(...fighter.Fighter) error

	// Enemy returns a fighter which must be attacked
	// 1. fight order number of attacked fighter; 2. a list of places where the fighter will be attacked
	Enemy() (int32, []string)

	// SetBlocks sets fighters blocks
	SetBlocks(int32, []string)

	// Attack receives command from a fighter who and where to hit
	// 1. attacking fighter's id; 2 attacked fighter's id; 3. a list of places where second fighter will be attacked
	Attack(string, []string, string, []string) error

	// Status returns fight's status
	Status(string) int32

	// FightersList returns fighters id list
	FightersList() []fighter.Fighter
}

type Params map[string]string

func New(cfg *Config, params Params) Fight {
	f := &fight{}

	// TODO here parallel programming starts

	return f
}

type fight struct {
	status   Status
	limit    int
	fighters []fighter.Fighter
}

func (f *fight) SetStatus(s int32) error {
	_, ok := statusList[Status(s)]
	if !ok {
		return errUnknownStatus
	}

	f.status = Status(s)

	return nil
}

func (f *fight) CanJoin() bool {
	if len(f.fighters) >= f.limit {
		return false
	}

	return true
}

func (f *fight) Join(fighters ...fighter.Fighter) error {
	if f.limit == len(f.fighters) {
		return errNoFreePlaces
	}
	f.fighters = append(f.fighters, fighters...)

	return nil
}

func (f *fight) Enemy() (int32, []string) {
	panic("implement me")
}

func (f *fight) SetBlocks(n int32, strings []string) {
	panic("implement me")
}

func (f *fight) Attack(s string, strings []string, s2 string, strings2 []string) error {
	panic("implement me")
}

func (f *fight) Status(s string) int32 {
	return int32(f.status)
}

func (f *fight) FightersList() []fighter.Fighter {
	return f.fighters
}
