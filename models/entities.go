package models

import "time"

type Person struct {
	FName   string
	LName   string
	Dob     DoB
	Country Country
	Email   string
}

func NewPerson(FName string, LName string, dob DoB) Person {
	return Person{FName: FName, LName: LName, Dob: dob}
}

type DoB struct {
	Day   int
	Month time.Month
	Year  int
}

type Country string
