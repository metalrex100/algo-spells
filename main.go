package main

import (
	"fmt"
)

const (
	squareSize = 25

	gridSymbol = "#  "
	dotSymbol  = ".  "
)

type spell struct {
	number    int
	condition func(x, y int) bool
}

func main() {
	spells := []spell{
		{number: 1, condition: func(x, y int) bool { return y > x }},
		{number: 2, condition: func(x, y int) bool { return x == y }},
		{number: 3, condition: func(x, y int) bool { return y == squareSize-(x-1) }},
		{number: 4, condition: func(x, y int) bool { return y-6 < squareSize-(x-1) }},
		{number: 5, condition: func(x, y int) bool { return (y-1)/2 == x-1 }},
		{number: 6, condition: func(x, y int) bool { return y <= 10 || x <= 10 }},
		{number: 7, condition: func(x, y int) bool { return y >= 15 && x >= 15 }},
		{number: 8, condition: func(x, y int) bool { return y <= 1 || x <= 1 }},
		{number: 9, condition: func(x, y int) bool { return y > x+10 || y < x-10 }},
		{number: 10, condition: func(x, y int) bool { return ((y+1)/2 <= x) && (x < y) }},
		{number: 13, condition: func(x, y int) bool { return y > squareSize-(x+4) && y < squareSize-(x-6) }},
		{number: 15, condition: func(x, y int) bool { return (y <= x+20 && y >= x-20) && (y > x+9 || y < x-9) }},
		{number: 18, condition: func(x, y int) bool { return (y <= 2 || x <= 2) && !(x == 1 && y == 1) }},
		{number: 19, condition: func(x, y int) bool { return (y <= 1 || x <= 1) || (y >= 25 || x >= 25) }},
		{number: 20, condition: func(x, y int) bool { return (x%2 == 0 && y%2 == 0) || (x%2 != 0 && y%2 != 0) }},
		{number: 23, condition: func(x, y int) bool { return x%3 == 1 && y%2 > 0 }},
		{number: 24, condition: func(x, y int) bool { return x == y || y == squareSize-(x-1) }},
		{number: 25, condition: func(x, y int) bool { return (x-1)%6==0 || (y-1)%6==0 }},
	}

	castSpells(spells)

	fmt.Printf("\nSpells mastered: %d", len(spells))
}

func castSpells(spells []spell) {
	for _, s := range spells {
		fmt.Printf("\nSpell #%d\n", s.number)
		castSpell(s)
		fmt.Println()
	}
}

func castSpell(s spell) {
	for x := 1; x <= squareSize; x++ {
		for y := 1; y <= squareSize; y++ {
			if s.condition(x, y) {
				fmt.Print(gridSymbol)
			} else {
				fmt.Print(dotSymbol)
			}
		}
		fmt.Println()
	}
}
