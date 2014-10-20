package md5

import (
	"crypto/md5"
	"github.com/landjur/go-uuid/layout"
	"github.com/landjur/go-uuid/version"
)

// New returns a new name-based uses MD5 hashing uuid.
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
	uuid := make([]byte, 16)
	copy(uuid, sum)

	version.Set(uuid, version.NameBasedMD5)
	layout.Set(uuid, layout.RFC4122)

	return uuid, nil
}
