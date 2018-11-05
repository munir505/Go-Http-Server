package main

import "fmt"

const s int = 999

func do_stuff(){
	fmt.Println("Hello World")
	fmt.Println(!true)
	fmt.Println("Number: ", 1356)
	str := "This is what"
	fmt.Println(str)
	number, number2 := 12466, 5467865
	fmt.Println(number2)
	fmt.Println(number)

	fmt.Println(s)
	number += 100000
	fmt.Println(number)
}

func loops_test(){
	fmt.Println("For Loop Testing")
	for i := 0;i < 20;i++  {
		fmt.Println("Loop one: ", i)
		if i == 10 {
			fmt.Println("Loop is gonna break")
			break
		}
	}
	fmt.Println("Loop is now broken")
}

func switch_test(){
	fmt.Println("Switch Case Testing")
	number := "World"
	switch number {
		case "Hello":
			fmt.Println("One")
		case "World":
			fmt.Println("Two")
		case "Munir":
			fmt.Println("Three")
	}
}

func array_test(){
	fmt.Println("Array Testing")
	var arr [10]int
	for i:=0;i<10;i++ {
		arr[i] = i * i
	}
	fmt.Println(arr)
}

func slice_test(){
	s := make([]int, 3)
	fmt.Println("Empty Slice: ", s)
	s[0] = 1
	s[1] = 2
	s[2] = 3
	fmt.Println("Count Slice: ", s)
	s = append(s, 4)
	fmt.Println("Count Slice: ", s)
	s = append(s, 5, 6, 7, 8, 9, 10)
	fmt.Println("Count Slice: ", s)
	
	other := make([]int, len(s))
	copy(other, s)
	fmt.Println("Copy Slice: ", other)
	
	slice_one := s[3:7]
	fmt.Println("Slice One: ", slice_one)
}

func map_test(){
	fmt.Println("Map Testing")
	map_one := make(map[string]int)
	map_one["V1"] = 1
	map_one["V2"] = 2
	map_one["V3"] = 3
	map_one["V4"] = 4
	fmt.Println(map_one)
	
	//delete(map_one, "V1")
	fmt.Println(map_one)
	
	_, prs := map_one["V1"]
    fmt.Println("prs:", prs)
}

func multi_return()(int, int){
	return 25, 50
}

func range_test(){
	fmt.Println("Range Test")
}

func main() {
	do_stuff()
	fmt.Println()
	loops_test()
	switch_test()
	array_test()
	slice_test()
	map_test()
	one, two := multi_return()
	fmt.Println(one)
	fmt.Println(two)
	range_test()
}
