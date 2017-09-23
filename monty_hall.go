package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// true if a car is behind the door, false otherwise
type Door bool

// NewDoors creates a set of n Doors. One Door is true.
// Returns the set of Doors and the index of the correct door.
func NewDoors(n int) ([]Door, error) {
	doors := make([]Door, n)

	if n < 3 {
		return doors, errors.New("Minimum 3 doors.")
	}

	answer := rand.Intn(n)
	doors[answer] = true

	return doors, nil
}

// reduceDoors takes a set of doors and an int indicating which door is chosen.
// Returns an index of another door.
func reduceDoors(d []Door, c int) (o int) {
	if d[c] == true {
		// choose any other door; they're all wrong
		for o == c {
			o = rand.Intn(len(d))
		}
		return
	}

	for i, e := range d {
		if e == true {
			o = i
			break
		}
	}
	return
}

// SimulateGameShow takes an int n indicating the number of doors to use,
// and a bool sd indicating whether we should switch doors.
// Returns a bool whether or not we won.
func SimulateGameShow(n int, sd bool) bool {

	d, err := NewDoors(n)
	if err != nil {
		log.Fatal(err)
	}

	// choose a random door
	c := rand.Intn(len(d))

	other := reduceDoors(d, c)

	if sd {
		c = other
	}

	return d[c] == true

}

func main() {
	var (
		doors int
		sw    string
		loops int
	)

	fmt.Print("How many doors should we use? ")
	fmt.Scanln(&doors)

	fmt.Print("Should we switch doors? (y/n) ")
	fmt.Scanln(&sw)
	sd := sw == "y"

	fmt.Print("How many iterations? ")
	fmt.Scanln(&loops)

	var wins []bool

	for i := 0; i < loops; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		w := SimulateGameShow(doors, sd)
		wins = append(wins, w)
	}

	fmt.Println(wins)

}
