package dcesecurity

import (
	"fmt"
	"github.com/wayn3h0/go-uuid"
	"testing"
)

func Test(t *testing.T) {
	u, err := New(DomainUser)
	if err != nil {
		t.Fatal(err)
	}

	if u.Version() != uuid.VersionDCESecurity {
		t.Fatal("wrong version")
	}

	if u.Layout() != uuid.LayoutRFC4122 {
		t.Fatal("wrong layout")
	}

	fmt.Println(u)
}
