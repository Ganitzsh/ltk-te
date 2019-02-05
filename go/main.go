package main

import (
	"fmt"
	"os"

	"./ltk"
	"github.com/davecgh/go-spew/spew"
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
	stack := []*ltk.Node{}
	args := os.Args
	if len(args) == 1 || len(args) > 2 {
		displayUsage()
		return
	}
	if err := ltk.AddRec(&total, args[1], &stack); err != nil {
		panic(err)
	}
	spew.Dump(stack)
	fmt.Printf("Full total is %d\n", total)
}
