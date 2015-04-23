package uuid

import (
	"errors"
)

// hex returns hex value for given char.
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

// Parse parses the UUID string.
func Parse(uuidString string) (UUID, error) {
	length := len(uuidString)
	buffer := make([]byte, 16)
	switch length {
	case 36:
		if uuidString[8] != '-' || uuidString[13] != '-' || uuidString[18] != '-' || uuidString[23] != '-' {
			return nil, errors.New("uuid: format of UUID string is invalid, it should be xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx (8-4-4-4-12)")
		}
		for i, v := range []int{0, 2, 4, 6, 9, 11, 14, 16, 19, 21, 24, 26, 28, 30, 32, 34} {
			buffer[i] = (hex(uuidString[v]) << 4) | hex(uuidString[v+1])
		}
	case 32:
		for i, v := range []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30} {
			buffer[i] = (hex(uuidString[v]) << 4) | hex(uuidString[v+1])
		}
	default:
		return nil, errors.New("uuid: length of uuid string is invalid, it should be 36 (standard) or 32 (without dash)")
	}

	uuid := UUID(buffer)

	if !uuid.Equal(Nil) {
		if uuid.Layout() == LayoutInvalid {
			return nil, errors.New("uuid: layout of UUID is invalid")
		}

		if uuid.Version() == VersionUnknown {
			return nil, errors.New("uuid: version of UUID is unknown")
		}
	}

	return uuid, nil
}
