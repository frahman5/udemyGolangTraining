// this is an exercise in exploring the sort package
package main

import (
	"fmt"
	"sort"
	"unicode"
)

type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Less(i, j int) bool {
	return unicode.ToLower(rune((p)[i][0])) < unicode.ToLower(rune((p)[j][0]))
}

func (p people) Swap(i, j int) {
	ival, jval := (p)[i], (p)[j]
	(p)[i] = jval
	(p)[j] = ival
}

// Sorts using the .Strings and .Ints functions, and also reverses
func way1() {
	fmt.Println("SORTING USING THE .STRINGS and .INTS FUNCTIONS!", "\n")
	// sort a people object using the sort package
	studyGroup := people{"Zeno",   "John", "Al", "Jenny"}
	fmt.Println("PEOPLE")
	fmt.Println("Start: ", studyGroup)
	sort.Strings(studyGroup)
	fmt.Println("Sorted: ", studyGroup)
	sort.Sort(sort.Reverse(studyGroup))
	fmt.Println("Reverse: ", studyGroup)

	//sort a []string using the sort package
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println("\n", "SLICE OF STRING")
	fmt.Println("Start: ", s)
	sort.Strings(s)
	fmt.Println("Sorted: ", s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println("Reverse: ", s)

	//sort a []int using the sort package
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println("\n", "SLICE OF INT")
	fmt.Println("Start: ", n)
	sort.Ints(n)
	fmt.Println("Sorted: ", n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println("Reverse: ", n)

}

// Sorts using the .SORT method
func way2() {
	fmt.Println("\n", "SORTING USING THE .SORT FUNCTIONS", "\n")
	// sort a people object
	studyGroup := people{"Zeno",   "John", "Al", "Jenny"}
	fmt.Println("PEOPLE")
	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)


	//sort a []string
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println("\n", "SLICE OF STRING")
	fmt.Println(s)
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)

	//sort a []int
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println("\n", "SLICE OF INT")
	fmt.Println(n)
	sort.Sort(sort.IntSlice(n))
	fmt.Println(n)
}

// Sorts using the .SLICESTABLE method
func way3() {
	fmt.Println("\n", "SORTING USING THE SLICESTABLE FUNCTIONS", "\n")
	// sort a people object
	studyGroup := people{"Zeno",   "John", "Al", "Jenny"}
	fmt.Println("PEOPLE")
	fmt.Println(studyGroup)
	sort.SliceStable(studyGroup, func(i, j int) bool {
		return unicode.ToLower(rune(studyGroup[i][0])) < unicode.ToLower(rune(studyGroup[j][0]))})
	fmt.Println(studyGroup)


	//sort a []string
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println("\n", "SLICE OF STRING")
	fmt.Println(s)
	sort.SliceStable(s, func(i, j int) bool {
		return unicode.ToLower(rune(s[i][0])) < unicode.ToLower(rune(s[j][0]))})
	fmt.Println(s)

	//sort a []int
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println("\n", "SLICE OF INT")
	fmt.Println(n)
	sort.SliceStable(n, func(i, j int) bool{ return n[i] < n[j] })
	fmt.Println(n)
}


// Sorts in reverse order
func main() {
	way1()
	way2()
	way3()

}
