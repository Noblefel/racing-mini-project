package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"
)

const bounds = 100

func main() {
	var sb strings.Builder

	car1 := NewCar(&sb, "Slow Car", 1)
	car2 := NewCar(&sb, "Average Car", 3)
	car3 := NewCar(&sb, "Fast Car", 6)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	go func() {
		fmt.Scanln()
		sig <- os.Interrupt
	}()

	go func() {
		<-sig
		fmt.Print("\033[J")
		fmt.Println("Race is finished")
		os.Exit(1)
	}()

	fmt.Print("\033[J\n")
	for {
		sb.WriteString(strings.Repeat("-", bounds+10))
		car1.Stat()
		car1.Move()
		sb.WriteString(strings.Repeat("-", bounds+10))
		car2.Stat()
		car2.Move()
		sb.WriteString(strings.Repeat("-", bounds+10))
		car3.Stat()
		car3.Move()
		sb.WriteString(strings.Repeat("-", bounds+10))

		sb.WriteString("\npress enter to end the race\n")
		sb.WriteString("\033[20A")
		fmt.Print(sb.String())
		sb.Reset()
		time.Sleep(40 * time.Millisecond)
	}
}
