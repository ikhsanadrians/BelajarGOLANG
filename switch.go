package main

import "fmt"

func main() {

	x := 7
	res := 0

	switch {

	case x > 8:
		res++
	case x > 2 && x%2 == 0:
		res++
	case x > 3:
		res++

	}

	fmt.Println(res)

}
