package dcesecurity

import (
	"encoding/binary"
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

// Generate returns a new DCE security uuid.
func New(domain Domain) ([]byte, error) {
	uuid, err := timebased.New()
	if err != nil {
		return nil, err
	}

	switch domain {
	case DomainUser:
		uid := os.Getuid()
		binary.BigEndian.PutUint32(uuid[0:], uint32(uid)) // network byte order
	case DomainGroup:
		gid := os.Getgid()
		binary.BigEndian.PutUint32(uuid[0:], uint32(gid)) // network byte order
	}

	version.Set(uuid, version.DCESecurity)
	layout.Set(uuid, layout.RFC4122)

	return uuid, nil
}
