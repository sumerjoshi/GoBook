package main

import (
	"bufio"
	"fmt"
	"os"
)

//Dup1 prints the text of each line that ppears more than
//once in the standard input, preceded by its count


func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input..Text()]++
	}
	//NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1{
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}


/*

Loop over the input, show some computation for each line of vode.
We will show three variants here called dup. It is inspired 
by the Unix uniq command which will look for adjacent lines.

For if statement, parentheses are required for the body.
There can be ab optional else part that is executed if the condition
is false

A map holds a set of k/v pairs and provides constant time operations
to store retrive or test for an item in a set.

Each time dup reads a line of input 
the line is used as a skey into the map and corresponding map is incremented.

counts[input.Text()]++ -> 
	line := input.Text()
	counts[line] = counts[line] + 1


Onward to the bufio package, which helps make input and ouput efficient
and convienient. One of its most useful features is a type called Scanner that
reads input and breaks into lines or words. Often the easiest way to process
input that comes naturally to lines

A map holds a set of key/valye pairs and provides constant time operations to 
store, retrieve, or test for items in the set. The key may be of any type whose values
can be compared with ==, strings being the most common example

Each time dup reads a line of input, the line is used as a key into the map
of the corresponding value.

We use another range-based for loop that iterates over the counts map. 
Each iteration produces two results (a key and the value of the map element for that key)

bufio.Scanner reads from program's standard input.


*/