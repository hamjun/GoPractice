package main

import (
	//"bufio"
	"fmt"
	//"os"
	//"strconv"
)

type Point struct {
	x int32
	y int32
}

func add(x, y int) (z, z1 int) {
	defer fmt.Println("hello")
	z = x + y
	z1 = x - y
	return
}

func main() {
	var p1 Point = Point{1, 2}
	fmt.Println(p1)

	x := func(x int) int {
		return x * -1
	}(8)
	fmt.Println(x)
	fmt.Println(add(1, 2))

	var arr [4]int
	arr1 := []int{4, 5, 6}
	arr1 = append(arr1, 1)

	for _, element := range arr1 {
		fmt.Printf("%d\t", element)
	}
	fmt.Println(arr, arr1)

	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}
	delete(mp, "apple")

	//check if value exists
	val, ok := mp["pear"]
	fmt.Println(val, ok)
	//mp1 = make(map[string]int)

	fmt.Println(mp)
	/*
		var num1 float64 = 8.3
		var num2 int = 4
		ans := num1 / float64(num2)
		fmt.Printf("%f", ans)


			scanner := bufio.NewScanner(os.Stdin)
			fmt.Printf("Type the year you were born: ")
			scanner.Scan()
			input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			fmt.Printf("You will be %d years old at the end of 2022", 2022-input)

			var name string
			name = "Hello Jun"

			var number = 260
			number1 := 6

			fmt.Println("Hello World!", name, number)
			fmt.Printf("%T", number1)

	*/
}
