package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part_one(grid [][]string) {
	engine_parts := GenerateEngineParts(grid)
	symbol_map := GenerateSymbolMap(grid)
	valid_parts := []EnginePart{}

	for _, engine_part := range engine_parts {
		symbol_locs := engine_part.generate_symbol_locations(len(grid[0]), len(grid))
		for _, symbol_loc := range symbol_locs {
			if _, ok := symbol_map[symbol_loc]; ok {
				valid_parts = append(valid_parts, engine_part)
			}
		}
	}

	acc := 0
	for _, vp := range valid_parts {
		acc += vp.val
	}

	fmt.Printf("Answer to part 1 is: %d", acc)
}

func main() {
	grid := CreateGrid("./input.txt")
	part_one(grid)

}

func CreateGrid(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	chars := make([][]string, len(lines))
	for i, line := range lines {
		chars[i] = strings.Split(line, "")
	}

	return chars
}

func GenerateSymbolMap(grid [][]string) map[SymbolLocation]string {
	symbol_map := make(map[SymbolLocation]string)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			_, err := strconv.Atoi(grid[i][j])
			if grid[i][j] != "." && err != nil {
				symbol_location := SymbolLocation{
					x: i,
					y: j,
				}
				symbol_map[symbol_location] = grid[i][j]
			}
		}
	}

	return symbol_map
}

func GenerateEngineParts(grid [][]string) []EnginePart {
	m := regexp.MustCompile(`[0-9]+`)
	engine_parts := []EnginePart{}
	for i := 0; i < len(grid); i++ {
		joined := strings.Join(grid[i], "")
		num_indicies := m.FindAllStringIndex(joined, -1)
		nums := m.FindAllString(joined, -1)

		for j, num_index := range num_indicies {
			val, err := strconv.Atoi(nums[j])
			if err != nil {
				panic(err)
			}
			engine_parts = append(engine_parts,
				EnginePart{
					x_start: num_index[0],
					x_end:   num_index[1] - 1,
					y:       i,
					val:     val,
				},
			)
		}
	}

	return engine_parts
}

type EnginePart struct {
	x_start int
	x_end   int
	y       int
	val     int
}

type SymbolLocation struct {
	x int
	y int
}

func (e *EnginePart) generate_symbol_locations(x_lim, y_lim int) []SymbolLocation {
	symbol_locations := []SymbolLocation{}

	_x_start := e.x_start - 1
	_y_start := e.y - 1

	if _x_start < 0 {
		_x_start = 0
	}
	if _y_start < 0 {
		_y_start = 0
	}

	_x_end := e.x_end + 1
	_y_end := e.y + 1

	if _x_end > x_lim {
		_x_end = x_lim
	}
	if _y_end > y_lim {
		_y_end = 0
	}

	for i := _y_start; i <= _y_end; i++ {
		for j := _x_start; j <= _x_end; j++ {
			if !(i == e.y && j >= e.x_start && j <= e.x_end) {
				symbol_locations = append(symbol_locations,
					SymbolLocation{
						x: i,
						y: j,
					},
				)
			}
		}
	}

	return symbol_locations

}
