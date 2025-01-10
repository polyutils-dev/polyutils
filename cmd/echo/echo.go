package main

import "os"

func main() {
	for i, arg := range os.Args {
		if i == 0 {
			continue // Skip the first argument, which is the executable name.
		}
		print(arg + " ")
	}
	println()
}
