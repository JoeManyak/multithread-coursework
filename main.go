package main

import (
	"fmt"
	"math"
	"multidfs/tree"
	"time"
)

const depth = 15

func main() {
	taskSize := int(math.Pow(2, depth))

	fmt.Println("--------------------")
	fmt.Println("Tree Generating...")
	t := tree.GenerateTree(depth)
	fmt.Println("--------------------")
	fmt.Println("--------------------")
	fmt.Println("Multi Thread Search...")
	now := time.Now()
	println(t.MultiTaskThreadSearch(4, taskSize, 3))
	//println(t.MultiThreadSearch(12))
	time.Sleep(time.Nanosecond)
	fmt.Println(time.Since(now))
	fmt.Printf("Multi thread estimate:  %d nanoseconds\n", time.Since(now).Nanoseconds())

	t.RemoveVisitors()
	time.Sleep(time.Second * 2)

	fmt.Println("--------------------")
	fmt.Println("Single Thread Search...")
	now = time.Now()
	println(t.SingleThreadSearch())
	time.Sleep(time.Nanosecond)
	fmt.Println(time.Since(now))
	fmt.Printf("Single thread estimate: %d nanoseconds\n", time.Since(now).Nanoseconds())

	t.RemoveVisitors()

	time.Sleep(time.Second * 5)

}
