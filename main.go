package main

import (
	"slices"
	"strings"
)

// Player — структура для хранения данных о футболисте
type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

// calculateRating рассчитывает рейтинг игрока по формуле
func (p *Player) calculateRating() {
	score := float64(p.Goals) + float64(p.Assists)/2.0
	if p.Misses == 0 {
		p.Rating = score
	} else {
		p.Rating = score / float64(p.Misses)
	}
}

// NewPlayer — конструктор для создания нового игрока с автоматическим расчётом рейтинга
func NewPlayer(name string, goals, misses, assists int) Player {
	p := Player{
		Name:    name,
		Goals:   goals,
		Misses:  misses,
		Assists: assists,
	}
	p.calculateRating()
	return p
}

// goalsSort сортирует игроков по убыванию количества голов.
// При совпадении голов — сортирует по имени в алфавитном порядке.
func goalsSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		if a.Goals != b.Goals {
			if a.Goals > b.Goals {
				return -1 // убывание
			}
			return 1
		}
		// Если голы равны, сортируем по имени (возрастание/алфавит)
		return strings.Compare(a.Name, b.Name)
	})
	return players
}

// ratingSort сортирует игроков по убыванию рейтинга.
// При совпадении рейтинга — сортирует по имени в алфавитном порядке.
func ratingSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		if a.Rating != b.Rating {
			if a.Rating > b.Rating {
				return -1 // убывание
			}
			return 1
		}
		// Если рейтинг равен, сортируем по имени (возрастание/алфавит)
		return strings.Compare(a.Name, b.Name)
	})
	return players
}

// gmSort сортирует игроков по убыванию отношения голов к промахам.
// При совпадении отношения — сортирует по имени в алфавитном порядке.
func gmSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		// Безопасно высчитываем отношение голов к промахам для игрока A
		var ratioA float64
		if a.Misses == 0 {
			ratioA = float64(a.Goals)
		} else {
			ratioA = float64(a.Goals) / float64(a.Misses)
		}

		// Безопасно высчитываем отношение голов к промахам для игрока B
		var ratioB float64
		if b.Misses == 0 {
			ratioB = float64(b.Goals)
		} else {
			ratioB = float64(b.Goals) / float64(b.Misses)
		}

		// Сравниваем полученные отношения (по убыванию)
		if ratioA != ratioB {
			if ratioA > ratioB {
				return -1 // убывание
			}
			return 1
		}

		// Если отношение одинаковое, сортируем по имени (алфавитный порядок)
		return strings.Compare(a.Name, b.Name)
	})
	return players
}
