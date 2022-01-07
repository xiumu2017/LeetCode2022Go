package main

import "os"

func main() {
	println("hello world ~")
	args := os.Args
	for i, i2 := range args {
		println(i)
		println(i2)
	}
	os.Exit(0)

}
