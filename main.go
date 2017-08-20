package main

import (
	"fmt"
  "github.com/dattran92/simplequeue/queue"
)

func main() {
  queue := queue.Init()

	go queue.Enqueue(3)
	go queue.Enqueue(4)
	go queue.Enqueue(5)
	go queue.Enqueue(6)
	go queue.Enqueue(9)

	var input string
  fmt.Scanln(&input)

	fmt.Println("Values ", queue.Values)
	pick := queue.Dequeue()

	fmt.Println("Pick ", pick)
	fmt.Println("Remain ", queue.Values)

  fmt.Println("done")
}
