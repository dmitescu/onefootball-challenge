package main

import (
	"errors"
	"sync"
	"time"
	"math/rand"
)

const (
	maxId = 9999
)

func TeamFinderHandler() func() (*QueryResponse, error) {
	currentId := 0
	teamMap := make(map[string]bool)
	var teamMutex sync.Mutex
	var counterMutex sync.Mutex

	var requiredTeams = [...]string{
		"Germany",
		"England",
		"France",
		"Spain",
		"Manchester United",
		"Arsenal",
		"Chelsea",
		"Barcelona",
		"Real Madrid",
		"Bayern Munich",
	}

	for _, team := range requiredTeams {
		teamMap[team] = false
	}
	
	return func() (*QueryResponse, error) {
		counterMutex.Lock()
		currentId = currentId + 1
		counterMutex.Unlock()
		
		var allFound bool = true
		var isMatch bool = false
		
		if (currentId > maxId) {
			return nil, errors.New("maximum id reached")
		}

		for _, f := range teamMap {
			allFound = allFound && f
		}

		if (allFound) {
			return nil, errors.New("all teams have been found")
		}

		response, _ := QueryId(currentId)

		for t, _ := range teamMap {
			if (response.Data.Team.Name == t) {
				teamMutex.Lock()
				teamMap[t] = true
				teamMutex.Unlock()
				
				isMatch = true
			}
		}

		if (isMatch) {
			return response, nil
		} else {
			return nil, nil
		}
	}
}

func TeamFinder() []QueryResponse  {
	handler := TeamFinderHandler()
	rand.Seed(time.Now().UnixNano())
	
	var responses []QueryResponse
	var errGlobal error

	var collectorMutex sync.Mutex
	var errMutex sync.Mutex

	processHandler := func() {
		var response *QueryResponse
		var err error

		response, err = handler()

		if (err != nil) {
			errMutex.Lock()
			errGlobal = err
			errMutex.Unlock()
		}

		if (response != nil) {
			collectorMutex.Lock()
			responses = append(responses, *response)
			collectorMutex.Unlock()
		}
	}
	
	for (errGlobal == nil) {
		go processHandler()
		time.Sleep(time.Millisecond * time.Duration(100 + rand.Intn(400)))
	}

	return responses
}
