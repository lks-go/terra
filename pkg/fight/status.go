package fight

type Status int

const (
	Unknown Status = iota
	Created
	Going
	Finished
)

var statusList = make(map[Status]struct{})

func init() {
	statusList[Unknown] = struct{}{}
	statusList[Created] = struct{}{}
	statusList[Going] = struct{}{}
	statusList[Finished] = struct{}{}
}
