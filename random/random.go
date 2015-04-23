package random

import (
	"crypto/rand"
	"github.com/wayn3h0/go-uuid"
)

// New returns a new randomly uuid.
func New() (uuid.UUID, error) {
	instance := make(uuid.UUID, 16)
	n, err := rand.Read(instance[:])
	if n != len(instance) || err != nil {
		return nil, err
	}

	uuid.SetVersion(instance, uuid.VersionRandom)
	uuid.SetLayout(instance, uuid.LayoutRFC4122)

	return instance, nil
}
