package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Learning Slices")
	var slc = []string{"Welcome", "to", "Learning", "Slices"}
	fmt.Printf("slc is of data type %T \n", slc)

	slc = append(slc, "in", "Golang")
	fmt.Println(slc)

	// slc = append(slc[1:3]) // last place is non-inclusive
	// fmt.Println(" slc after appending slc[1:3] : ", slc)

	slc = append(slc[:3]) // last place is non-inclusive
	fmt.Println(" slc after appending slc[1:3] : ", slc)

	highScores := make([]int, 4)
	fmt.Println(highScores)
	highScores[0] = 100
	highScores[1] = 500
	highScores[2] = 200
	highScores[3] = 300
	fmt.Println(highScores)

	// highScores[4] = 600
	highScores = append(highScores, 600)

	fmt.Println(sort.IntsAreSorted(highScores))

	fmt.Println(highScores)

	sort.Ints(highScores)

	// slices.Sort(highScores)

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	// remove a element from slice based on index
	var courses = []string{"Golang", "Python", "Javascript", "Java"}
	fmt.Println(courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
