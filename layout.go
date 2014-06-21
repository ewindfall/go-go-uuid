package uuid

// Layout represents layout of UUID.
type Layout byte

const (
	LayoutInvalid   Layout = iota // Invalid Layout
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
	}

	return "Invalid(Unknown) Layout"
}
