package utils

import "fmt"

func Dump(args ...interface{})  {
	fmt.Print("\n----------------------------------------\n")
	for _, arg := range args {
		fmt.Printf("%v", arg)
	}
	fmt.Print("\n----------------------------------------\n")
}
