package style

// Style represents the style of UUID string.
type Style byte

const (
	Standard    Style = iota + 1 // xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx (8-4-4-4-12, length: 36)
	WithoutDash                  // xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx (length: 32)
)

// String returns English description of style.
func (this Style) String() string {
	switch this {
	case Standard:
		return "Style: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx (8-4-4-4-12, length: 36)"
	case WithoutDash:
		return "Style: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx (length: 32)"
	}

	return "Unknown style"
}
