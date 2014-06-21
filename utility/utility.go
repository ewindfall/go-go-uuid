package utility

// Layout returns the layout of given UUID.
func Layout(uuid []byte) byte {
	if len(uuid) != 16 {
		panic("uuid: Length of UUID is invalid, it should be 16")
	}

	switch {
	case (uuid[8] & 0x80) == 0x00:
		return 1 //LayoutNCS
	case (uuid[8] & 0xc0) == 0x80:
		return 2 //LayoutRFC4122
	case (uuid[8] & 0xe0) == 0xc0:
		return 3 //LayoutMicrosoft
	case (uuid[8] & 0xe0) == 0xe0:
		return 4 //LayoutFuture
	}

	return 0 //LayoutInvalid
}

// SetLayout sets the layout of given UUID. See page 5 in RFC 4122.
// This is intended to be called from the packages that implement uuid functions.
func SetLayout(uuid []byte, layout byte) {
	if len(uuid) != 16 {
		panic("uuid: Length of UUID is invalid, it should be 16")
	}

	switch layout {
	case 1: //LayoutNCS: // NCS
		uuid[8] = (uuid[8] | 0x00) & 0x0f // Msb0=0
	case 2: // RFC4122:
		uuid[8] = (uuid[8] | 0x80) & 0x8f // Msb0=1, Msb1=0
	case 3: // Microsoft:
		uuid[8] = (uuid[8] | 0xc0) & 0xcf // Msb0=1, Msb1=1, Msb2=0
	case 4: // Future:
		uuid[8] = (uuid[8] | 0xe0) & 0xef // Msb0=1, Msb1=1, Msb2=1
	default:
		panic("uuid: Invalid Layout of UUID")
	}
}

// Version returns version of given UUID.
func Version(uuid []byte) byte {
	if len(uuid) != 16 {
		panic("uuid: Length of UUID is invalid, it should be 16")
	}

	version := uuid[6] >> 4
	if version > 0 && version < 6 {
		return version
	}

	return 0 //VersionUnknown
}

// SetVersion sets version if UUID. See page 7 in RFC 4122.
// This is intended to be called from the packages that implement uuid functions.
func SetVersion(uuid []byte, version byte) {
	if len(uuid) != 16 {
		panic("uuid: Length of UUID is invalid, it should be 16")
	}

	switch version {
	case 1: //VersionTime: // V1
		uuid[6] = (uuid[6] | 0x10) & 0x1f
	case 2: //VersionDCE: // V2
		uuid[6] = (uuid[6] | 0x20) & 0x2f
	case 3: //VersionMD5: // V3
		uuid[6] = (uuid[6] | 0x30) & 0x3f
	case 4: //VersionRandom: // V4
		uuid[6] = (uuid[6] | 0x40) & 0x4f
	case 5: //VersionSHA1: //V5
		uuid[6] = (uuid[6] | 0x50) & 0x5f
	default:
		panic("uuid: Unknwon Version of UUID")
	}
}
