package main

import (
	"fmt"
	"os"
)

// The number of cell by one side.
const l = 19

// State is whether the cell is black, white or empty.
type State int

const (
	Empty State = iota
	Black
	White
)

// Field is whole game board.
type Field [l][l]State

var verbose = false

func main() {
	if len(os.Args) >= 2 && os.Args[1] == "-v" {
		verbose = true
	}

	field := Field{}

	myState := Black
	opponentState := White

	for {
		printField(field)

		var opponentY, opponentX int
		if n, err := fmt.Scanf("%d %d", &opponentY, &opponentX); n != 2 || err != nil {
			fmt.Printf("n: %s, error: %s", n, err)
			break
		}

		if opponentY == -1 && opponentX == -1 {
			myState = White
			opponentState = Black
		} else {
			field[opponentY][opponentX] = opponentState
		}

		printField(field)

		myY, myX, _ := calcNextPlacing(field)

		fmt.Printf("%d %d\n", myY, myX)
		field[myY][myX] = myState
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
	if !verbose {
		return
	}

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
