package engine

import (
	"math/rand"
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

type Player struct {
	dice Dice
}

func (p *Player) RollDice() {

}
