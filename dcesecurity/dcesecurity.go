package dcesecurity

import (
	"encoding/binary"
	"errors"
	"github.com/wayn3h0/go-uuid"
	"github.com/wayn3h0/go-uuid/timebased"
	"os"
)

// Domain represents the identifier for a local domain
type Domain byte

const (
	DomainUser  Domain = iota + 1 // POSIX UID domain
	DomainGroup                   // POSIX GID domain
)

// Generate returns a new DCE security uuid.
func New(domain Domain) (uuid.UUID, error) {
	instance, err := timebased.New()
	if err != nil {
		return nil, err
	}

	switch domain {
	case DomainUser:
		uid := os.Getuid()
		binary.BigEndian.PutUint32(instance[0:], uint32(uid)) // network byte order
	case DomainGroup:
		gid := os.Getgid()
		binary.BigEndian.PutUint32(instance[0:], uint32(gid)) // network byte order
	default:
		return nil, errors.New("uuid/dcesecurity: domain is invalid")
	}

	uuid.SetVersion(instance, uuid.VersionDCESecurity)
	uuid.SetLayout(instance, uuid.LayoutRFC4122)

	return instance, nil
}
