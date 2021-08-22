package fight

import "errors"

// fight
const (
	ErrNoFreePlaces                      = "no free places"
	ErrToManyFightersTriedToJoinTheFight = "to many fighters tried to join the fight"
	ErrNoRoundsInFight                   = "there is no rounds on fight yet"
	ErrFailedToGetCurrentRound           = "failed to get current round"
)

// round
const (
	ErrCantActInThisRoundAnymore         = "can't make actions in this round anymore"
	ErrFightersCountMustNotBeEqualToZero = "fighters count must not be equal to zero"
	ErrFightersCountMustNotBeLessThanTwo = "fighters count must not be less than two"
	ErrFailedToCreateNewRound            = "failed to create a new round"
	ErrRoundNotFinished                  = "round now finished yet"
)

// statuses
const (
	ErrUnknownStatus         = "unknown status"
	ErrFightNotStarted       = "the fight not started yet"
	ErrFightNotGoing         = "the fight is not going"
	ErrFightHasUnknownStatus = "the fight has unknown status"
	ErrFightFinished         = "the fight already finished"
)

var (
	errNoFreePlaces                      = errors.New(ErrNoFreePlaces)
	errUnknownStatus                     = errors.New(ErrUnknownStatus)
	errCantActInThisRoundAnymore         = errors.New(ErrCantActInThisRoundAnymore)
	errFightersCountMustNotBeEqualToZero = errors.New(ErrFightersCountMustNotBeEqualToZero)
	errFightersCountMustNotBeLessThanTwo = errors.New(ErrFightersCountMustNotBeLessThanTwo)
	errToManyFightersTriedToJoinTheFight = errors.New(ErrToManyFightersTriedToJoinTheFight)
	errFightNotStarted                   = errors.New(ErrFightNotStarted)
	errFightFinished                     = errors.New(ErrFightFinished)
	errFightHasUnknownStatus             = errors.New(ErrFightHasUnknownStatus)
	errFailedToCreateNewRound            = errors.New(ErrFailedToCreateNewRound)
	errRoundNotFinished                  = errors.New(ErrRoundNotFinished)
	errNoRoundsInFight                   = errors.New(ErrNoRoundsInFight)
	errFightNotGoing                     = errors.New(ErrFightNotGoing)
	errFailedToGetCurrentRound           = errors.New(ErrFailedToGetCurrentRound)
)
