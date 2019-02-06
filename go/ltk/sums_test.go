package ltk_test

import (
	"testing"

	ltk "."
)

func TestAddRec(t *testing.T) {
	total, err := ltk.AddRec("tests/test_set/a.txt")
	if err != nil {
		t.Fatal(err)
	}
	if total != 221 {
		t.Fatalf("Invalid total, expected 221 got %d", total)
	}
	t.Logf("Total is %d\n", total)
}

func TestAddRecLoopSimple(t *testing.T) {
	if _, err := ltk.AddRec("tests/test_set_loop_simple/a.txt"); err == nil {
		t.Fatal("Expected an error")
	} else {
		t.Log(err)
	}
}

//
// func TestAddRecLoopComplex(t *testing.T) {
// 	var total int
// 	if err := ltk.AddRec(&total, "./test_set_loop_complex/a.txt", ""); err == nil {
// 		t.Fatal("Expected an error")
// 	} else {
// 		t.Log(err)
// 	}
// }
