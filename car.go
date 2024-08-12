package main

import (
	"fmt"
	"strings"
)

var template = []string{
	"  ______",
	"  /|_||_\\`.__",
	" (   _    _ _\\",
	"=`-(_)--(_)-'",
}

type Car struct {
	sb    *strings.Builder
	name  string
	speed int
	pos   int
	laps  int
}

func NewCar(sb *strings.Builder, name string, speed int) *Car {
	return &Car{
		sb:    sb,
		name:  name,
		speed: speed,
	}
}

func (c *Car) Move() {
	if c.pos >= bounds {
		c.pos = 0
		c.laps++
	} else {
		c.pos = min(c.pos+c.speed, bounds)
	}

	c.writeCar()
}

func (c *Car) writeCar() {
	leftPad := strings.Repeat(" ", c.pos)
	rightPad := strings.Repeat(" ", bounds-c.pos)

	for _, s := range template {
		c.sb.WriteString(leftPad)
		c.sb.WriteString(s)
		c.sb.WriteString(rightPad)
		c.sb.WriteString("\n")
	}
}

func (c Car) Stat() {
	fmt.Fprintf(c.sb, "%s	| Laps: %d	| Speed: %d\n",
		c.name,
		c.laps,
		c.speed,
	)
}
