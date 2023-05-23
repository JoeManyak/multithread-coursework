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
		fmt.Println("Multi thread estimate: ", time.Since(now))

		fmt.Println("--------------------")
		fmt.Println("Single Thread Search...")
		now = time.Now()
		fmt.Println(t.SingleThreadSearch())
		fmt.Println("Single thread estimate: ", time.Since(now))
		time.Sleep(time.Second * 5)
	}
}
