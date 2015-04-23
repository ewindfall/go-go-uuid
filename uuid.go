package uuid

import (
	"bytes"
)

var (
	Nil = make(UUID, 16) // Nil UUID
)

// UUID respresents an UUID type compliant with specification in RFC 4122.
type UUID []byte

// Layout returns layout of UUID.
func (this UUID) Layout() Layout {
	return ParseLayout(this)
}

func (this UUID) SetLayout(layout Layout) {
	SetLayout(this, layout)
}

// Version returns version of UUID.
func (this UUID) Version() Version {
	return ParseVersion(this)
}

func (this UUID) SetVersion(version Version) {
	SetVersion(this, version)
}

// Equal returns true if current uuid equal to passed uuid.
func (this UUID) Equal(another UUID) bool {
	return bytes.EqualFold(this, another)
}

// String returns the string of UUID (standard style: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx | 8-4-4-4-12)
func (this UUID) String() string {
	return Format(this, StyleStandard)
}

// Format returns the formatted string of UUID.
func (this UUID) Format(style Style) string {
	return Format(this, style)
}
