package main

import (
	"fmt"
	"strconv"
)

// TODO: put on github/bitbucket
// TODO: handle invalid input
// TODO: make case-insensitive
// TODO: write tests and refactor

// TODO: seperate plateau and rover into Golang packages?
// TODO: right some todos of some things that I think should be done

type Plateau struct {
	width, height int
}

type Vector2 struct {
	x, y int
}

type Rover struct {
	x, y int
	dir string
	name string
	instructions string
	plateau *Plateau
}

var (
	numRovers int = 2
	plateau Plateau
	moves = map[string]Vector2 {
		"N": Vector2 {
			0, 1,
		},
		"E": Vector2 {
			1, 0,
		},
		"S": Vector2 {
			0, -1,
		},
		"W": Vector2 {
			-1, 0,
		},
	}
)

func main() {
	plateau.getSize()

	rovers := make([]Rover, numRovers)

	for i := 0; i < len(rovers); i++ {
		rovers[i].plateau = &plateau
		rovers[i].name = strconv.Itoa(i + 1);
		rovers[i].getRoverLanding()
		rovers[i].getInstructions()
		rovers[i].activate()
	}
}

func (plateau *Plateau) getSize() {
	var w, h int

	fmt.Println("Plateau: ")
	fmt.Scan(&w, &h)
	  
	plateau.width = w;
	plateau.height = h;
}

func (rover *Rover) getRoverLanding() {
	var (
		x, y int
		dir string
	)

	fmt.Printf("Rover%s Landing: ", rover.name)
	fmt.Scan(&x, &y, &dir)

	if _, ok := moves[dir]; ok {
		// fmt.Println("nice moves keep it up proud of you")
	}

	rover.x = x
	rover.y = y
	rover.dir = dir
}

func (rover *Rover) getInstructions() {
	var instructions string

	fmt.Printf("Rover%s Instructions: ", rover.name)
	fmt.Scan(&instructions)

	rover.instructions = instructions
}

func (rover *Rover) move() {
	var x, y int

	x = rover.x + moves[rover.dir].x
	y = rover.y + moves[rover.dir].y

	if (x >= 0 && x <= rover.plateau.width) {
		rover.x = x
	}

	if (y >= 0 && y <= rover.plateau.height) {
		rover.y = y
	}
}

func (rover *Rover) rotate(c string) {
	if c == "R" {
		switch rover.dir {
		case "N":
			rover.dir = "E"
		case "E":
			rover.dir = "S"
		case "S":
			rover.dir = "W"
		case "W":
			rover.dir = "N"
		}
	} else if c == "L" {
		switch rover.dir {
		case "N":
			rover.dir = "W"
		case "E":
			rover.dir = "N"
		case "S":
			rover.dir = "E"
		case "W":
			rover.dir = "S"
		}
	}
}

func (rover *Rover) activate() {
	for _, char := range rover.instructions {
		instruction := string(char);

		if instruction == "M" {
			rover.move()
		} else if instruction == "L" || instruction == "R" {
			rover.rotate(instruction)
		}
	}

	rover.print()
}

func (rover *Rover) print() {
	fmt.Printf("Rover%s: %d %d %s\n", rover.name, rover.x, rover.y, rover.dir)
}