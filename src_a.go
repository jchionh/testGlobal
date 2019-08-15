package main

import "fmt"

func ModifyAndPrintCount() {

	// increment the global var
	count++

	// print the value after modification
	fmt.Printf("Modified count value: %d\n", count)
}
