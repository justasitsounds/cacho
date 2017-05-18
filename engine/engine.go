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

func (d *Dice) Count(value int) int {
	var count int
	for i := range d.dice {
		diceVal := d.dice[i]
		if diceVal == 1 || diceVal == value {
			count = count + 1
		}
	}
	return count
}

func (d *Dice) Score() []int {
	sort.Sort(sort.Reverse(sort.IntSlice(d.dice)))
	var prevSlots []int
	for i := 0; i < len(d.dice); i = i + 1 {
		var slots []int
		var val = d.dice[i] //4
		for j := 0; j < d.Count(val); j = j + 1 {
			slots = append(slots, val)
		}
		if len(slots) > len(prevSlots) {
			prevSlots = slots
		}
	}
	return prevSlots
}

type Player struct {
	dice Dice
}

func (p *Player) RollDice() {

}
