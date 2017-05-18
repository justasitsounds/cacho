package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		playerNames[i] = strconv.Itoa(i)
	}
	var g = game.New(playerNames)
}
