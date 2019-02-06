package main

import (
	"fmt"
	"os"

	"./ltk"
)

func displayUsage() {
	fmt.Println(`Usage: sums file.txt

Description

This program takes a file and counts the total of each numbers at each line
including the subtotal of other subsequent files specified instead of a
number`)
}

func main() {
	var total int
	args := os.Args
	if len(args) == 1 || len(args) > 2 {
		displayUsage()
		return
	}
	total, err := ltk.AddRec(args[1])
	if err != nil {
		panic(err)
	}
	fmt.Printf("Full total is %d\n", total)
}
