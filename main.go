package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/heedson/AoC2016/day1"
	"github.com/heedson/AoC2016/day2"
	"github.com/heedson/AoC2016/day3"
)

func main() {
	var day = flag.Int("day", 0, "Specify what days challenge to run.")
	flag.Parse()

	if *day == 0 {
		log.Fatal("A day needs to be selected to continue")
	}

	err := runDay(*day)
	if err != nil {
		log.Fatal(err)
	}
}

func runDay(day int) error {
	switch day {
	case 1:
		return day1.Run()
	case 2:
		return day2.Run()
	case 3:
		return day3.Run()
	default:
		return fmt.Errorf("Day '%d' is not implemented (yet)", day)
	}
}
