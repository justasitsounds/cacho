package engine

import (
	"math/rand"
	"sort"
	"time"
)

// Dice are the group of one or more Die's
type Dice struct {
	randSeed int64
	dice     []int
}

//NewDice returns a new set of Dice, with (count) number of Die's
func NewDice(count int) *Dice {
	return &Dice{
		randSeed: time.Now().UnixNano(),
		dice:     make([]int, count),
	}
}

// SetSeed sets the RNG seed that determines how the bones fall
func (d *Dice) SetSeed(seed int64) {
	d.randSeed = seed
}

// Roll makes them bones roll
func (d *Dice) Roll() {
	r := rand.New(rand.NewSource(d.randSeed))
	for i := range d.dice {
		d.dice[i] = r.Intn(5) + 1
	}
}

//Score the Dice
func (d *Dice) Score() []int {
	sort.Sort(sort.Reverse(sort.IntSlice(d.dice)))
	return d.dice
}

//Player No prizes
type Player struct {
	dice Dice
	name string
}
