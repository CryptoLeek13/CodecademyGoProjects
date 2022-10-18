package main

import "fmt"

// define the comic vars i.e publisher, writer, artist, title
func main() {
	var publisher, writer, artist, title string
	var year, pageNumber int
	var grade float32
	title, writer, artist, publisher, year, pageNumber, grade = "Mr. GoToSLeep", "Tracy Hatchet", "Jewel Tampson", "DizzyBooks Publishing Inc.", 1997, 14, 6.5
	// Print out the requested data
	fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "in year", year, "on page number", pageNumber, "it was graded", grade)
	// Alter the data vars
	title, writer, artist, publisher, year, pageNumber, grade = "Epic Vol. 1", "Ryan N. Shawn", "Phoebe Paperclips", "DizzyBooks Publishing Inc.", 2013, 160, 9.0
	// print out altered vars
	fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "in year", year, "on page number", pageNumber, "it was graded", grade)
}
