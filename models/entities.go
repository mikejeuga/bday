package models

import "time"

type Person struct {
	FName   string  `json:"firstName"`
	LName   string  `json:"lastName"`
	Dob     DoB     `json:"date_of_birth"`
	Country Country `json:"country"`
	Email   string  `json:"email"`
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
