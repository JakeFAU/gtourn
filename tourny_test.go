package main

import (
	"math"
	"testing"
)

func getPlayerOne() Player {
	p := Player{Number: 1,
		Handicap: 6,
		Scores:   []int{1, 3, 6, 2, 1, 4, 3, 2, 4}}
	return p
}

func getPlayerTwo() Player {
	p := Player{Number: 2,
		Handicap: 3,
		Scores:   []int{2, 2, 2, 2, 4, 4, 4, 2, 2}}
	return p
}

func getPlayerThree() Player {
	p := Player{Number: 3,
		Handicap: -1,
		Scores:   []int{3, 3, 4, 5, 3, 4, 2, 5, 2}}
	return p
}

func getPlayerFour() Player {
	p := Player{Number: 4,
		Handicap: 2,
		Scores:   []int{4, 5, 4, 3, 4, 1, 3, 5, 4}}
	return p
}

func getPlayers() []Player {
	a := getPlayerOne()
	b := getPlayerTwo()
	c := getPlayerThree()
	d := getPlayerFour()
	players := []Player{a, b, c, d}
	return players
}

func TestTournamentName(t *testing.T) {
	players := getPlayers()
	tourney := Tournament{Name: "Test Tourney", Players: players}
	if tourney.Name != "Test Tourney" {
		t.Errorf("Wanted Test Tourney, Got %v", tourney.Name)
	}
}

func TestTournamentPlayerCount(t *testing.T) {
	players := getPlayers()
	tourney := Tournament{Name: "Test Tourney", Players: players}
	statistics, err := tourney.GetStatistics()
	if err != nil {
		t.Error(err.Error())
	}
	if statistics.Count != 3 {
		t.Errorf("Wanted 3, Got %v", statistics.Count)
	}
}

func TestTournamentLow(t *testing.T) {
	players := getPlayers()
	tourney := Tournament{Name: "Test Tourney", Players: players}
	statistics, err := tourney.GetStatistics()
	if err != nil {
		t.Error(err.Error())
	}
	if statistics.Low != 20 {
		t.Errorf("Wanted 20, Got %v", statistics.Low)
	}
}

func TestTournamentHigh(t *testing.T) {
	players := getPlayers()
	tourney := Tournament{Name: "Test Tourney", Players: players}
	statistics, err := tourney.GetStatistics()
	if err != nil {
		t.Error(err.Error())
	}
	if statistics.High != 31 {
		t.Errorf("Wanted 31, Got %v", statistics.High)
	}
}

func TestTournamentAverage(t *testing.T) {
	players := getPlayers()
	tourney := Tournament{Name: "Test Tourney", Players: players}
	statistics, err := tourney.GetStatistics()
	if err != nil {
		t.Error(err.Error())
	}
	if statistics.Average != 24.0 {
		t.Errorf("Wanted 24.0, Got %v", statistics.Average)
	}
}

func TestTournamentStdDev(t *testing.T) {
	players := getPlayers()
	tourney := Tournament{Name: "Test Tourney", Players: players}
	statistics, err := tourney.GetStatistics()
	if err != nil {
		t.Error(err.Error())
	}
	if math.Abs(statistics.StdDev-4.97) > .01 {
		t.Errorf("Got %v, expected %v", statistics.StdDev, 4.97)
	}
}

func TestWinner(t *testing.T) {
	players := getPlayers()
	tourney := Tournament{Name: "Test Tourney", Players: players}
	winner, err := tourney.GetWinner()
	if err != nil {
		t.Error(err.Error())
	}
	if winner.Number != 1 {
		t.Errorf("Got %v, expected %v", winner.Number, 1)
	}
}
