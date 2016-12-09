package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := "L1, R3, R1, L5, L2, L5, R4, L2, R2, R2, L2, R1, L5, R3, L4, L1, L2, R3, R5, L2, R5, L1, R2, L5, R4, R2, R2, L1, L1, R1, L3, L1, R1, L3, R5, R3, R3, L4, R4, L2, L4, R1, R1, L193, R2, L1, R54, R1, L1, R71, L4, R3, R191, R3, R2, L4, R3, R2, L2, L4, L5, R4, R1, L2, L2, L3, L2, L1, R4, R1, R5, R3, L5, R3, R4, L2, R3, L1, L3, L3, L5, L1, L3, L3, L1, R3, L3, L2, R1, L3, L1, R5, R4, R3, R2, R3, L1, L2, R4, L3, R1, L1, L1, R5, R2, R4, R5, L1, L1, R1, L2, L4, R3, L1, L3, R5, R4, R3, R3, L2, R2, L1, R4, R2, L3, L4, L2, R2, R2, L4, R3, R5, L2, R2, R4, R5, L2, L3, L2, R5, L4, L2, R3, L5, R2, L1, R1, R3, R3, L5, L2, L2, R5"

	commands := strings.Split(input, ", ")

	x := 0
	y := 0
	heading := 0

	visitedCoords := map[string]bool{}

	for _, c := range commands {
		switch {
		case strings.Contains(c, "L"):
			heading--
		case strings.Contains(c, "R"):
			heading++
		default:
			log.Fatalf("unsupported direction '%s'", c[:1])
		}

		d := c[1:]
		distance, err := strconv.Atoi(d)
		if err != nil {
			log.Fatal(err)
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
			log.Fatalf("unsupported heading '%d'", heading)
		}

		found := false

		for i := 0; i < distance; i++ {
			x += xDirection
			y += yDirection

			key := fmt.Sprintf("%d,%d", x, y)
			found = visitedCoords[key]
			if found {
				fmt.Printf("First visited twice: (%d, %d)\n", x, y)
				fmt.Println("Distance away:", math.Abs(float64(x))+math.Abs(float64(y)))
				break
			}

			visitedCoords[key] = true
		}

		if found {
			break
		}
	}

	fmt.Printf("Finishing destination: (%d, %d)\n", x, y)
	fmt.Println("Distance away:", math.Abs(float64(x))+math.Abs(float64(y)))
}
