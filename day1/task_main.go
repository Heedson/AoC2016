package day1

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/heedson/AoC2016/common"
)

// Run runs the task for the day. Same as a standalone 'main.go'.
func Run() error {
	reader, err := os.Open("./day1/day1.txt")
	if err != nil {
		return err
	}

	input, err := common.LoadStringFromReader(reader)
	if err != nil {
		return err
	}

	commands := strings.Split(input, ", ")

	x := 0
	y := 0
	heading := 0

	visitedCoords := map[string]struct{}{}

	for _, c := range commands {
		switch {
		case strings.Contains(c, "L"):
			heading--
		case strings.Contains(c, "R"):
			heading++
		default:
			return fmt.Errorf("unsupported direction '%s'", c[:1])
		}

		d := c[1:]
		distance, err := strconv.Atoi(d)
		if err != nil {
			return err
		}

		xDirection := 0
		yDirection := 0

		switch heading {
		case 4:
			heading = 0
			fallthrough
		case 0:
			// North
			yDirection = 1
			xDirection = 0
		case 1:
			// East
			xDirection = 1
			yDirection = 0
		case 2:
			// South
			yDirection = -1
			xDirection = 0
		case -1:
			heading = 3
			fallthrough
		case 3:
			// West
			xDirection = -1
			yDirection = 0
		default:
			return fmt.Errorf("unsupported heading '%d'", heading)
		}

		found := false

		for i := 0; i < distance; i++ {
			x += xDirection
			y += yDirection

			key := fmt.Sprintf("%d,%d", x, y)
			_, found = visitedCoords[key]
			if found {
				fmt.Printf("First visited twice: (%d, %d)\n", x, y)
				fmt.Println("Distance away:", math.Abs(float64(x))+math.Abs(float64(y)))
				return nil
			}

			visitedCoords[key] = struct{}{}
		}
	}

	return errors.New("did not cross any location twice")
}
