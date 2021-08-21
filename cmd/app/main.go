package main

import "github.com/lks-go/terra/pkg/fighter"

func main() {

	bodyParts := make([]fighter.DamageGetter, 0)

	bodyParts = append(bodyParts, fighter.NewPart(&fighter.PartConfig{}))

	fighterCfg := &fighter.Config{Name: "lks"}
	_ = fighter.New(fighterCfg, bodyParts, nil)

}
