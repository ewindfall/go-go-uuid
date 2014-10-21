package dcesecurity

import (
	"encoding/binary"
	"errors"
	"github.com/landjur/go-uuid/generating"
	"github.com/landjur/go-uuid/generating/timebased"
	"github.com/landjur/go-uuid/layout"
	"github.com/landjur/go-uuid/version"
	"os"
)

// Domain represents the identifier for a local domain
type Domain byte

const (
	DomainUser  Domain = iota + 1 // POSIX UID domain
	DomainGroup                   // POSIX GID domain
)

// Errors
var (
	ErrConfigurationInvalid = errors.New("uuid: wrong configuration for DCE security generator, it should be Configure(dcesecurity.Domain)")
)

// init registers the generator self.
func init() {
	generating.DCESecurity.Register(New)
}

// New returns a new generator for generating time-based uuid.
func New() generating.Generator {
	instance := new(generator)
	instance.domain = DomainUser

	return instance
}

// generator represents the time-based uuid generator.
type generator struct {
	domain Domain
}

// Configure configures the generator.
// Usage: Configure(Domain)
func (this *generator) Configure(args ...interface{}) error {
	if len(args) != 1 {
		return ErrConfigurationInvalid
	}

	domain, ok := args[0].(Domain)
	if !ok {
		return ErrConfigurationInvalid
	}

	this.domain = domain

	return nil
}

// Generate returns a new DCE security uuid.
func (this generator) Generate() ([]byte, error) {
	uuid, err := timebased.New().Generate()
	if err != nil {
		return nil, err
	}

	switch this.domain {
	case DomainUser:
		uid := os.Getuid()
		binary.BigEndian.PutUint32(uuid[0:], uint32(uid)) // network byte order
	case DomainGroup:
		gid := os.Getgid()
		binary.BigEndian.PutUint32(uuid[0:], uint32(gid)) // network byte order
	}

	version.Set(uuid, version.DCESecurity)
	layout.Set(uuid, generating.Layout)

	return uuid, nil
}
