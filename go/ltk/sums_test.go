package ltk_test

import (
	"testing"

	ltk "."
)

func TestAddRec(t *testing.T) {
	var total int
	stack := []*ltk.Node{}
	if err := ltk.AddRec(&total, "./test_set/a.txt", &stack); err != nil {
		t.Fatal(err)
	}
	if total != 221 {
		t.Fatalf("Invalid total, expected 221 got %d", total)
	}
	t.Logf("Total is %d\n", total)
}
