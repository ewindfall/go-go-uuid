package uuid

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/landjur/go-uuid/dce"
	"github.com/landjur/go-uuid/md5"
	"github.com/landjur/go-uuid/random"
	"github.com/landjur/go-uuid/sha1"
	"github.com/landjur/go-uuid/time"
	"github.com/landjur/go-uuid/utility"
)

// Consts
const (
	// The standard namespaces for UUIDs
	NamespaceDNS  = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	NamespaceURL  = "6ba7b811-9dad-11d1-80b4-00c04fd430c8"
	NamespaceOID  = "6ba7b812-9dad-11d1-80b4-00c04fd430c8"
	NamespaceX500 = "6ba7b814-9dad-11d1-80b4-00c04fd430c8"
)

var (
	// The EmptyUUID is specified to have all 128 bits set to zero. Specified in RFC 4122(section 4.1.7).
	Empty = new(UUID)
)

// import dce.Domain
var (
	DomainUser  = dce.DomainUser
	DomainGroup = dce.DomainGroup
)

// Errors
var (
	ErrLengthInvalid  = errors.New("uuid: length of UUID string is invalid, it should be 36")
	ErrFormatInvalid  = errors.New("uuid: format of UUID string is invalid, it should be xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx [8-4-4-4-12])")
	ErrLayoutInvalid  = errors.New("uuid: layout of UUID is invalid")
	ErrVersionUnknown = errors.New("uuid: version of UUID is unknown")
)

// hex returns hex value for given char. Called by Parse.
func hex(v byte) byte {
	switch v {
	case '0':
		return 0x00
	case '1':
		return 0x01
	case '2':
		return 0x02
	case '3':
		return 0x03
	case '4':
		return 0x04
	case '5':
		return 0x05
	case '6':
		return 0x06
	case '7':
		return 0x07
	case '8':
		return 0x08
	case '9':
		return 0x09
	case 'A', 'a':
		return 0x0a
	case 'B', 'b':
		return 0x0b
	case 'C', 'c':
		return 0x0c
	case 'D', 'd':
		return 0x0d
	case 'E', 'e':
		return 0x0e
	case 'F', 'f':
		return 0x0f
	}

	return 0xff
}

// Parse returns instance of UUID by parse from string
// It returns ErrLengthInvalid if length of UUID string is invalid.
// It returns ErrFormatInvalid if format of UUID string is invalid.
// It returns ErrLayoutInvalid if layout of UUID parsed is invalid.
// It returns ErrVersionInvalid if version of UUID parsed is invalid.
func Parse(v string) (*UUID, error) {
	if len(v) != 36 {
		return nil, ErrLengthInvalid
	}

	// 8-4-4-4-12: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	if v[8] != '-' || v[13] != '-' || v[18] != '-' || v[23] != '-' {
		return nil, ErrFormatInvalid
	}

	id := new(UUID)

	for i, v2 := range []int{0, 2, 4, 6, 9, 11, 14, 16, 19, 21, 24, 26, 28, 30, 32, 34} {
		id[i] = (hex(v[v2]) << 4) | hex(v[v2+1])
	}

	if !id.Equal(Empty) {
		if id.Layout() == LayoutInvalid {
			return nil, ErrLayoutInvalid
		}

		if id.Version() == VersionUnknown {
			return nil, ErrVersionUnknown
		}
	}

	return id, nil
}

// IsValid checks the passed value whether is a valid UUID string.
func IsValid(v string) bool {
	_, err := Parse(v)
	if err != nil {
		return false
	}

	return true
}

// NewTimeUUID returns a new Version 1 UUID.
// Same as NewV1UUID function.
func NewTimeUUID() (*UUID, error) {
	u, err := time.New()
	if err != nil {
		return nil, err
	}

	uuid := new(UUID)
	copy(uuid[:], u)

	return uuid, nil
}

// NewV1UUID call NewTimeUUID directly.
func NewV1UUID() (*UUID, error) {
	return NewTimeUUID()
}

// NewDceUUID returns a new Version 2 UUID.
// Same as NewV2UUID function.
func NewDceUUID(domain dce.Domain) (*UUID, error) {
	u, err := dce.New(domain)
	if err != nil {
		return nil, err
	}

	uuid := new(UUID)
	copy(uuid[:], u)

	return uuid, nil
}

// NewV2UUID call NewDceUUID directly.
func NewV2UUID(domain dce.Domain) (*UUID, error) {
	return NewDceUUID(domain)
}

// NewMD5UUID returns a new Version 3 UUID.
// Same as NewV3UUID function.
func NewMD5UUID(namespace, name string) (*UUID, error) {
	u, err := md5.New(namespace, name)
	if err != nil {
		return nil, err
	}

	uuid := new(UUID)
	copy(uuid[:], u)

	return uuid, nil
}

// NewV3UUID call NewMD5UUID directly.
func NewV3UUID(namespace, name string) (*UUID, error) {
	return NewMD5UUID(namespace, name)
}

// NewRandomUUID returns a new Version 4 UUID.
// Same as NewV4UUID function.
// Same as New function.
func NewRandomUUID() (*UUID, error) {
	u, err := random.New()
	if err != nil {
		return nil, err
	}

	uuid := new(UUID)
	copy(uuid[:], u)

	return uuid, nil
}

// NewV4UUID call NewRandomUUID directly.
func NewV4UUID() (*UUID, error) {
	return NewRandomUUID()
}

// NewSHA1UUID returns a new Version 5 UUID.
// Same as NewV5UUID function.
func NewSHA1UUID(namespace, name string) (*UUID, error) {
	u, err := sha1.New(namespace, name)
	if err != nil {
		return nil, err
	}

	uuid := new(UUID)
	copy(uuid[:], u)

	return uuid, nil
}

// NewV5UUID call NewSHA1UUID directly.
func NewV5UUID(namespace, name string) (*UUID, error) {
	return NewSHA1UUID(namespace, name)
}

// UUID respresents a UUID object compliant with specification in RFC 4122.
type UUID [16]byte

// Layout returns layout of UUID.
func (this UUID) Layout() Layout {
	return Layout(utility.Layout(this[:]))
}

// Version returns version of UUID.
func (this UUID) Version() Version {
	return Version(utility.Version(this[:]))
}

// Equal returns true if current uuid equal to passed uuid.
func (this UUID) Equal(v *UUID) bool {
	return bytes.Equal(this[:], (*v)[:])
}

// String returns formatted string for UUID (format[8-4-4-4-12]: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)
func (this UUID) String() string {
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x", this[:4], this[4:6], this[6:8], this[8:10], this[10:])
}
