// Copyright 2014 Landjur. All rights reserved.

package md5

import (
	"bitbucket.org/landjur-golang/uuid/utility"
	"crypto/md5"
)

// New returns a new uuid by namespace and name given.
func New(namespace, name string) ([]byte, error) {
	hash := md5.New()
	_, err := hash.Write([]byte(namespace))
	if err != nil {
		return nil, err
	}
	_, err = hash.Write([]byte(name))
	if err != nil {
		return nil, err
	}

	sum := hash.Sum(nil)
	id := make([]byte, 16)
	copy(id, sum)

	utility.SetVersion(id, 3) // Version 3
	utility.SetLayout(id, 2)  // RFC4122 Layout

	return id, nil
}
