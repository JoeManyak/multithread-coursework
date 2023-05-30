package main

import (
	"fmt"
	"math"
	"multidfs/tree"
	"time"
)

const (
	depth     = 20
	depthTask = 11
)

func main() {
	//should use this in case of binary tree as maximum tasks
	taskSize := int(math.Pow(2, depth))

	fmt.Println("--------------------")
	fmt.Println("Tree Generating...")
	t := tree.GenerateTree(depth)
	fmt.Println("--------------------")
	fmt.Println("--------------------")
	fmt.Println("Multi Thread Search...")
	now := time.Now()
	println(t.MultiTaskThreadSearch(12, taskSize, depthTask))
	time.Sleep(time.Nanosecond)
	fmt.Println(time.Since(now))
	fmt.Printf("Multi thread estimate:  %d nanoseconds\n", time.Since(now).Nanoseconds())
}
