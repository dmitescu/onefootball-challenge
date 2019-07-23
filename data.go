package main

import (
	"strconv"
	"strings"
	"errors"
)

type QueryResponse struct {
	Message string `json:"message"`
	Code int `json:"code"`
	Status string `json:"status"`
	Data TeamEntry  `json:"data"`
}

type TeamEntry struct {
	Team TeamData `json:"team"`
}

type TeamData struct {
	Name string `json:"name"`
	Id int `json:"id"`
	Players []PlayerEntry `json:"players"`
}

type PlayerEntry struct {
	Id string `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Age string `json:"age"`
}

type Player struct {
	Id int
	Name string
	Age int
	Teams []string
}

func (p *Player) EnrichPlayer(entry PlayerEntry, teamName string) error {
	if (p.Id == 0) {
		var intId int
		var fullName string
		var age int
		
		var err error
		
		intId, err = strconv.Atoi(entry.Id)
		if (err != nil) { 
			return err
		}
		
		fullName = strings.Join(
			[]string{
				entry.FirstName,
				entry.LastName,
			}, " ")
		
		age, err = strconv.Atoi(entry.Age)
		if (err != nil) {
			return err
		}

		p.Id = intId
		p.Name = fullName
		p.Age = age
		p.Teams = []string{teamName}
	} else {
		var intId int
		var err error
		
		intId, err = strconv.Atoi(entry.Id)		

		if (err != nil) {
			return err
		}
		
		if (p.Id != intId) {
			return errors.New("validity check failed")
		}

		p.Teams = append(p.Teams, teamName)
	}
	return nil
}

func (p *Player) ToString() string {
	name := p.Name
	ageStr := strconv.Itoa(p.Age)
	teamsStr := strings.Join(p.Teams, ", ")
	
	return strings.Join(
		[]string{
			name, ageStr, teamsStr,
		}, "; ")
}
