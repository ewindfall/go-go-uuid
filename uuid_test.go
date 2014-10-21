package uuid

import (
	"fmt"
	"github.com/landjur/go-uuid/layout"
	"github.com/landjur/go-uuid/version"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	s := "00000000-0000-0000-0000-000000000000"
	start := time.Now()
	uuid, err := Parse(s)
	spent := time.Since(start)
	if err != nil {
		t.Fatalf("UUID parsing failed: %s", err)
	}

	if !uuid.Equal(Nil) {
		t.Fatal("UUID should equal to nil uuid")
	}

	fmt.Printf("Parse UUID(%s) spent %v seconds (%v nanoseconds)\n", uuid, spent.Seconds(), spent.Nanoseconds())
}

func TestUUID(t *testing.T) {
	// V1
	uuid, err := NewTimeBased()
	if err != nil {
		t.Fatalf("create time uuid failed: %s", err)
	}
	v := uuid.Version()
	if v != version.TimeBased {
		t.Fatalf("Time-Based UUID has wrong version: %s", v)
	}
	l := uuid.Layout()
	if l != layout.RFC4122 {
		t.Fatalf("Time-Based UUID has wrong layout: %s", l)
	}
	fmt.Println("Version 1 (Time Based):", uuid, "     ", uuid.Format(StyleWithoutDash, false))

	// V2
	uuid, err = NewDCESecurity(DomainUser)
	if err != nil {
		t.Fatalf("create dce uuid failed: %s", err)
	}
	v = uuid.Version()
	if v != version.DCESecurity {
		t.Fatalf("DCE Security UUID has wrong version: %s", v)
	}
	l = uuid.Layout()
	if l != layout.RFC4122 {
		t.Fatalf("DCE Security UUID has wrong layout: %s", l)
	}
	fmt.Println("Version 2 (DCE Security):", uuid, "     ", uuid.Format(StyleWithoutDash, false))

	// V3
	uuid, err = NewNameBasedMD5("namespace", "name")
	if err != nil {
		t.Fatalf("create md5 uuid failed: %s", err)
	}
	v = uuid.Version()
	if v != version.NameBasedMD5 {
		t.Fatalf("Name-Based (MD5) UUID has wrong version: %s", v)
	}
	l = uuid.Layout()
	if l != layout.RFC4122 {
		t.Fatalf("Name-Based (MD5) UUID has wrong layout: %s", l)
	}
	fmt.Println("Version 3 (Name-Based uses MD5 hashing):", uuid, "     ", uuid.Format(StyleWithoutDash, false))

	// V4
	uuid, err = NewRandomly()
	if err != nil {
		t.Fatalf("create randomly uuid failed: %s", err)
	}
	v = uuid.Version()
	if v != version.Randomly {
		t.Fatalf("Randomly UUID has wrong version: %s", v)
	}
	l = uuid.Layout()
	if l != layout.RFC4122 {
		t.Fatalf("Randomly UUID has wrong layout: %s", l)
	}
	fmt.Println("Version 4 (Randomly):", uuid, "", uuid.Format(StyleWithoutDash, false))

	// V5
	uuid, err = NewNameBasedSHA1("namespace", "name")
	if err != nil {
		t.Fatalf("create sha1 uuid failed: %s", err)
	}
	v = uuid.Version()
	if v != version.NameBasedSHA1 {
		t.Fatalf("Name-Based (SHA1) UUID has wrong version: %s", v)
	}
	l = uuid.Layout()
	if l != layout.RFC4122 {
		t.Fatalf("Name-Based (SHA1) UUID has wrong layout: %s", l)
	}
	fmt.Println("Version 5 (Name-Based uses SHA-1 hashing):", uuid, "     ", uuid.Format(StyleWithoutDash, false))
}
