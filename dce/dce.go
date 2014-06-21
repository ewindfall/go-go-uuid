package dce

import (
	"encoding/binary"
	"errors"
	"github.com/landjur/go-uuid/time"
	"github.com/landjur/go-uuid/utility"
	"os"
)

var (
	ErrDomainInvalid = errors.New("uuid/dce: domain of uuid is invalid, it should be DomainUser or DomainGroup")
)

// Domain represents the identifier for a local domain
type Domain byte

const (
	DomainUser  Domain = iota + 1 // POSIX UID domain
	DomainGroup                   // POSIX GID domain
)

// New returns the uuid by given domain(UID or GID).
func New(domain Domain) ([]byte, error) {
	id, err := time.New()
	if err != nil {
		return nil, err
	}

	switch domain {
	case DomainUser:
		uid := os.Getuid()
		binary.BigEndian.PutUint32(id[0:], uint32(uid)) // network byte order
	case DomainGroup:
		gid := os.Getgid()
		binary.BigEndian.PutUint32(id[0:], uint32(gid)) // network byte order
	default:
		return nil, ErrDomainInvalid
	}

	utility.SetVersion(id, 2) // Version 2
	utility.SetLayout(id, 2)  // RFC4122 Layout

	return id, nil
}
