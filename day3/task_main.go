package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/heedson/AoC2016/common"
)

// Run runs the task for the day. Same as a standalone 'main.go'.
func Run() error {
	reader, err := os.Open("./day3/day3.txt")
	if err != nil {
		return err
	}

	input, err := common.LoadStringFromReader(reader)
	if err != nil {
		return err
	}

	triangles := strings.Split(input, "\n")

	totalRealTriangles := 0

	reg, err := regexp.Compile(`^\s*([0-9]+)\s*([0-9]+)\s*([0-9]+)\s*$`)
	if err != nil {
		return err
	}

	vertices := make([][]string, len(triangles))
	for i := 0; i < len(triangles); i += 3 {
		vertices[i] = reg.FindStringSubmatch(triangles[i])
		vertices[i] = vertices[i][1:]
		vertices[i+1] = reg.FindStringSubmatch(triangles[i+1])
		vertices[i+1] = vertices[i+1][1:]
		vertices[i+2] = reg.FindStringSubmatch(triangles[i+2])
		vertices[i+2] = vertices[i+2][1:]

		for j := 0; j < 3; j++ {
			values := make([]int, 3)
			for k := 0; k < 3; k++ {
				values[k], err = strconv.Atoi(vertices[i+k][j])
				if err != nil {
					return err
				}
			}

			totalRealTriangles += calculateTriangle(values)
		}
	}

	fmt.Println("Total entries:", len(triangles))
	fmt.Println("Total possible triangles:", totalRealTriangles)

	return nil
}

func calculateTriangle(values []int) int {
	if values[0]+values[1] <= values[2] {
		return 0
	}

	if values[1]+values[2] <= values[0] {
		return 0
	}

	if values[0]+values[2] <= values[1] {
		return 0
	}

	return 1
}
