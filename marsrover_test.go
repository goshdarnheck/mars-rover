package main

import "testing"

func TestRotate(t *testing.T) {
	var rover Rover

	cases := []struct {
		start, rotation, end string
	}{
		{"N", "R", "E"},
		{"N", "L", "W"},
		{"E", "R", "S"},
		{"E", "L", "N"},
		{"S", "R", "W"},
		{"S", "L", "E"},
		{"W", "R", "N"},
		{"W", "L", "S"},
	}

	for _, c := range cases {
		rover.dir = c.start
		rover.rotate(c.rotation)

		if rover.dir != c.end {
			t.Error()
		}
	}
}

func TestMove(t *testing.T) {
	var (
		rover Rover
		plateau Plateau
	)

	cases := []struct {
		width, height, sx, sy, ex, ey int
		dir  string
	}{
		{5, 5, 0, 0, 0, 1, "N"},
		{5, 5, 5, 5, 5, 5, "N"},
		{5, 5, 3, 3, 2, 3, "W"},
		{2, 2, 2, 0, 2, 0, "E"},
		{5, 5, 1, 1, 2, 1, "E"},
		{5, 5, 0, 0, 0, 0, "S"},
		{5, 5, 1, 1, 1, 0, "S"},
	}

	rover.plateau = &plateau

	for _, c := range cases {
		plateau.width = c.width
		plateau.height = c.height
		rover.dir = c.dir
		rover.x = c.sx
		rover.y = c.sy

		rover.move()

		if rover.x != c.ex || rover.y != c.ey {
			t.Error()
		}
	}
}

func TestActivate(t *testing.T) {
	var (
		rover Rover
		plateau Plateau
	)

	cases := []struct {
		width, height, sx, sy, ex, ey int
		instructions, sdir, edir string
	}{
		{5, 5, 0, 0, 0, 2, "MM", "N", "N"},
		{5, 5, 0, 0, 2, 0, "RMM", "N", "E"},
		{5, 5, 1, 2, 1, 3, "LMLMLMLMM", "N", "N"},
		{5, 5, 3, 3, 5, 1, "MMRMMRMRRM", "E", "E"},
	}

	rover.plateau = &plateau

	for _, c := range cases {
		plateau.width = c.width
		plateau.height = c.height
		rover.x = c.sx
		rover.y = c.sy
		rover.instructions = c.instructions
		rover.dir = c.sdir

		rover.activate()

		if rover.x != c.ex || rover.y != c.ey || rover.dir != c.edir {
			t.Error()
		}
	}
}