package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

type Pivot struct {
	// {x, y}
	Destination []int
	CurrentLocation []int
	//  left right up down
	// [   1,    1, 1,  1]
	CanMove []int
	StepCount int
	// TODO implement foresight
	// This will tell us whether to traverse forward or reverse
	StartedLineLeft bool
	StartedLineRight bool
	Grid [][]int
}

func (p *Pivot) Init(grid [][]int) {
	p.Destination = make([]int, 2)
	p.Destination[0] = len(grid[0]) - 1
	p.Destination[1] = len(grid) - 1
	p.CurrentLocation = []int{0, 0}
	p.CanMove = []int{1, 1, 1, 1}
	p.StepCount = 0
	p.Grid = grid
}

func (p *Pivot) MoveLeft() {
	fmt.Println("moving left")
	p.CurrentLocation[1] = p.CurrentLocation[1] - 1
	p.StepCount++
	p.FigureOutPossiblePaths()
}

func (p *Pivot) MoveRight() {
	fmt.Println("moving right")
	p.CurrentLocation[1] = p.CurrentLocation[1] + 1
	p.StepCount++
	p.FigureOutPossiblePaths()
}

func (p *Pivot) MoveUp() {
	fmt.Println("moving up")
	p.CurrentLocation[0] = p.CurrentLocation[0] - 1
	p.StepCount++
	p.FigureOutPossiblePaths()
}

func (p *Pivot) MoveDown() {
	fmt.Println("moving down")
	p.CurrentLocation[0] = p.CurrentLocation[0] + 1
	if p.CurrentLocation[1] == 0 {
		p.StartedLineLeft = true
	} else {
		p.StartedLineRight = true
	}
	p.StepCount++
	p.FigureOutPossiblePaths()
}

func (p *Pivot) FigureOutPossiblePaths() {
	// vertical
	i := p.CurrentLocation[0]
	// horizontal
	j := p.CurrentLocation[1]
	// Boundaries
	// Can't move left
	if j == 0 {
		p.CanMove[0] = 0
	} else {
		p.CanMove[0] = 1
		// are there obstacles left?
		if p.Grid[i][j - 1] == 1{
			p.CanMove[0] = 0
		}
	}
	// Can't move right
	if j == len(p.Grid[0]) - 1 {
		p.CanMove[1] = 0
	} else {
		p.CanMove[1] = 1
		// are there obstacles right?
		if p.Grid[i][j + 1] == 1{
			p.CanMove[1] = 0
		}
	}
	// Can't move up
	if i == 0 {
		p.CanMove[2] = 0
	} else {
		p.CanMove[2] = 1
		// are there obstacles up?
		if p.Grid[i - 1][j] == 1{
			p.CanMove[2] = 0
		}
	}
	// Can't move down
	if i == len(p.Grid) - 1 {
		p.CanMove[3] = 0
	} else {
		p.CanMove[3] = 1
		// are there obstacles down?
		if p.Grid[i + 1][j] == 1{
			p.CanMove[3] = 0
		}
	}
}

func (p *Pivot) Traverse() {
	for {
		if p.CanMove[0] == 0 && p.CanMove[1] == 0 &&
			p.CanMove[2] == 0 && p.CanMove[3] == 0 {
			fmt.Printf("Reached a dead end :c \n")
			p.StepCount = -1
			break
		}
		// Can move only left
		if p.CanMove[0] == 1 && p.CanMove[1] == 0 &&
			p.CanMove[2] == 0 && p.CanMove[3] == 0 {
			p.MoveLeft()
			continue
		}
		// Can move only right
		if p.CanMove[0] == 0 && p.CanMove[1] == 1 &&
			p.CanMove[2] == 0 && p.CanMove[3] == 0 {
			p.MoveRight()
			continue
		}
		// Can move only up
		if p.CanMove[0] == 0 && p.CanMove[1] == 0 &&
			p.CanMove[2] == 1 && p.CanMove[3] == 0 {
			p.MoveUp()
			continue
		}
		// Can move only down
		if p.CanMove[0] == 0 && p.CanMove[1] == 0 &&
			p.CanMove[2] == 0 && p.CanMove[3] == 1 {
			p.MoveDown()
			continue
		}
		// Can move left or right
		// prefer to move right if p.StartedLineLeft
		// prefer to move left if p.StartedLineRight
		if p.CanMove[0] == 1 || p.CanMove[1] == 1 {
			if (p.StartedLineLeft) {
				p.MoveRight()
				continue
			} else if (p. StartedLineRight) {
				p.MoveLeft()
				continue
			}
		}
		// TODO Can move up or down
		if p.CanMove[2] == 1 || p.CanMove[3] == 1 {
			// For now we're just going down
			p.MoveDown()
			continue
		}
		if p.CurrentLocation[0] == p.Destination[0] &&
			p.CurrentLocation[1] == p.Destination[1] {
			fmt.Printf("Reached destination in %d steps :)\n", p.StepCount)
			break
		}
	}
}

func shortestPath(grid [][]int, k int) int {
	pivot := &Pivot{}
	pivot.Init(grid)
	pivot.Traverse()
    return pivot.StepCount
}
