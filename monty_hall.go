package main

import (
	"fmt"
	"math/rand"
	"time"
)

// true if a car is behind the door, false otherwise
type Door bool

// New creates a set of n Doors. One Door is true.
// Returns the set of Doors and the index of the correct door.
func New(n int) ([]Door, int) {
	doors := make([]Door, n)

	rand.Seed(time.Now().UTC().UnixNano())
	answer := rand.Intn(n)
	doors[answer] = true

	return doors, answer
}

func main() {
	doors, _ := New(3)
	fmt.Println(doors)
}
