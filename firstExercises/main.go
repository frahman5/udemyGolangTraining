package main

import . "fmt"

func HelloWorld(a string) {
	Println(a)
}

func HelloUser() {
	var userName string
	Print("Please enter your name: ")
	Scan(&userName)
	Printf("Hello %v!\n", userName)
}

//func largeNumSmallNum() {
////	var smallNum int
////	var largeNum int
////	Print("Please enter an integer: ")
////	numItems, err := Scan(&smallNum)
////	Printf("Num items scanned: %v\n", numItems)
////	if err != nil {
////		for err.Error() == "expected integer" {
////			Print("Nono. Please enter an INTEGER: a positive, zero, or negative whole number: ")
////			numItems, err = Scan(&smallNum)
////			Printf("Num items scanned: %v\n", numItems)
////		}
////	}
////	Print("Fantastic! Now please enter a larger integer: ")
////	Scan(&largeNum)
////	for largeNum < smallNum {
////		Println("Oops! The second number must be bigger than the first. Please try again.")
////		Printf("Please enter an integer larger than %v: ", smallNum)
////		Scan(&largeNum)
////	}
////	Printf("The remainder of %v by %v is %v!\n", largeNum, smallNum, largeNum % smallNum)
////}

func allEvenNumsZeroToOneHun() {
	for i := 0; i <= 100; i++ {
		if i % 2 == 0 {
			Println(i)
		}
	}
}

func FizzBuzz() {
	for i := 0; i <= 100; i++ {
		switch  {
		case i%15 == 0:
			Println("FizzBuzz")
		case i%3 == 0:
			Println("Fizz")
		case i%5 == 0:
			Println("Buzz")
		default:
			Println(i)
		}
	}
}

func sumOfMultsOfThreeAndFiveBelowANum(n int)  {
	sum := 0
	for i := 1; i < n; i++ {
		if i%3 == 0 {
			sum += i
		} else if i%5 == 0 {
			sum += i
		} else {
			continue
		}
	}
	Printf("The sum of the multiples of 3 and 5 below %v is...%v!\n", n, sum)
}
func main() {
	sumOfMultsOfThreeAndFiveBelowANum(1000)
}