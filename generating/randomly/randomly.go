package randomly

import (
	"crypto/rand"
	"github.com/landjur/go-uuid/layout"
	"github.com/landjur/go-uuid/version"
)

// New returns a new randomly uuid.
func New() ([]byte, error) {
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid[:])
	if n != len(uuid) || err != nil {
		return nil, err
	}

	version.Set(uuid, version.Randomly)
	layout.Set(uuid, layout.RFC4122)

	return uuid, nil
}
