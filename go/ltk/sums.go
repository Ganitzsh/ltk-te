package ltk

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Node struct {
	Name   string
	Parent string
}

func exists(s string, arr []*Node) bool {
	for _, v := range arr {
		if s == v.Parent {
			return true
		}
	}
	return false
}

// AddRec takes a file as parameter and sums the numbers at each line, it also
// accepts other files as input in the very same file and sums them up as well
func AddRec(total *int, path string, history *[]*Node) error {
	var subTotal, line int
	if total == nil {
		return errors.New("Nil total, cannot compute")
	}
	if history == nil {
		history = &[]*Node{&Node{Name: path, Parent: ""}}
	}
	b, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Error opening file %s: %s", path, err)
	}
	defer b.Close()
	rd := bufio.NewReader(b)
	for {
		l, _, err := rd.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("Error reading line: %s", err)
		}
		lStr := string(l)
		v, err := strconv.Atoi(lStr)
		if err != nil {
			if lStr == path {
				return fmt.
					Errorf("Recursive loop in %s, self called at line %d", path, line+1)
			}
			// if exists(lStr, history) {
			// 	return fmt.Errorf(
			// 		"File %s: recusrive loop, file %s already in the stack",
			// 		path, lStr,
			// 	)
			// }
			if e := AddRec(total, lStr, history); e != nil {
				return fmt.Errorf("File %s: %s", path, e)
			}
			*history = append(*history, &Node{Name: lStr, Parent: path})
		}
		subTotal += v
		*total += v
		line += 1
	}
	fmt.Printf("Subtotal for file %s is %d\n", path, subTotal)
	return nil
}
