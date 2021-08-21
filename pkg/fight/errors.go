package fight

import "errors"

const (
	ErrNoFreePlaces                      = "no free places"
	ErrUnknownStatus                     = "unknown status"
	ErrCantActInThisRoundAnymore         = "can't make actions in this round anymore"
	ErrFightersCountMustNotBeEqualToZero = "fighters count must not be equal to zero"
	ErrFightersCountMustNotBeLessThanTwo = "fighters count must not be less than two"
)

var (
	errNoFreePlaces                      = errors.New(ErrNoFreePlaces)
	errUnknownStatus                     = errors.New(ErrUnknownStatus)
	errCantActInThisRoundAnymore         = errors.New(ErrCantActInThisRoundAnymore)
	errFightersCountMustNotBeEqualToZero = errors.New(ErrFightersCountMustNotBeEqualToZero)
	errFightersCountMustNotBeLessThanTwo = errors.New(ErrFightersCountMustNotBeLessThanTwo)
)
