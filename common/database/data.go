package database

import (
	"math/rand"
)

type Node interface {
	IsNode()
}

type isNode struct{}

func (isNode) IsNode() {}

func init() {
	rand.Seed(100)

	for i := 0; i < 100; i++ {
		name := randomName()

		pets := make([]Pet, rand.Intn(5))
		for j := range pets {
			pets[j] = Pet{
				Name:    randomName(),
				Species: randomSpec(),
			}
		}

		Users[name] = User{
			Name: name,
			ID:   name,
			Pets: pets,
		}
	}
}

func randomName() string {
	return Names[rand.Intn(len(Names))]
}

func randomSpec() string {
	return Species[rand.Intn(len(Species))]
}
