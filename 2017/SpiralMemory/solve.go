package main

import (
	"fmt"
	"math"
	"errors"
	"os"
)


type Vector struct {
	x, y int
}

func (v *Vector) ManhattanDistance() (int){
	return Abs(v.x) + Abs(v.y)

}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	if n == 0 {
		return 0
	}
	return n
}


type Cell struct {
	// The ID of this cell
	id int

	// The address of this cell in the grid
	vector Vector

	// The value stored at this cell
	value int
}

func (c *Cell) adjacent() ([8]Vector){

	var adjacent [8]Vector

	var x = c.vector.x
	var y = c.vector.y

	adjacent[0] = Vector{x+1, y-1} //  1,-2
	adjacent[1] = Vector{x+0, y-1} //  0,-2
	adjacent[2] = Vector{x-1, y-1} // -1,-2

	adjacent[3] = Vector{x+1, y+0} //  1,-1
	adjacent[4] = Vector{x-1, y+0} // -1,-1

	adjacent[5] = Vector{x+1, y+1}
	adjacent[6] = Vector{x+0, y+1}
	adjacent[7] = Vector{x-1, y+1}



	return adjacent
}


func (c *Cell) setAddress() {
	if c.id == 1 {
		c.vector = Vector{0,0}
		c.value = 1
		return
	}

	distance := int(math.Ceil((math.Sqrt(float64(c.id)) - 1) /2))
	width :=  distance * 2 + 1

	bottomRightId := int(math.Pow(float64(width) - 2, 2) + 1)
	centerRightId := int(float64(bottomRightId) + math.Floor((float64(width)-3)/2))

	startDiff := c.id - bottomRightId
	centerDiff := c.id - centerRightId
	side := width - 1

	if startDiff < side {
		c.vector = Vector{
			distance,
			centerDiff,
			}
		return
	}

	if c.id - bottomRightId < side * 2 {
		c.vector = Vector{
			-(centerDiff - side),
			distance,
			}
		return
	}

	if c.id - bottomRightId < side * 3 {
		c.vector = Vector{
			-distance,
			-(centerDiff - 2 * side),
			}
		return
	}

	c.vector = Vector{
		centerDiff - 3 * side,
		-distance,
		}
}

type Grid struct {
	// the number of Cells in this Grid
	size, target int
	targetCell *Cell
	cells []*Cell
}

func (g *Grid) init (size int) {
	g.size = size
	solve := true
	for i := 1; i <= size; i++ {
		cell := new(Cell)
		cell.id = i
		cell.setAddress()
		if solve {
			for _, v := range cell.adjacent(){
				cell.value = cell.value + g.getCellByVector(v).value
				if cell.value > g.target {
					solve = false
					g.targetCell = cell
				}
			}
		}
		g.cells = append(g.cells, cell)
	}
}

func (g *Grid) getCellByVector(v Vector) (*Cell){
	for _, c := range g.cells{
		if c.vector == v{
			return c
		}
	}
	return &Cell{}
}


func (g * Grid) getCellById(id int) (*Cell, error){
	if id > len(g.cells){
		return nil, errors.New("requested ID is out of bounds")
	}

	return g.cells[id-1], nil
}

func (g *Grid) ManhattanDistance(id int) (int, error){
	cell, err := g.getCellById(id)
	if err != nil {
		return -1, err
	}

	return cell.vector.ManhattanDistance(), nil
}


func main() {
	address := 265149
	grid := Grid{address+1, address, nil, nil}
	grid.init(address)

	distance, err := grid.ManhattanDistance(address)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Printf("Manhattan Distance for address %v is %v.\n", address, distance)
	fmt.Printf("Next Largest value is %v\n", grid.targetCell.value)

}