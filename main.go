package main

import (
	"log"
	"fmt"
	"sort"
)

type ByName []Player

func (p ByName) Len() int           { return len(p) }
func (p ByName) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByName) Less(i, j int) bool { return p[i].Name < p[j].Name }

func main() {
	log.Println("starting querying service")

	var players map[string]*Player = make(map[string]*Player)
	responses := TeamFinder()

	for _, resp := range responses {
		for _, entry := range resp.Data.Team.Players {
			newPlayer := &Player{}
			err := newPlayer.EnrichPlayer(entry, resp.Data.Team.Name)

			if (err != nil) {
				log.Println("error: ", err)
			}

			player, exists := players[newPlayer.Name]

			if (!exists) {
				players[newPlayer.Name] = newPlayer
			} else {
				err = player.EnrichPlayer(entry, resp.Data.Team.Name)
				if (err != nil) {
					log.Println("error: ", err)
				}
			}
		}
	}

	var playersList []Player

	for _, p := range players {
		playersList = append(playersList, *p)
	}

	sort.Sort(ByName(playersList))

	for i, p := range playersList {
		fmt.Printf("%d. %s\n", i + 1, p.ToString())
	}
}
