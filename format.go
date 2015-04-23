package uuid

import (
	"fmt"
)

// Format returns the formatted uuid string by given style.
func Format(uuid UUID, style Style) string {
	if len(uuid) != 16 {
		panic("uuid: UUID is invalid")
	}

	buffer := []byte(uuid)
	switch style {
	case StyleStandard:
		return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x", buffer[:4], buffer[4:6], buffer[6:8], buffer[8:10], buffer[10:])
	case StyleWithoutDash:
		return fmt.Sprintf("%x", buffer[:])
	default:
		panic("uuid: style of UUID is invalid")
	}
}
