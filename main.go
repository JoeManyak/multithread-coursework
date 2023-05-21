package main

import (
	"fmt"
	"multidfs/tree"
)

const depth = 3

func main() {
	t := tree.GenerateTree(depth)
	fmt.Println(t)
}
