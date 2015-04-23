package uuid

// Layout represents the layout of UUID. See page 5 in RFC 4122.
type Layout byte

const (
	LayoutInvalid   Layout = iota // Invalid
	LayoutNCS                     // Reserved, NCS backward compatibility. (Values: 0x00-0x07)
	LayoutRFC4122                 // The variant specified in RFC 4122. (Values: 0x08-0x0b)
	LayoutMicrosoft               // Reserved, Microsoft Corporation backward compatibility. (Values: 0x0c-0x0d)
	LayoutFuture                  // Reserved for future definition. (Values: 0x0e-0x0f)
)

// String returns English description of layout.
func (this Layout) String() string {
	switch this {
	case LayoutNCS:
		return "Layout: Reserved For NCS"
	case LayoutRFC4122:
		return "Layout: RFC 4122"
	case LayoutMicrosoft:
		return "Layout: Reserved For Microsoft"
	case LayoutFuture:
		return "Layout: Reserved For Future"
	default:
		return "Layout: Invalid"
	}
}

// ParseLayout parses the layout of UUID.
func ParseLayout(uuid UUID) Layout {
	if len(uuid) == 16 {
		switch {
		case (uuid[8] & 0x80) == 0x00:
			return LayoutNCS
		case (uuid[8] & 0xc0) == 0x80:
			return LayoutRFC4122
		case (uuid[8] & 0xe0) == 0xc0:
			return LayoutMicrosoft
		case (uuid[8] & 0xe0) == 0xe0:
			return LayoutFuture
		}
	}

	return LayoutInvalid
}

// SetLayout sets the layout of UUID.
func SetLayout(uuid UUID, layout Layout) {
	if len(uuid) != 16 {
		panic("uuid: length of UUID is invalid")
	}

	switch layout {
	case LayoutNCS:
		uuid[8] = (uuid[8] | 0x00) & 0x0f // Msb0=0
	case LayoutRFC4122:
		uuid[8] = (uuid[8] | 0x80) & 0x8f // Msb0=1, Msb1=0
	case LayoutMicrosoft:
		uuid[8] = (uuid[8] | 0xc0) & 0xcf // Msb0=1, Msb1=1, Msb2=0
	case LayoutFuture:
		uuid[8] = (uuid[8] | 0xe0) & 0xef // Msb0=1, Msb1=1, Msb2=1
	default:
		panic("uuid: layout of UUID is invalid")
	}
}
