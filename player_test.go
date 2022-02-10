package main

import "testing"

func getValidPlayer() Player {
	p := Player{
		Number:   1,
		Handicap: 10,
		Scores:   []int{5, 6, 6, 7, 5, 6, 5, 4, 5},
	}
	return p
}

func getInvalidHandicap() Player {
	p := Player{
		Number:   1,
		Handicap: -1,
		Scores:   []int{5, 6, 6, 7, 5, 6, 5, 4, 5},
	}
	return p
}

func getInvalidScore() Player {
	p := Player{
		Number:   1,
		Handicap: 10,
		Scores:   []int{5, 6, 6, -7, 5, 6, 5, 4, 5},
	}
	return p
}

func TestPlayerHandicap(t *testing.T) {
	validPlayer := getValidPlayer()
	invalidPlayer := getInvalidHandicap()

	v, err := validPlayer.NetScore()
	if err != nil {
		t.Errorf(err.Error())
	}
	if v != 39 {
		t.Errorf("Wanted 39, got %v", v)
	}
	_, err = invalidPlayer.NetScore()
	if err == nil {
		t.Error("The player is invalid")
	}
}

func TestPlayerScores(t *testing.T) {
	validPlayer := getValidPlayer()
	invalidPlayer := getInvalidScore()

	v, err := validPlayer.NetScore()
	if err != nil {
		t.Errorf(err.Error())
	}
	if v != 39 {
		t.Errorf("Wanted 39, got %v", v)
	}
	_, err = invalidPlayer.NetScore()
	if err == nil {
		t.Error("The player is invalid")
	}
}

func TestGrossScore(t *testing.T) {
	validPlayer := getValidPlayer()
	invalidPlayer := getInvalidScore()

	v, err := validPlayer.GrossScore()
	if err != nil {
		t.Errorf(err.Error())
	}
	if v != 49 {
		t.Errorf("Wanted 49, got %v", v)
	}
	_, err = invalidPlayer.GrossScore()
	if err == nil {
		t.Error("The player is invalid")
	}
}

func TestNetScore(t *testing.T) {
	validPlayer := getValidPlayer()
	invalidPlayer := getInvalidScore()

	v, err := validPlayer.NetScore()
	if err != nil {
		t.Errorf(err.Error())
	}
	if v != 39 {
		t.Errorf("Wanted 39, got %v", v)
	}
	_, err = invalidPlayer.NetScore()
	if err == nil {
		t.Error("The player is invalid")
	}
}

func TestUnderPar(t *testing.T) {
	validPlayer := getValidPlayer()
	up30, err := validPlayer.UnderPar(30)
	if err != nil {
		t.Error(err.Error())
	}
	up40, err := validPlayer.UnderPar(40)
	if err != nil {
		t.Error(err.Error())
	}
	if up30 {
		t.Error("Player should not be under par30")
	}
	if !up40 {
		t.Error("Player should be under par40")
	}
}
