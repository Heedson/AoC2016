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

		switch heading {
		case 4:
			heading = 0
			fallthrough
		case 0:
			// North
			y += distance
		case 1:
			// East
			x += distance
		case 2:
			// South
			y -= distance
		case -1:
			heading = 3
			fallthrough
		case 3:
			// West
			x -= distance
		default:
			log.Fatalf("unsupported heading '%d'", heading)
		}
	}

	fmt.Printf("Finishing destination: (%d, %d)\n", x, y)
	fmt.Println("Distance away:", math.Abs(float64(x))+math.Abs(float64(y)))
}
