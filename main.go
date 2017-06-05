package main

import (
	"fmt"
	"os"

	"k8s.io/kubernetes/pkg/api/resource"
)

func main() {
	qty, err := resource.ParseQuantity(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	i, ok := qty.AsInt64()
	if !ok {
		fmt.Println("Unable to output as int64")
		os.Exit(1)
	}
	fmt.Println(i)
}
