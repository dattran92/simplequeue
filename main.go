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

	fmt.Println("---Enter to continue---")
	var input string
  fmt.Scanln(&input)

	fmt.Println("---Log data in queue before and after picking---")
	fmt.Println("Values ", queue.Values)
	pick := queue.Dequeue()

	fmt.Println("Pick ", *pick)
	fmt.Println("Remain ", queue.Values)

	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()

	fmt.Println("---In case nothing else left---")
	fmt.Println("Pick ", queue.Dequeue())

  fmt.Println("done")
}
