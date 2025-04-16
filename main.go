package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Coord represents a coordinate (x, y) on the 2D grid.
// We use int64 to support extremely large positions.
type Coord struct {
	X, Y int64
}

// neighbors defines the 8 directions around a cell.
var neighbors = []Coord{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func main() {
	// Read the initial live cells from input.txt
	liveCells := readInput("input.txt")

	// Number of generations to simulate
	generations := 10

	// Simulate and print each generation
	for i := 0; i < generations; i++ {
		fmt.Printf("Generation %d:\n", i+1)
		printGrid(liveCells, -5, 5, -5, 5) // Print a fixed area of the grid
		fmt.Println()
		liveCells = step(liveCells) // Advance one generation
	}

	// Save the final generation to output.txt
	writeOutput("output.txt", liveCells)
}

// readInput reads live cells from a file formatted as Life 1.06
func readInput(filename string) map[Coord]bool {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cells := make(map[Coord]bool)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			// Skip comments and blank lines
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			// Skip malformed lines
			continue
		}

		// Parse coordinates from strings to int64
		x, _ := strconv.ParseInt(parts[0], 10, 64)
		y, _ := strconv.ParseInt(parts[1], 10, 64)

		// Add the coordinate to the map of live cells
		cells[Coord{x, y}] = true
	}

	return cells
}

// writeOutput writes the final live cells to a file in Life 1.06 format
func writeOutput(filename string, cells map[Coord]bool) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintln(writer, "#Life 1.06") // Write header

	// Write each live cell's coordinates
	for cell := range cells {
		fmt.Fprintf(writer, "%d %d\n", cell.X, cell.Y)
	}

	writer.Flush()
}

// step advances the simulation by one generation using Game of Life rules
func step(current map[Coord]bool) map[Coord]bool {
	// Count how many times each cell is a neighbor to a live cell
	neighborCount := make(map[Coord]int)

	// For each live cell, increment neighbor counts
	for cell := range current {
		for _, offset := range neighbors {
			neighbor := Coord{cell.X + offset.X, cell.Y + offset.Y}
			neighborCount[neighbor]++
		}
	}

	// Create the next generation
	next := make(map[Coord]bool)

	// Apply the rules of Game of Life
	for cell, count := range neighborCount {
		// Rule 1: A dead cell with exactly 3 neighbors becomes alive
		// Rule 2: A live cell with 2 or 3 neighbors stays alive
		if count == 3 || (count == 2 && current[cell]) {
			next[cell] = true
		}
		// Otherwise, the cell stays dead or dies
	}

	return next
}

// printGrid displays a rectangular region of the grid in the terminal
func printGrid(cells map[Coord]bool, minX, maxX, minY, maxY int64) {
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if cells[Coord{x, y}] {
				fmt.Print("■ ") // Live cell
			} else {
				fmt.Print(". ") // Dead cell
			}
		}
		fmt.Println()
	}
}
