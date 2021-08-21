package fight

import (
	"fmt"
	"sync"
)

type Round interface {
	// Finished returns information if a round finished or not
	Finished() bool

	// AddAction
	// 1 - fighter's number, 2 - enemy's number, 3 - blocks of own parts, 4 - attacks of enemy parts
	AddAction(int, int, []int, []int) error

	// ActionsList returns list of blocks and attacks
	ActionsList() ([]action, []action)
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
		enemies:            e,
		totalActs:          ta,
		blocks:             make([]action, 0),
		attacks:            make([]action, 0),
		maxPossibleActions: (len(fightersNumbers) * (len(fightersNumbers) - 1)) / len(fightersNumbers),
	}

	return r, nil
}

type action struct {
	fighter int
	enemy   int
	parts   []int
}

type round struct {
	// enemies contains list of fighters
	enemies []int
	// totalActs contains list of fighters actions against their enemies
	totalActs map[string]int
	// blocks is a list of fighters blocks against enemies
	blocks []action
	// attacks is a list of fighters attacks against enemies
	attacks []action
	// maxPossibleActions returns actions count which each fighter can make against each enemy
	maxPossibleActions int

	mu sync.Mutex
}

func (r *round) AddAction(fighterNumber int, enemyNumber int, blockedBodyParts []int, attackedBodyParts []int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.maxPossibleActions == r.totalActs[fmt.Sprintf("%d", fighterNumber)] {
		return errCantActInThisRoundAnymore
	}

	r.blocks = append(r.blocks, action{
		fighter: fighterNumber,
		enemy:   enemyNumber,
		parts:   blockedBodyParts,
	})

	r.attacks = append(r.attacks, action{
		fighter: fighterNumber,
		enemy:   enemyNumber,
		parts:   attackedBodyParts,
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

func (r *round) ActionsList() ([]action, []action) {
	return r.blocks, r.attacks
}
