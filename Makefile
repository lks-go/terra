tests:
	go test ./...

testsv: test_fighter test_fight test_bot

test_fighter:
	go test ./pkg/fighter -v

test_fight:
	go test ./pkg/fight -v

test_bot:
	go test ./pkg/bot -v
