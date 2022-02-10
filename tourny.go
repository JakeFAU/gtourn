package main

import "math"

type Tournament struct {
	Name    string
	Players []Player
	stats   Statistics
	winner  Player
}

type Statistics struct {
	Count   int
	Low     int
	High    int
	Average float64
	StdDev  float64
}

func (t *Tournament) GetStatistics() (Statistics, error) {
	if t.stats == (Statistics{}) {
		t.calcStats()
	}
	return t.stats, nil
}

func (t *Tournament) GetWinner() (Player, error) {
	if t.stats == (Statistics{}) {
		t.calcStats()
	}
	return t.winner, nil
}

func (t *Tournament) calcStats() error {
	validPlayers := make([]Player, 0)
	for _, p := range t.Players {
		if p.IsValid() {
			validPlayers = append(validPlayers, p)
		}
	}

	var stats Statistics
	stats.Count = len(validPlayers)

	var _low, _high, _sum int
	_low = math.MaxInt

	//now stdev
	var sumSquaredScores, sumScores, scoresSquared, N float64

	N = float64(len(validPlayers))
	for _, p := range validPlayers {
		s, err := p.NetScore()
		if err != nil {
			return err
		}
		if s < _low {
			_low = s
			t.winner = p
		}
		if s > _high {
			_high = s
		}
		_sum += s
		stats.Low = _low
		stats.High = _high
		stats.Average = float64(_sum) / float64(stats.Count)
		sumSquaredScores += (math.Pow(float64(s), 2))
		sumScores += float64(s)
	}
	scoresSquared = math.Pow(sumScores, 2)
	stats.StdDev = math.Pow((sumSquaredScores-(scoresSquared/N))/N, .5)
	t.stats = stats
	return nil
}
