package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

/*

The os package provides functions and other values for dealing with the OS
-> anywhere outside of the os package is os.Args

The var declaration declares two variables s and sep, of type string. A variable cabe initialized
as part of the declaration. If it is not explicitly initialized, it
is implicityly initialized to find zer o value for its type, which is 0 for numeric types
and the empty string "" for strings.

When applied to strings however the + operator concatenates the values

operator += is an assignment operator.

For loop is the only loop statement in Go.

for initializaiton; condition; post {
	// zero or more statements
}

Optional initialization statement is initialized before the loop starts. condition is a boolean statement

//traditional while loop
for condition {

}

//traditional infinite loop
for {
	//
}



*/
