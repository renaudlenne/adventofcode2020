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
			} else if seat == '#' && occupiedNeighbors >= 4 {
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
	inFirstCol := col == 0
	inLastCol := col == len(r.rows[0])-1
	var res []byte
	if row > 0 {
		prevRow := r.rows[row-1]
		if !inFirstCol {
			res = append(res, prevRow[col-1])
		}
		res = append(res, prevRow[col])
		if !inLastCol {
			res = append(res, prevRow[col+1])
		}
	}
	if !inFirstCol {
		res = append(res, r.rows[row][col-1])
	}
	if !inLastCol {
		res = append(res, r.rows[row][col+1])
	}
	if row < len(r.rows)-1 {
		nextRow := r.rows[row+1]
		if !inFirstCol {
			res = append(res, nextRow[col-1])
		}
		res = append(res, nextRow[col])
		if !inLastCol {
			res = append(res, nextRow[col+1])
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
