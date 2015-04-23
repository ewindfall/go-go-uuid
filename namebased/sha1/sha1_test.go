package sha1

import (
	"fmt"
	"github.com/wayn3h0/go-uuid"
	"github.com/wayn3h0/go-uuid/namebased"
	"testing"
)

func Test(t *testing.T) {
	u, err := New(namebased.NamespaceDNS, "name")
	if err != nil {
		t.Fatal(err)
	}

	if u.Version() != uuid.VersionNameBasedSHA1 {
		t.Fatal("wrong version")
	}

	if u.Layout() != uuid.LayoutRFC4122 {
		t.Fatal("wrong layout")
	}

	fmt.Println(u)
}
