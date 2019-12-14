package main

import "fmt"

func main() {
	w := workout{
		name: "Stuff",
		typeOfWorkout: Cardio,
		weekday: Monday,
	}

	fmt.Println(w.name)
}

type Weekday int
type TypeofWorkout int

type workout struct {
	name string
	typeOfWorkout TypeofWorkout
	weekday Weekday
}

const (
	Monday    Weekday = 1
	Tuesday   Weekday = 2
	Wednesday Weekday = 3
	Thursday  Weekday = 4
	Friday    Weekday = 5
	Saturday  Weekday = 6
	Sunday    Weekday = 7
)

const (
	Cardio TypeofWorkout = 1
	Strength TypeofWorkout = 2
	Stretchy TypeofWorkout = 3
)