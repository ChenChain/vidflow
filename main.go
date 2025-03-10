package main

import (
	"fmt"
	"vidflow/handler/facade"
)

func main() {
	fmt.Println("vid flow is making")
	facade.NewHello().Healthy(nil)
}
