package concurrency

import "fmt"

func sum(s []int, c chan int) {
	sum := 0

	for _, v := range s {
		sum += v
	}

	c <- sum
}

// ShowChannel
func ShowChannel() {

	println("======================== Concurrency with channels =====================")

	sl := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int)

	go sum(sl[:len(sl)/2], c)
	go sum(sl[len(sl)/2:], c)

	x, y := <-c, <-c

	fmt.Println("Sum: ", (x + y))

	c2 := make(chan int, 5)
	go fibonacci(cap(c2), c2)
	for i := range c2 {
		fmt.Println(i)
	}

}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
