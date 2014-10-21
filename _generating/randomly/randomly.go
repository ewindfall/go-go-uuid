package randomly

import (
	"crypto/rand"
	"errors"
	"github.com/landjur/go-uuid/generating"
	"github.com/landjur/go-uuid/layout"
	"github.com/landjur/go-uuid/version"
)

// init registers the generator self.
func init() {
	generating.Randomly.Register(New)
}

// New returns a new generator for generating time-based uuid.
func New() generating.Generator {
	return new(generator)
}

// generator represents the time-based uuid generator.
type generator struct {
}

// Configure configures the generator.
func (this generator) Configure(args ...interface{}) error {
	return errors.New("uuid: the randomly generator is not configurable")
}

// Generate returns a new randomly uuid.
func (this generator) Generate() ([]byte, error) {
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid[:])
	if n != len(uuid) || err != nil {
		return nil, err
	}

	version.Set(uuid, version.Randomly)
	layout.Set(uuid, generating.Layout)

	return uuid, nil
}
