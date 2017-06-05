package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"k8s.io/kubernetes/pkg/api/resource"
)

func main() {
	var (
		input string
		err   error
	)
	if len(os.Args) != 2 {
		reader := bufio.NewReader(os.Stdin)
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}
	} else {
		input = os.Args[1]
	}
	input = strings.Trim(input, "\n")
	qty, err := resource.ParseQuantity(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	d := qty.AsDec()
	i, ok := d.Unscaled()
	if !ok {
		fmt.Println("Unable to output as int64")
		os.Exit(1)
	}
	fmt.Println(i)
}
