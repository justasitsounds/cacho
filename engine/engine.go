package engine

import (
	"math/rand"
	"sort"
	"time"
)

type Dice struct {
	randSeed int64
	dice     []int
}

func NewDice(count int) *Dice {
	return &Dice{
		randSeed: time.Now().UnixNano(),
		dice:     make([]int, count),
	}
}

func (d *Dice) SetSeed(seed int64) {
	d.randSeed = seed
}

func (d *Dice) Roll() {
	r := rand.New(rand.NewSource(d.randSeed))
	for i, _ := range d.dice {
		d.dice[i] = r.Intn(5) + 1
	}
}

func (d *Dice) Score() []int {
	sort.Sort(sort.Reverse(sort.IntSlice(d.dice)))
	return d.dice
}

type Player struct {
	dice Dice
}

func (p *Player) RollDice() {

}
