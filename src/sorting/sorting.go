package main

import (
	"fmt"
	"github.com/ThompsonJonM/fantasy-emailer/m/v2/src/players"
	"math/rand"
	"time"
)

func main() {
	ps := players.ImportPlayers("../../csvFiles/leagueFour.csv")

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ps), func(i, j int) {
		ps[i], ps[j] = ps[j], ps[i]
	})

	fmt.Println(ps)
}
