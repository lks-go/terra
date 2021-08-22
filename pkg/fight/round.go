package fight

import (
	"fmt"
	"sync"
)

type Round interface {
	// Finished returns information if a round finished or not
	Finished() bool

	// AddAction
	// params:
	// 	1 - fighter's number;
	//	2 - enemy's number;
	//	3 - blocks
	//	4 - attacks
	AddAction(int, int, []int, []int) error

	// Actions returns a list of fighters actions made during the fight
	Actions() []action

	// ShowEnemyNumber returns a fighters number which not attacked yet in this round
	// -1 means there are no enemies to attack anymore
	ShowEnemyNumber(int) int
}

func NewRound(fightersNumbers []int) (Round, error) {
	if len(fightersNumbers) < 2 {
		return nil, errFightersCountMustNotBeLessThanTwo
	}

	e := make([]int, 0)
	ta := make(map[string]int)
	for _, fn := range fightersNumbers {
		e = append(e, fn)
		ta[fmt.Sprintf("%d", fn)] = 0
	}

	r := &round{
		fighters:           e,
		totalActs:          ta,
		actions:            make([]action, 0),
		maxPossibleActions: (len(fightersNumbers) * (len(fightersNumbers) - 1)) / len(fightersNumbers),
	}

	return r, nil
}

type action struct {
	fighter int
	enemy   int
	blocks  []int
	attacks []int
}

type round struct {
	// fighters contains list of fighters
	fighters []int
	// totalActs contains list of fighters actions against their fighters
	totalActs map[string]int
	// actions is a list of fighters actions
	actions []action
	// maxPossibleActions returns actions count which each fighter can make against each enemy
	maxPossibleActions int

	mu sync.Mutex
}

func (r *round) ShowEnemyNumber(fighterNumber int) int {

MainLoop:
	for _, checkedEnemyNumber := range r.fighters {
		if fighterNumber == checkedEnemyNumber {
			continue
		}

		for _, completedAction := range r.actions {
			if completedAction.fighter != fighterNumber {
				continue
			}

			if completedAction.enemy == checkedEnemyNumber {
				continue MainLoop
			}
		}

		return checkedEnemyNumber
	}

	return -1
}

func (r *round) AddAction(fighterNumber int, enemyNumber int, blockedBodyParts []int, attackedBodyParts []int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.maxPossibleActions == r.totalActs[fmt.Sprintf("%d", fighterNumber)] {
		return errCantActInThisRoundAnymore
	}

	r.actions = append(r.actions, action{
		fighter: fighterNumber,
		enemy:   enemyNumber,
		blocks:  blockedBodyParts,
		attacks: attackedBodyParts,
	})

	r.totalActs[fmt.Sprintf("%d", fighterNumber)]++

	return nil
}

func (r *round) Finished() bool {
	for fn := range r.totalActs {
		if r.totalActs[fn] < r.maxPossibleActions {
			return false
		}
	}

	return true
}

func (r *round) Actions() []action {
	return r.actions
}
