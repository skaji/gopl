package main

func main() {
	var x []int
	go func() {
		x = make([]int, 10)
	}()
	go func() {
		x = make([]int, 1000)
	}()
	x[999] = 1
}
