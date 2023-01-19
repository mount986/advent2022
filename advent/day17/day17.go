package day17

import (
	"bufio"
	"fmt"
)


func RunSim(scanner *bufio.Scanner, cycles int) Chamber {	
	chamber := NewChamber(7)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	jets := scanner.Text()
	jetIndex := 0
	shapeIndex := 0

	shapes := []string{
		"hLine",
		"cross",
		"angle",
		"vLine",
		"square",	
	}

	for i := 0; i < cycles; i++ {
		shape := NewShape(shapes[shapeIndex%len(shapes)], chamber.Height)

		for !shape.Settled {
			shape.Move(chamber, jets[jetIndex%len(jets)])
			jetIndex++
		}

		chamber.SettlePiece(shape)

		shapeIndex++
	}

	return chamber
}