package main

import (
	"fmt"
	"multidfs/tree"
	"time"
)

const depth = 27

func main() {
	fmt.Println("Tree Generating...")
	t := tree.GenerateTree(depth)
	for {
		fmt.Println("--------------------")
		fmt.Println("Multi Thread Search...")
		now := time.Now()
		fmt.Println(t.MultiThreadSearch(5))
		fmt.Printf("Multi thread estimate:  %d nanoseconds\n", time.Since(now).Nanoseconds())

		fmt.Println("--------------------")
		fmt.Println("Single Thread Search...")
		now = time.Now()
		fmt.Println(t.SingleThreadSearch())
		fmt.Printf("Single thread estimate: %d nanoseconds\n", time.Since(now).Nanoseconds())
		time.Sleep(time.Second * 5)
	}
}
