tests:
	go test ./...

testsv: test_fighter test_fight

test_fighter:
	go test ./pkg/fighter -v

test_fight:
	go test ./pkg/fight -v
