package random

import (
	"fmt"
	"github.com/wayn3h0/go-uuid"
	"testing"
)

func Test(t *testing.T) {
	u, err := New()
	if err != nil {
		t.Fatal(err)
	}

	if u.Version() != uuid.VersionRandom {
		t.Fatal("wrong version")
	}

	if u.Layout() != uuid.LayoutRFC4122 {
		t.Fatal("wrong layout")
	}

	fmt.Println(u)
}
