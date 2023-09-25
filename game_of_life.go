package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Width  = 40
	Height = 20
	Alive  = '■'
	Dead   = ' '
)

func main() {
	rand.Seed(time.Now().UnixNano())
	grid := initializeGrid()
	for {
		printGrid(grid)
		grid = nextGeneration(grid)
		time.Sleep(100 * time.Millisecond)
	}
}

func initializeGrid() [][]bool {
	grid := make([][]bool, Height)
	for i := range grid {
		grid[i] = make([]bool, Width)
		for j := range grid[i] {
			grid[i][j] = rand.Intn(2) == 1
		}
	}
	return grid
}

func printGrid(grid [][]bool) {
	for _, row := range grid {
		for _, cell := range row {
			if cell {
				fmt.Print(string(Alive))
			} else {
				fmt.Print(string(Dead))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func nextGeneration(grid [][]bool) [][]bool {
	newGrid := make([][]bool, Height)
	for i := range newGrid {
		newGrid[i] = make([]bool, Width)
	}

	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			neighbors := countNeighbors(grid, i, j)
			newGrid[i][j] = (grid[i][j] && neighbors == 2) || neighbors == 3
		}
	}

	return newGrid
}

func countNeighbors(grid [][]bool, x, y int) int {
	neighbors := 0
	deltas := []struct{ dx, dy int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, /*{0, 0},*/ {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, delta := range deltas {
		nx, ny := x+delta.dx, y+delta.dy
		if nx >= 0 && nx < Height && ny >= 0 && ny < Width && grid[nx][ny] {
			neighbors++
		}
	}

	return neighbors
}
