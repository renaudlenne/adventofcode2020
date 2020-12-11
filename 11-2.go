package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Room struct {
	rows [][]byte
}

func (r Room) print() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("")
	for _, row := range r.rows {
		fmt.Println(string(row))
	}
}

func (r Room) countOccupiedSeats() int {
	count := 0
	for _, row := range r.rows {
		for _, seat := range row {
			if seat == '#' {
				count++
			}
		}
	}
	return count
}

func (r Room) updateState() (bool, Room) {
	stateChanged := false
	var newRows [][]byte
	for i, row := range r.rows {
		var newRow []byte
		for j, seat := range row {
			neighborhood := r.neighbors(i, j)
			occupiedNeighbors := countOccupied(neighborhood)
			if seat == 'L' && occupiedNeighbors == 0 {
				stateChanged = true
				newRow = append(newRow, '#')
			} else if seat == '#' && occupiedNeighbors >= 5 {
				stateChanged = true
				newRow = append(newRow, 'L')
			} else {
				newRow = append(newRow, seat)
			}
		}
		newRows = append(newRows, newRow)
	}
	if stateChanged {
		return stateChanged, Room{rows: newRows}
	}
	return stateChanged, r
}

func (r Room) neighbors(row int, col int) []byte {
	var res []byte
	rowLength := len(r.rows[0])
	colLength := len(r.rows)

	// line left
	for i := col - 1; i >= 0; i-- {
		seat := r.rows[row][i]
		if seat == '#' {
			res = append(res, seat)
			break
		}
		if seat == 'L' {
			res = append(res, seat)
			break
		}
	}

	// line right
	for i := col + 1; i < rowLength; i++ {
		seat := r.rows[row][i]
		if seat == '#' {
			res = append(res, seat)
			break
		}
		if seat == 'L' {
			res = append(res, seat)
			break
		}
	}

	// col top
	for i := row - 1; i >= 0; i-- {
		seat := r.rows[i][col]
		if seat == '#' {
			res = append(res, seat)
			break
		}
		if seat == 'L' {
			res = append(res, seat)
			break
		}
	}

	// col bottom
	for i := row + 1; i < colLength; i++ {
		seat := r.rows[i][col]
		if seat == '#' {
			res = append(res, seat)
			break
		}
		if seat == 'L' {
			res = append(res, seat)
			break
		}
	}

	// diag top left
	for i := 1; row-i >= 0 && col-i >= 0; i++ {
		seat := r.rows[row-i][col-i]
		if seat == '#' {
			res = append(res, seat)
			break
		}
		if seat == 'L' {
			res = append(res, seat)
			break
		}
	}

	// diag top right
	for i := 1; row-i >= 0 && col+i < rowLength; i++ {
		seat := r.rows[row-i][col+i]
		if seat == '#' {
			res = append(res, seat)
			break
		}
		if seat == 'L' {
			res = append(res, seat)
			break
		}
	}

	// diag top left
	for i := 1; row+i < colLength && col-i >= 0; i++ {
		seat := r.rows[row+i][col-i]
		if seat == '#' {
			res = append(res, seat)
			break
		}
		if seat == 'L' {
			res = append(res, seat)
			break
		}
	}

	// diag bottom right
	for i := 1; row+i < colLength && col+i < rowLength; i++ {
		seat := r.rows[row+i][col+i]
		if seat == '#' {
			res = append(res, seat)
			break
		}
		if seat == 'L' {
			res = append(res, seat)
			break
		}
	}

	return res
}

func countOccupied(seats []byte) int {
	res := 0
	for _, seat := range seats {
		if seat == '#' {
			res++
		}
	}
	return res
}

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rows [][]byte
	for scanner.Scan() {
		rows = append(rows, []byte(scanner.Text()))
	}
	room := Room{rows: rows}
	var changed bool
	changed, room = room.updateState()
	for changed {
		changed, room = room.updateState()
	}
	fmt.Printf("%d\n", room.countOccupiedSeats())
}
