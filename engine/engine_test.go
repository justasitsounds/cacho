package engine

import (
	"fmt"
	"testing"
)

func TestDice(t *testing.T) {
	d := NewDice(5)
	d.SetSeed(23)
	for _, die := range d.dice {
		if die != 0 {
			t.Fail()
		}
	}
	d.Roll()
	fmt.Println(d.dice[0])
}
