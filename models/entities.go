package models

import "time"

type Person struct {
	LName string `json:"lastName"`
	FName string `json:"firstName"`
	Dob   DoB    `json:"date_of_birth"`
}

func NewPerson(FName string, LName string, dob DoB) Person {
	return Person{FName: FName, LName: LName, Dob: dob}
}

type DoB struct {
	Year  int        `json:"year"`
	Month time.Month `json:"month"`
	Day   int        `json:"day"`
}

type Country string
