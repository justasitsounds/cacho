package engine

import "testing"

func deterministicDice(seed int64) *Dice {
	d := NewDice(5)
	d.SetSeed(seed)
	d.Roll()
	return d
}

func setDice(setdice []int) *Dice {
	return &Dice{
		dice: setdice,
	}
}

func TestDice(t *testing.T) {
	d := deterministicDice(23)

	for _, die := range d.dice {
		if die == 0 {
			t.Fail()
		}
	}
}

func TestScoreDice(t *testing.T) {
	var diceTests = []struct {
		setDice       []int
		expectedScore []int
	}{
		{[]int{1, 3, 4, 3, 1}, []int{3, 3, 3, 3}},
		{[]int{2, 3, 4, 5, 1}, []int{5, 5}},
		{[]int{2, 2, 4, 2, 1}, []int{2, 2, 2, 2}},
		{[]int{6, 2, 4, 3, 5}, []int{6}},
		{[]int{5, 3, 5}, []int{5, 5}},
		{[]int{3}, []int{3}},
		{[]int{6, 6, 5, 5, 1}, []int{6, 6, 6}},
		{[]int{6, 6, 6, 5, 1}, []int{6, 6, 6, 6}},
		{[]int{2, 2, 6, 5, 1}, []int{2, 2, 2}},
		{[]int{2, 2, 5}, []int{2, 2}},
	}

	for _, tt := range diceTests {
		actual := setDice(tt.setDice).Score()
		if len(actual) != len(tt.expectedScore) {
			t.Errorf("Expected score: %v, but got: %v", tt.expectedScore, actual)
			break
		}
		for i, ad := range actual {
			if ad != tt.expectedScore[i] {
				t.Errorf("Expected die at position:%v to be %v, but got %v", i, ad, tt.expectedScore)
			}
		}
	}
}
