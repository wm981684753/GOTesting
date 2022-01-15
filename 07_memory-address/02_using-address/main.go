package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	var meters2 float64
	fmt.Print("Enter meters2 swam: ")
	fmt.Scan(&meters2)
	yards := meters * meters2 * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")
}
