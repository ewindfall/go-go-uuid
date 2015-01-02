package style

import (
	"fmt"
	"strings"
)

// Format returns the formatted uuid string by given style.
func Format(uuid []byte, style Style, upper bool) string {
	if len(uuid) != 16 {
		panic("uuid: length of UUID is invalid, it should be 16")
	}

	var formatted string

	switch style {
	case Standard:
		formatted = fmt.Sprintf("%08x-%04x-%04x-%04x-%012x", uuid[:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
	case WithoutDash:
		formatted = fmt.Sprintf("%x", uuid[:])
	}

	if upper {
		return strings.ToUpper(formatted)
	}

	return formatted

	panic("uuid: invalid style for formatting")
}
