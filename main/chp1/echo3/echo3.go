package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""
	for index, arg := range os.Args[1:] {
		s += sep + strconv.Itoa(index) + arg
		sep = "\n"
	}
	fmt.Println(s)

}

/*
1.1 - Invoke name of file that called this
	- /var/folders/dv/br1s6gnx3h760vvl_0slwtzr0000gn/T/go-build2007631752/b001/exe/echo3 a b c

1.2 - With the echo program here is that we are looping
	through index, arg with a range here
	- Convert index to a string with a seperation here

*/
