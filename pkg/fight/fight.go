package fight

import (
	"fmt"

	"github.com/lks-go/terra/pkg/fighter"
)

type Fight interface {
	// SetStatus sets a new status to the Created fight
	SetStatus(Status) error

	// StartNewRound starts a new round within the fight
	StartNewRound() error

	// CanJoin tells to fighter is it possible to join the fight
	CanJoin() bool

	// Join join's a gamer to the fight
	// this method is used to let new fighters to take a part in a started and not finished  yet fight
	Join(...fighter.Fighter) error

	// ShowEnemy returns for the fighter an enemy which must be attacked
	// input param is a number of the current fighter
	// response params the order number of attacked fighter
	ShowEnemy(int) int

	// Actions receives command from a fighter who and where to hit
	// input param:
	//	1. attacking fighter's id;
	//	2. attacked fighter's id;
	//	3. the first fighter's blocks;
	//	4. a list of places where second fighter will be attacked
	Actions(int, int, []int, []int) error

	// Status returns fight's status
	Status() Status

	// FightersList returns fighters id list
	FightersList() []fighter.Fighter
}

func New(cfg *Config) Fight {
	if cfg.FightersLimit == 0 {
		cfg.FightersLimit = defaultFightersLimit
	}

	f := &fight{
		cfg:    cfg,
		rounds: make([]Round, 0),
	}

	return f
}

type fight struct {
	status   Status
	cfg      *Config
	fighters []fighter.Fighter
	rounds   []Round
}

func (f *fight) Actions(fighter int, enemy int, blocks []int, attacks []int) error {
	cr := f.currentRound()
	if cr == nil {
		return errFailedToGetCurrentRound
	}

	return cr.AddAction(fighter, enemy, blocks, attacks)
}

func (f *fight) StartNewRound() error {
	if f.status != Going {
		return errFightNotGoing
	}

	if currentRound := f.currentRound(); currentRound != nil && !currentRound.Finished() {
		return fmt.Errorf("%s: %w", errFailedToCreateNewRound, errRoundNotFinished)
	}

	fnList := make([]int, 0)
	for fighterNumber, _ := range f.aliveFighters() {
		fnList = append(fnList, fighterNumber)
	}

	r, err := NewRound(fnList)
	if err != nil {
		return fmt.Errorf("%s: %w", errFailedToCreateNewRound, err)
	}

	f.rounds = append(f.rounds, r)

	return nil
}

func (f *fight) aliveFighters() []fighter.Fighter {
	survivors := make([]fighter.Fighter, 0)

	for _, fighter := range f.fighters {
		if fighter.CurrentHealth() > 0 {
			survivors = append(survivors, fighter)
		}
	}

	return survivors
}

func (f *fight) ShowEnemy(fighterNumber int) int {
	currentRound := f.currentRound()

	if currentRound == nil {
		return -1
	}

	n := currentRound.ShowEnemyNumber(fighterNumber)

	return n
}

func (f *fight) CanJoin() bool {
	if len(f.fighters) >= f.cfg.FightersLimit {
		return false
	}

	return true
}

func (f *fight) Join(newFighters ...fighter.Fighter) error {
	if f.cfg.FightersLimit == len(f.fighters) {
		return errNoFreePlaces
	}

	if f.cfg.FightersLimit < len(f.fighters)+len(newFighters) {
		return errToManyFightersTriedToJoinTheFight
	}

	f.fighters = append(f.fighters, newFighters...)

	return nil
}

func (f *fight) SetStatus(s Status) error {
	_, ok := statusList[s]
	if !ok || s == Unknown {
		return errUnknownStatus
	}

	f.status = s

	return nil
}

func (f *fight) Status() Status {
	return f.status
}

func (f *fight) FightersList() []fighter.Fighter {
	return f.fighters
}

func (f *fight) currentRound() Round {
	if len(f.rounds) == 0 {
		return nil
	}

	return f.rounds[len(f.rounds)-1]
}
