// Copyright 2014 Landjur. All rights reserved.

package uuid

import (
	"fmt"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	s := "00000000-0000-0000-0000-000000000000"
	start := time.Now()
	u, err := Parse(s)
	spent := time.Since(start)
	if err != nil {
		t.Fatalf("uuid parse failed: %s", err)
	}

	if !u.Equal(Empty) {
		t.Fatal("uuid parsed should equal emtpy")
	}

	fmt.Printf("Parse UUID(%s) spent %v seconds (%v nanoseconds)\n", u, spent.Seconds(), spent.Nanoseconds())
}

func TestUUID(t *testing.T) {
	// V1
	uuid, err := NewTimeUUID()
	if err != nil {
		t.Fatalf("create time uuid failed: %s", err)
	}
	version := uuid.Version()
	if version != VersionTime {
		t.Fatalf("Time Version UUID has wrong version: %s", version)
	}
	layout := uuid.Layout()
	if layout != LayoutRFC4122 {
		t.Fatalf("Time Version UUID has wrong layout: %s", layout)
	}
	fmt.Println("Version 1 (Time Based):\t", uuid)

	// V2
	uuid, err = NewDceUUID(DomainUser)
	if err != nil {
		t.Fatalf("create dce uuid failed: %s", err)
	}
	version = uuid.Version()
	if version != VersionDCE {
		t.Fatalf("DCE Version UUID has wrong version: %s", version)
	}
	layout = uuid.Layout()
	if layout != LayoutRFC4122 {
		t.Fatalf("DCE Version UUID has wrong layout: %s", layout)
	}
	fmt.Println("Version 2 (DCE):\t", uuid)

	// V3
	uuid, err = NewMD5UUID("namespace", "name")
	if err != nil {
		t.Fatalf("create md5 uuid failed: %s", err)
	}
	version = uuid.Version()
	if version != VersionMD5 {
		t.Fatalf("MD5 Version UUID has wrong version: %s", version)
	}
	layout = uuid.Layout()
	if layout != LayoutRFC4122 {
		t.Fatalf("MD5 Version UUID has wrong layout: %s", layout)
	}
	fmt.Println("Version 3 (MD5):\t", uuid)

	// V4
	uuid, err = NewRandomUUID()
	if err != nil {
		t.Fatalf("create random uuid failed: %s", err)
	}
	version = uuid.Version()
	if version != VersionRandom {
		t.Fatalf("Random Version UUID has wrong version: %s", version)
	}
	layout = uuid.Layout()
	if layout != LayoutRFC4122 {
		t.Fatalf("Random Version UUID has wrong layout: %s", layout)
	}
	fmt.Println("Version 4 (Randomly):\t", uuid)

	// V5
	uuid, err = NewSHA1UUID("namespace", "name")
	if err != nil {
		t.Fatalf("create sha1 uuid failed: %s", err)
	}
	version = uuid.Version()
	if version != VersionSHA1 {
		t.Fatalf("SHA1 Version UUID has wrong version: %s", version)
	}
	layout = uuid.Layout()
	if layout != LayoutRFC4122 {
		t.Fatalf("SHA1 Version UUID has wrong layout: %s", layout)
	}
	fmt.Println("Version 5 (SHA1):\t", uuid)
}
