package players

import (
	"encoding/csv"
	"log"
	"os"
)

type Player struct {
	Name  string
	Email string
}

func ImportPlayers(file string) []Player {
	var ps []Player

	f, err := os.Open(file)
	if err != nil {
		log.Fatalln("Could not open file", err)
	}

	rdr := csv.NewReader(f)
	rows, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln("Could not read CSV", err)
	}

	for _, v := range rows {
		ps = append(ps, Player{
			Name:  v[0],
			Email: v[1],
		})
	}

	return ps
}
