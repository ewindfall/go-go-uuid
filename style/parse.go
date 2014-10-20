package style

import (
	"bytes"
	"errors"
	"github.com/landjur/go-uuid/layout"
	"github.com/landjur/go-uuid/version"
)

var (
	nilUUID = make([]byte, 16)
)

var (
	ErrLayoutInvalid  = errors.New("uuid: layout is invalid")
	ErrVersionUnknown = errors.New("uuid: version is invalid")
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

// parse parses the standard uuid string, format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx (8-4-4-4-12, length: 36).
func parseStandard(uuidString string) ([]byte, error) {
	if len(uuidString) != 36 {
		return nil, errors.New("uuid: length of uuid string is invalid, it should be 36")
	}

	if uuidString[8] != '-' || uuidString[13] != '-' || uuidString[18] != '-' || uuidString[23] != '-' {
		return nil, errors.New("uuid: format of uuid string is invalid, it should be xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx (8-4-4-4-12)")
	}

	uuid := make([]byte, 16)

	for i, v := range []int{0, 2, 4, 6, 9, 11, 14, 16, 19, 21, 24, 26, 28, 30, 32, 34} {
		uuid[i] = (hex(uuidString[v]) << 4) | hex(uuidString[v+1])
	}

	if !bytes.Equal(uuid, nilUUID) {
		if layout.Get(uuid) == layout.Invalid {
			return nil, ErrLayoutInvalid
		}

		if version.Get(uuid) == version.Unknown {
			return nil, ErrVersionUnknown
		}
	}

	return uuid, nil
}

// parse parses the uuid string without dash, format: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx (length: 32).
func parseWithoutDash(uuidString string) ([]byte, error) {
	if len(uuidString) != 32 {
		return nil, errors.New("uuid: length of uuid string is invalid, it should be 36")
	}

	uuid := make([]byte, 16)

	for i, v := range []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30} {
		uuid[i] = (hex(uuidString[v]) << 4) | hex(uuidString[v+1])
	}

	if !bytes.Equal(uuid, nilUUID) {
		if layout.Get(uuid) == layout.Invalid {
			return nil, ErrLayoutInvalid
		}

		if version.Get(uuid) == version.Unknown {
			return nil, ErrVersionUnknown
		}
	}

	return uuid, nil
}

// Parse parses the uuid string.
func Parse(uuidString string) ([]byte, error) {
	switch len(uuidString) {
	case 36:
		return parseStandard(uuidString)
	case 32:
		return parseWithoutDash(uuidString)
	}

	return nil, errors.New("uuid: length of uuid string is invalid, it should be 36 (standard) or 32 (without dash)")
}
