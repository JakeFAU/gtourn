package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := os.Args[1]
	if fileName == "" {
		fileName = "golfdata.txt"
	}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var players []Player
	for scanner.Scan() {
		aLine := scanner.Text()
		tokens := strings.Fields(aLine)
		golferNumber, err := strconv.Atoi(tokens[0])
		if err != nil {
			log.Fatal("Error reading golfer number")
		}
		golferHandicap, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatal("Error reading golfer handicap")
		}
		var golferScores []int
		for _, score := range tokens[2:] {
			_s, err := strconv.Atoi(score)
			if err != nil {
				log.Fatal("Error reading scores")
			}
			golferScores = append(golferScores, _s)
		}
		player := Player{Number: uint(golferNumber),
			Handicap: golferHandicap,
			Scores:   golferScores}
		players = append(players, player)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	t := Tournament{Name: "Test Tournament",
		Players: players}
	printTournament(t)

}

func printTournament(t Tournament) {
	fmt.Println("       ---------------------------------------------------")
	fmt.Printf("       %s\n", t.Name)
	fmt.Println("       ---------------------------------------------------")
	fmt.Println()
	fmt.Println()
	fmt.Println("Player   Handicap       Scores                       Total   Result   Under Par?")
	fmt.Println("------   --------   ------------------               -----   ------   ----------")
	for _, player := range t.Players {
		if player.IsValid() {
			g, _ := player.GrossScore()
			n, _ := player.NetScore()
			p, _ := player.UnderPar(30)
			fmt.Printf("%2v       %2v         %v                %v     %v        %v       \n", player.Number, player.Handicap, player.Scores, g, n, p)
		} else {
			fmt.Printf("%2v       %2v         %v               *** INVALID DATA ***       \n", player.Number, player.Handicap, player.Scores)
		}
	}
	stats, err := t.GetStatistics()
	if err != nil {
		log.Fatal("Error getting tournament stats")
	}
	fmt.Println()
	fmt.Printf("Number of Players Counted in Results: %v\n", stats.Count)
	fmt.Printf("Lowest Score: %v\n", stats.Low)
	fmt.Printf("Highest Score: %v\n", stats.High)
	fmt.Printf("Average Score: %v\n", stats.Average)
	fmt.Printf("Standard Deviation: %v\n", stats.StdDev)
	fmt.Printf("Winner: Golfer #%v\n", t.winner.Number)
	fmt.Println()
	fmt.Println()
}
