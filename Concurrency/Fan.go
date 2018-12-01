package Concurrency

func FanOut(input chan int, task func(int, chan int)) chan int {
	result := make(chan int)
	go func() {
		for {

		}
	}()
}
