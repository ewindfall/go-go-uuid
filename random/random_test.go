package random

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	id, err := New()
	if err != nil {
		t.Fatalf("generate v4 uuid failed: %s", err)
	}

	fmt.Printf("%08x-%04x-%04x-%04x-%012x\n", id[:4], id[4:6], id[6:8], id[8:10], id[10:])
}
