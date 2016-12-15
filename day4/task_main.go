package day4

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/heedson/AoC2016/common"
)

type pair struct {
	Letter string
	Count  int
}

type pairList []pair

func (p pairList) Len() int { return len(p) }

// Less is inversed to sort the list (1->) then (a->).
func (p pairList) Less(i, j int) bool {
	return (p[i].Count > p[j].Count || (p[i].Count == p[j].Count && p[i].Letter < p[j].Letter))
}
func (p pairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// Run runs the task for the day. Same as a standalone 'main.go'.
func Run() error {
	reader, err := os.Open("./day4/day4.txt")
	if err != nil {
		return err
	}

	input, err := common.LoadStringFromReader(reader)
	if err != nil {
		return err
	}

	reg, err := regexp.Compile(`^([a-z]+)([0-9]+)\[([a-z]{5})\]$`)
	if err != nil {
		return err
	}

	rooms := strings.Split(input, "\n")

	sumOfIDs := 0

	for _, r := range rooms {
		singleRoom := strings.Replace(r, "-", "", -1)
		roomData := reg.FindStringSubmatch(singleRoom)
		roomData = roomData[1:]

		foundBefore := map[string]struct{}{}
		pairs := pairList{}

		for _, c := range roomData[0] {
			if _, found := foundBefore[string(c)]; found {
				continue
			}

			foundBefore[string(c)] = struct{}{}
			p := pair{
				Letter: string(c),
				Count:  strings.Count(roomData[0], string(c)),
			}

			pairs = append(pairs, p)
		}

		sort.Sort(pairs)

		pairs = pairs[:5]
		var key string
		for _, p := range pairs {
			key += p.Letter
		}

		if key == roomData[2] {
			id, err := strconv.Atoi(roomData[1])
			if err != nil {
				return err
			}

			sumOfIDs += id
		}
	}

	fmt.Println("Sum of real IDs:", sumOfIDs)

	return nil
}
