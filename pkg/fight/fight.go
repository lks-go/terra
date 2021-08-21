package fight

import (
	"github.com/lks-go/terra/pkg/fighter"
)

type Fight interface {
	// SetStatus sets a new status to the Created fight
	SetStatus(int) error

	// CanJoin tells to fighter is it possible to join the fight
	CanJoin() bool

	// Join join's a gamer to the fight
	// this method is used to let new fighters to take a part in a started and not finished  yet fight
	Join(...fighter.Fighter) error

	// Enemy returns a fighter which must be attacked
	// 1. fight order number of attacked fighter; 2. a list of places where the fighter will be attacked
	Enemy() (int, []string)

	// SetBlocks sets fighters blocks
	SetBlocks(int, []string)

	// Attack receives command from a fighter who and where to hit
	// 1. attacking fighter's id; 2 attacked fighter's id; 3. a list of places where second fighter will be attacked
	Attack(string, []string, string, []string) error

	// Status returns fight's status
	Status() int

	// FightersList returns fighters id list
	FightersList() []fighter.Fighter
}

func New(cfg *Config) Fight {
	if cfg.FightersLimit == 0 {
		cfg.FightersLimit = defaultFightersLimit
	}

	f := &fight{
		cfg: cfg,
	}

	return f
}

type fight struct {
	status   Status
	cfg      *Config
	fighters []fighter.Fighter
}

func (f *fight) Enemy() (int, []string) {
	panic("implement me")
}

func (f *fight) SetBlocks(n int, strings []string) {
	panic("implement me")
}

func (f *fight) Attack(s string, strings []string, s2 string, strings2 []string) error {
	panic("implement me")
}

func (f *fight) CanJoin() bool {
	if len(f.fighters) >= f.cfg.FightersLimit {
		return false
	}

	return true
}

func (f *fight) Join(fighters ...fighter.Fighter) error {
	if f.cfg.FightersLimit == len(f.fighters) {
		return errNoFreePlaces
	}

	f.fighters = append(f.fighters, fighters...)

	return nil
}

func (f *fight) SetStatus(s int) error {
	_, ok := statusList[Status(s)]
	if !ok || Status(s) == Unknown {
		return errUnknownStatus
	}

	f.status = Status(s)

	return nil
}

func (f *fight) Status() int {
	return int(f.status)
}

func (f *fight) FightersList() []fighter.Fighter {
	return f.fighters
}
