package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/justasitsounds/cacho/engine"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("How many players?")

	line, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("  ==  ", err.Error())
	}

	playerCount, err := strconv.Atoi(string(line))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(playerCount)

	playerNames := make([]string, playerCount)
	for i := 0; i < playerCount; i = i + 1 {
		playerNames[i] = strconv.Itoa(i + 1)
	}
	var g = game.New(playerNames)
}

type game struct {
	players []*engine.Player
}

func (g *game) New(playerNames []string) *game {

	players := make([]*engine.Player, len(playerNames))
	for name := range playerNames {
		players = append(players, engine.NewPlayer(name))
	}
	return &game{players}
}
