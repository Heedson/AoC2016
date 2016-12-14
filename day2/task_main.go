package day2

import (
	"os"
	"strings"

	"github.com/heedson/AoC2016/common"
)

// Run runs the task for the day. Same as a standalone 'main.go'.
func Run() error {
	reader, err := os.Open("./day2/day2.txt")
	if err != nil {
		return err
	}

	input, err := common.LoadStringFromReader(reader)
	if err != nil {
		return err
	}

	keypad := [][]int{
		{0, 0, 1, 0, 0},
		{0, 2, 3, 4, 0},
		{5, 6, 7, 8, 9},
		{0, 10, 11, 12, 0}, // {0, A, B, C, 0},
		{0, 0, 13, 0, 0},   // {0, 0, D, 0, 0},
	}

	commands := strings.Split(input, "\n")

	row := 2
	col := 0

	for _, c := range commands {
		for i := range c {
			letter := string(c[i])
			switch letter {
			case "U":
				if row == 0 {
					continue
				}

				if keypad[row-1][col] == 0 {
					continue
				}

				row--
			case "R":
				if col == 4 {
					continue
				}

				if keypad[row][col+1] == 0 {
					continue
				}

				col++
			case "D":
				if row == 4 {
					continue
				}

				if keypad[row+1][col] == 0 {
					continue
				}

				row++
			case "L":
				if col == 0 {
					continue
				}

				if keypad[row][col-1] == 0 {
					continue
				}

				col--
			}
		}

		println("Next digit:", keypad[row][col])
	}

	return nil
}
