package layout

// Layout represents layout of UUID.
type Layout byte

const (
	Invalid   Layout = iota // Invalid
	NCS                     // Reserved, NCS backward compatibility. (Values: 0x00-0x07)
	RFC4122                 // The variant specified in RFC 4122. (Values: 0x08-0x0b)
	Microsoft               // Reserved, Microsoft Corporation backward compatibility. (Values: 0x0c-0x0d)
	Future                  // Reserved for future definition. (Values: 0x0e-0x0f)
)

// String returns English description of layout.
func (this Layout) String() string {
	switch this {
	case NCS:
		return "Reserved layout for NCS"
	case RFC4122:
		return "RFC 4122 layout"
	case Microsoft:
		return "Reserved layout for Microsoft"
	case Future:
		return "Reserved layout for Future"
	}

	return "Unknown layout"
}

// Get gets the layout of UUID.
func Get(uuid []byte) Layout {
	if len(uuid) != 16 {
		panic("uuid: length of UUID is invalid, it should be 16")
	}

	switch {
	case (uuid[8] & 0x80) == 0x00:
		return NCS
	case (uuid[8] & 0xc0) == 0x80:
		return RFC4122
	case (uuid[8] & 0xe0) == 0xc0:
		return Microsoft
	case (uuid[8] & 0xe0) == 0xe0:
		return Future
	}

	return Invalid
}

// Set sets the layout of UUID.
func Set(uuid []byte, l Layout) {
	if len(uuid) != 16 {
		panic("uuid: length of UUID is invalid, it should be 16")
	}

	switch l {
	case NCS:
		uuid[8] = (uuid[8] | 0x00) & 0x0f // Msb0=0
	case RFC4122:
		uuid[8] = (uuid[8] | 0x80) & 0x8f // Msb0=1, Msb1=0
	case Microsoft:
		uuid[8] = (uuid[8] | 0xc0) & 0xcf // Msb0=1, Msb1=1, Msb2=0
	case Future:
		uuid[8] = (uuid[8] | 0xe0) & 0xef // Msb0=1, Msb1=1, Msb2=1
	default:
		panic("uuid: invalid layout of UUID")
	}
}
