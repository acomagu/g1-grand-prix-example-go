package main

import (
	"fmt"
)

// The number of cell by one side.
const l = 19

// Color is whether the cell is black, white or empty.
type Color int

const (
	Empty Color = iota
	Black
	White
)

// Field is whole game board.
type Field [l][l]Color

func main() {
	field := Field{}

	myColor := Black
	opponentColor := White

	for {
		printField(field)

		var opponentY, opponentX int
		if n, err := fmt.Scanf("%d %d", &opponentY, &opponentX); n != 2 || err != nil {
			fmt.Printf("n: %s, error: %s", n, err)
			break
		}

		if opponentY == -1 && opponentX == -1 {
			myColor = White
			opponentColor = Black
		} else {
			field[opponentY][opponentX] = opponentColor
		}

		printField(field)

		myY, myX, _ := calcNextPlacing(field)
		field[myY][myX] = myColor
	}
}

func calcNextPlacing(field Field) (int, int, error) {
	for y := 0; y < l; y++ {
		for x := 0; x < l; x++ {
			if field[y][x] == Empty {
				return y, x, nil
			}
		}
	}
	return 0, 0, fmt.Errorf("no placable cells")
}

func printField(field Field) {
	fmt.Print("  ")
	for x := 0; x < l; x++ {
		fmt.Printf("%2d", x)
	}
	fmt.Println()

	for y := 0; y < l; y++ {
		fmt.Printf("%2d", y)
		for x := 0; x < l; x++ {
			switch field[y][x] {
			case Empty:
				fmt.Print("  ")
			case Black:
				fmt.Print(" x")
			case White:
				fmt.Print(" o")
			}
		}
		fmt.Printf("%2d\n", y)
	}

	fmt.Print("  ")
	for x := 0; x < l; x++ {
		fmt.Printf("%2d", x)
	}
	fmt.Println()
}
