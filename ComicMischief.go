package main

import (
	"fmt"
	"time"
)

func main() {

	//string variables
	var title string
	var writer string
	var artist string
	var publisher string

	//integer variables
	var year int32
	var pageNumber int32

	//float variables
	var grade float32

	//Creating our commic book
	title = "Mr.GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publisher Inc."
	year = 1997
	pageNumber = 14
	grade = 6.5

	//Printing the values
	fmt.Println("Title :", title, "Writer :", writer, "Artist :", artist, "Publisher", publisher)
	fmt.Println("Year of publication :", year, "Number of pages :", pageNumber, "Grade :", grade, "Age :", time.Now().Year()-int(year))

	//Changing values of our comic book
	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	year = 2013
	pageNumber = 160
	grade = 9.0

	//Printing the values
	fmt.Println("Title :", title, "Writer :", writer, "Artist :", artist, "Publisher", publisher)
	fmt.Println("Year of publication :", year, "Number of pages :", pageNumber, "Grade :", grade, "Age :", time.Now().Year()-int(year))

}
