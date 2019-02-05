package ltk

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// AddRec takes a file as parameter and sums the numbers at each line, it also
// accepts other files as input in the very same file and sums them up as well
func AddRec(total *int, path string, previouses ...string) error {
	var subTotal, line int
	if total == nil {
		return errors.New("Nil total, cannot compute")
	}
	for _, previous := range previouses {
		if previous == path {
			return fmt.Errorf(
				"Recursive loop: %s previously called %s",
				path, strings.Join(previouses[1:len(previouses)], "->"),
			)
		}
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
			if e := AddRec(total, lStr, append(previouses, path)...); e != nil {
				return e
			}
		}
		subTotal += v
		*total += v
		line += 1
	}
	fmt.Printf("Subtotal for file %s is %d\n", path, subTotal)
	return nil
}
