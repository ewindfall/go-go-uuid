package md5

import (
	"crypto/md5"
	"github.com/wayn3h0/go-errors"
	"github.com/wayn3h0/go-uuid"
)

// New returns a new name-based uses SHA-1 hashing uuid.
func New(namespace, name string) (uuid.UUID, error) {
	hash := md5.New()
	_, err := hash.Write([]byte(namespace))
	if err != nil {
		return nil, errors.Wrap(err, "uuid/namebased/md5: computing hash value failed")
	}
	_, err = hash.Write([]byte(name))
	if err != nil {
		return nil, errors.Wrap(err, "uuid/namebased/md5: computing hash value failed")
	}

	sum := hash.Sum(nil)
	instance := make(uuid.UUID, 16)
	copy(instance, sum)

	uuid.SetVersion(instance, uuid.VersionNameBasedMD5)
	uuid.SetLayout(instance, uuid.LayoutRFC4122)

	return instance, nil
}
