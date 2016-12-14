package main

import (
	"flag"
	"log"

	"github.com/heedson/AoC2016/day1"
	"github.com/heedson/AoC2016/day2"
)

func main() {
	var day = flag.Int("day", 0, "Specify what days challenge to run.")
	flag.Parse()

	if *day == 0 {
		log.Fatal("A day needs to be selected to continue")
	}

	switch *day {
	case 1:
		err := day1.Run()
		if err != nil {
			log.Fatal(err)
		}
	case 2:
		err := day2.Run()
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("Day '%d' is not implemented (yet)", *day)
	}

}
