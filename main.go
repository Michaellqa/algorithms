package main

import "fmt"

func main() {
	sl := []int{0, 1, 2, 3, 4}
	mid := sl[1:4]
	fmt.Println(mid)
	mid[0] += 10
	fmt.Println(mid)
	fmt.Println(sl)
}
