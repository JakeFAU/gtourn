package main

import "fmt"

type Player struct {
	Number   uint
	Handicap int
	Scores   []int
}

func (p *Player) IsValid() bool {
	if p.Handicap < 0 {
		return false
	}

	for _, score := range p.Scores {
		if score <= 0 {
			return false
		}
	}

	return true
}

type IllegalScoreError struct {
	handicap int
	scores   []int
}

func (e *IllegalScoreError) Error() string {
	return fmt.Sprintf("Either the handicap %v or scores %v are incorrect", e.handicap, e.scores)
}

func (p *Player) GrossScore() (int, error) {
	if !p.IsValid() {
		return -1, &IllegalScoreError{handicap: p.Handicap, scores: p.Scores}
	}
	score := 0
	for _, s := range p.Scores {
		score += s
	}
	return score, nil
}

func (p *Player) NetScore() (int, error) {
	gross, err := p.GrossScore()
	if err != nil {
		return -1, err
	}

	return gross - p.Handicap, nil

}

func (p *Player) UnderPar(par int) (bool, error) {
	net, err := p.NetScore()
	if err != nil {
		return false, err
	}
	return net < par, nil
}
