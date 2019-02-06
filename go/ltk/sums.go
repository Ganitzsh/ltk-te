package ltk

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// AddRec takes a file as parameter and sums the numbers at each line, it also
// accepts other files as input in the very same file and sums them up as well
func AddRec(path string, previouses ...string) (int, error) {
	var total, subTotal, line int
	abs, err := filepath.Abs(path)
	if err != nil {
		return 0, err
	}
	for _, previous := range previouses {
		if previous == abs {
			return 0, fmt.Errorf(
				"Recursive loop: %s previously called %s",
				abs, strings.Join(previouses[1:len(previouses)], "->"),
			)
		}
	}
	b, err := os.Open(abs)
	if err != nil {
		return 0, fmt.Errorf("Error opening file %s: %s", abs, err)
	}
	defer b.Close()
	rd := bufio.NewReader(b)
	for {
		l, _, err := rd.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, fmt.Errorf("Error reading line: %s", err)
		}
		lStr := string(l)
		v, err := strconv.Atoi(lStr)
		if err != nil {
			if t, e := AddRec(lStr, append(previouses, abs)...); e != nil {
				return 0, e
			} else {
				total += t
			}
		}
		subTotal += v
		line += 1
	}
	fmt.Printf("Subtotal for file %s is %d\n", abs, subTotal)
	return subTotal + total, nil
}
