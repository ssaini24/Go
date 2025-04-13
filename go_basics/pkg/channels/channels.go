package channels

import "fmt"

func Init() {
	fmt.Println("\n Go Channels Init...")
	var c = make(chan int)
	go process(c)
	for i := range c {
		fmt.Println(i)
	}
}

func process(c chan int) {
	defer close(c)
	for i := 1; i < 5; i++ {
		c <- i
	}
}
