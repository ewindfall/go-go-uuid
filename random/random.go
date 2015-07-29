package random

import (
	"crypto/rand"
	"github.com/wayn3h0/go-errors"
	"github.com/wayn3h0/go-uuid"
)

// New returns a new randomly uuid.
func New() (uuid.UUID, error) {
	instance := make(uuid.UUID, 16)
	n, err := rand.Read(instance[:])
	if err != nil {
		return nil, errors.Wrap(err, "uuid/random: generating pseudorandom numbers failed")
	}
	if n != len(instance) {
		return nil, errors.New("uuid/random: generating pseudorandom numbers failed")
	}

	uuid.SetVersion(instance, uuid.VersionRandom)
	uuid.SetLayout(instance, uuid.LayoutRFC4122)

	return instance, nil
}
