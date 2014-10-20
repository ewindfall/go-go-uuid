package uuid

import (
	"bytes"
	"github.com/landjur/go-uuid/generating/dcesecurity"
	"github.com/landjur/go-uuid/generating/namebased"
	"github.com/landjur/go-uuid/generating/namebased/md5"
	"github.com/landjur/go-uuid/generating/namebased/sha1"
	"github.com/landjur/go-uuid/generating/randomly"
	"github.com/landjur/go-uuid/generating/timebased"
	"github.com/landjur/go-uuid/layout"
	"github.com/landjur/go-uuid/style"
	"github.com/landjur/go-uuid/version"
)

const (
	StandardStyle    = style.Standard
	WithoutDashStyle = style.WithoutDash
)
const (
	UserDomain  = dcesecurity.DomainUser
	GroupDomain = dcesecurity.DomainGroup
)
const (
	NamespaceDNS  = namebased.NamespaceDNS
	NamespaceURL  = namebased.NamespaceURL
	NamespaceOID  = namebased.NamespaceOID
	NamespaceX500 = namebased.NamespaceX500
)

var (
	Nil = make(UUID, 16) // Nil UUID
)

// NewTimeBased returns a new time-based uuid.
func NewTimeBased() (UUID, error) {
	uuid, err := timebased.New()
	if err != nil {
		return nil, err
	}

	return UUID(uuid), nil
}

// NewV1 call NewTimeBased function.
func NewV1() (UUID, error) {
	return NewTimeBased()
}

// NewDCESecurity returns a new DCE security uuid.
func NewDCESecurity(domain dcesecurity.Domain) (UUID, error) {
	uuid, err := dcesecurity.New(domain)
	if err != nil {
		return nil, err
	}

	return UUID(uuid), nil
}

// NewV2 calls NewDCESecurity function.
func NewV2(domain dcesecurity.Domain) (UUID, error) {
	return NewDCESecurity(domain)
}

// NewNameBasedMD5 returns a new name-based uuid uses MD5 hashing.
func NewNameBasedMD5(namespace, name string) (UUID, error) {
	uuid, err := md5.New(namespace, name)
	if err != nil {
		return nil, err
	}

	return UUID(uuid), nil
}

// NewV3 calls NewNameBasedMD5 function.
func NewV3(namespace, name string) (UUID, error) {
	return NewNameBasedMD5(namespace, name)
}

// NewRandomly returns a new randomly uuid.
func NewRandomly() (UUID, error) {
	uuid, err := randomly.New()
	if err != nil {
		return nil, err
	}

	return UUID(uuid), nil
}

// NewV4 calls NewRandomly function.
func NewV4() (UUID, error) {
	return NewRandomly()
}

// NewNameBasedSHA1 returns a new name-based uuid uses SHA-1 hashing.
func NewNameBasedSHA1(namespace, name string) (UUID, error) {
	uuid, err := sha1.New(namespace, name)
	if err != nil {
		return nil, err
	}

	return UUID(uuid), nil
}

// NewV5 calls NewNameBasedSHA1 function.
func NewV5(namespace, name string) (UUID, error) {
	return NewNameBasedSHA1(namespace, name)
}

// Parse parses the uuid string.
func Parse(uuidString string) (UUID, error) {
	return style.Parse(uuidString)
}

// IsValid checks the passed value whether is a valid UUID string.
func IsValid(uuidString string) bool {
	_, err := Parse(uuidString)
	if err != nil {
		return false
	}

	return true
}

// UUID respresents an UUID type compliant with specification in RFC 4122.
type UUID []byte

// Layout returns layout of UUID.
func (this UUID) Layout() layout.Layout {
	return layout.Get(this)
}

func (this UUID) SetLayout(l layout.Layout) {
	layout.Set(this, l)
}

// Version returns version of UUID.
func (this UUID) Version() version.Version {
	return version.Get(this)
}

// Equal returns true if current uuid equal to passed uuid.
func (this UUID) Equal(another UUID) bool {
	return bytes.Equal(this, another)
}

// String returns the string of UUID (standard style: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx | 8-4-4-4-12)
func (this UUID) String() string {
	return style.Format(this, style.Standard, false)
}

// Format returns the formatted string of UUID.
func (this UUID) Format(s style.Style, upper bool) string {
	return style.Format(this, s, upper)
}
