# LTK-TE

This repository contains a function that adds the numbers at each line of a file
including sub files.

## Usage

### Prerequisite

To run this project, make sure you are using:

  * Python 3
  * Any version of Go

### Go

### Function usage

Here is a simple example of usage for the function:

```go
package main

import "tlk"

func main() {
  total, err := ltk.AddRec(args[1])
  if err != nil {
    panic(err)
  }
  fmt.Printf("Full total is %d\n", total)
}
```

#### Program from binary

Binaries are available for the following platform:

-   MacOS 64-bit
-   Linux 64-bit
-   Windows 64-bit

#### Build from source

To build from source you need `go` installed and then you can run the Makefile:

    $ > make
    GOOS=darwin GOARCH=amd64 go build -o ./bin/sums_darwin_amd64
    GOOS=windows GOARCH=amd64 go build -o ./bin/sums_win_amd64.exe
    GOOS=linux GOARCH=amd64 go build -o ./bin/sums_linux_amd64

#### Run tests

To run the tests you can use the Makefile:

    $ > make test
    === RUN   TestAddRec
    Subtotal for file <some_path>/tests/test_set/d.txt is 6
    Subtotal for file <some_path>/tests/test_set/e.txt is 22
    Subtotal for file <some_path>/tests/test_set/b.txt is 78
    Subtotal for file <some_path>/tests/test_set/e.txt is 22
    Subtotal for file <some_path>/tests/test_set/c.txt is 15
    Subtotal for file <some_path>/tests/test_set/a.txt is 78
    --- PASS: TestAddRec (0.00s)
        sums_test.go:17: Total is 221
    === RUN   TestAddRecSimple
    Subtotal for file <some_path>/tests/test_set_simple/b.txt is 5
    Subtotal for file <some_path>/tests/test_set_simple/a.txt is 5
    --- PASS: TestAddRecSimple (0.00s)
        sums_test.go:28: Total is 10
    === RUN   TestAddRecLoopSimple
    --- PASS: TestAddRecLoopSimple (0.00s)
        sums_test.go:24: Recursive loop: <some_path>/tests/test_set_loop_simple/a.txt previously called
    PASS

### Python

### Function usage

Here is a simple example of usage for the function:

```python
from tlk import sum, RecursiveLoopException

def someOtherFunction():
  try:
    total = sum('/path/to/text/file')
  except RecursiveLoopException as e:
    raise e

  print('Total is {:d}'.format(total))
```

#### Run tests

To run the tests you can use run:

    $ > ./tlk_test.py

The output should be the following:

    Subtotal for tests/test_set_loop_complex/d.txt is 6
    Subtotal for tests/test_set/d.txt is 6
    Subtotal for tests/test_set/e.txt is 22
    Subtotal for tests/test_set/b.txt is 78
    Subtotal for tests/test_set/e.txt is 22
    Subtotal for tests/test_set/c.txt is 15
    Subtotal for tests/test_set/a.txt is 78
    Subtotal for tests/test_set_simple/b.txt is 5
    Subtotal for tests/test_set_simple/a.txt is 5
    ...
    ----------------------------------------------------------------------
    Ran 3 tests in 0.003s

    OK
