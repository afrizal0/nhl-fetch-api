package main

import (
	"io"
	"log"
	nhlfetchproject "nhlfetchproject/nhlApi"
	"os"
	"time"
)

func main() {
	now := time.Now()
	rosterFile, err := os.OpenFile("roster.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening the file: %v", err)
	}
	defer rosterFile.Close()
	wrt := io.MultiWriter(os.Stdout, rosterFile)
	log.SetOutput(wrt)
	teams, err := nhlfetchproject.GetAllTeams()
	if err != nil {
		log.Fatalf("Error while getting all teams: %v", err)
	}

	for _, team := range teams {
		log.Println("---------------------")
		log.Println("Name :", team.Name)
		log.Println("---------------------")
	}
	log.Printf("too %v", time.Now().Sub(now).String())
}
