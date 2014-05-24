// Copyright 2014 Landjur. All rights reserved.

package random

import (
	"bitbucket.org/landjur-golang/uuid/utility"
	"crypto/rand"
)

// New returns a new uuid randomly.
func New() ([]byte, error) {
	id := make([]byte, 16)
	n, err := rand.Read(id[:])
	if n != len(id) || err != nil {
		return nil, err
	}

	utility.SetVersion(id, 4) // Version 4
	utility.SetLayout(id, 2)  // RFC4122 Layout

	return id, nil
}
