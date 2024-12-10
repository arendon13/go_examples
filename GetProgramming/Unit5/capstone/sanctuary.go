package main

import (
	"fmt"
	"math/rand"
	"time"
)

type animal interface {
	move() string
	eat() string
}

type cat struct {
	name string
}

func (c cat) move() string {
	switch rand.Intn(2) {
	case 0:
		return "struts with catlike pride"
	default:
		return "runs away from the sprinklers"
	}
}

func (c cat) eat() string {
	switch rand.Intn(2) {
	case 0:
		return "eats Purina Fancy Feast"
	default:
		return "eats bland generic cat food"
	}
}

func (c cat) String() string {
	return c.name
}

type dog struct {
	name string
}

func (d dog) move() string {
	switch rand.Intn(2) {
	case 0:
		return "walks without a care in the world"
	default:
		return "chases a squirrel"
	}
}

func (d dog) eat() string {
	switch rand.Intn(2) {
	case 0:
		return "eats Pedigree dry dog food"
	default:
		return "eats gourmet wet dog food"
	}
}

func (d dog) String() string {
	return d.name
}

type parrot struct {
	name string
}

func (p parrot) move() string {
	switch rand.Intn(2) {
	case 0:
		return "takes flight"
	default:
		return "hops around"
	}
}

func (p parrot) eat() string {
	switch rand.Intn(2) {
	case 0:
		return "eat Premium Encore bird food"
	default:
		return "eats bland generic bird food"
	}
}

func (p parrot) String() string {
	return p.name
}

const sunrise, sunset = 6, 17

func main() {
	rand.Seed(time.Now().UnixNano())

	animals := []animal{
		cat{name: "White-Shoes"},
		dog{name: "Rosco"},
		parrot{name: "Iago"},
	}

	var sol, hour int

	for {
		fmt.Printf("%2d:00 ", hour)
		if hour < sunrise || hour >= sunset {
			fmt.Println("All the animals are sleeping...")
		} else {
			i := rand.Intn(len(animals))
			performAction(animals[i])
		}

		time.Sleep(500 * time.Millisecond)

		hour++

		// stopping condition
		if hour >= 24 {
			hour = 0
			sol++
			if sol >= 3 {
				break
			}
		}
	}
}

func performAction(a animal) {
	switch rand.Intn(2) {
	case 0:
		fmt.Printf("%v %v.\n", a, a.move())
	default:
		fmt.Printf("%v %v.\n", a, a.eat())
	}
}
