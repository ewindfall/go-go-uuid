package version

// Version represents the version of UUID.
type Version byte

const (
	Unknown       Version = iota // Unknwon
	TimeBased                    // V1: The time-based version
	DCESecurity                  // V2: The DCE security version, with embedded POSIX UIDs
	NameBasedMD5                 // V3: The name-based version that uses MD5 hashing
	Randomly                     // V4: The randomly or pseudo-randomly generated version
	NameBasedSHA1                // V5: The name-based version that uses SHA-1 hashing
)

// String returns English description of version.
func (this Version) String() string {
	switch this {
	case TimeBased:
		return "Time-based version"
	case DCESecurity:
		return "DCE security version"
	case NameBasedMD5:
		return "Name-based version that uses MD5 hashing"
	case Randomly:
		return "Randomly or pseudo-randomly generated version"
	case NameBasedSHA1:
		return "Name-based version that uses SHA-1 hashing"
	}

	return "Unknwon version"
}

// Get gets the version of UUID.
func Get(uuid []byte) Version {
	if len(uuid) != 16 {
		panic("uuid: length of UUID is invalid, it should be 16")
	}

	version := uuid[6] >> 4
	if version > 0 && version < 6 {
		return Version(version)
	}

	return Unknown
}

// Set sets version if UUID.
func Set(uuid []byte, v Version) {
	if len(uuid) != 16 {
		panic("uuid: length of UUID is invalid, it should be 16")
	}

	switch v {
	case TimeBased:
		uuid[6] = (uuid[6] | 0x10) & 0x1f
	case DCESecurity:
		uuid[6] = (uuid[6] | 0x20) & 0x2f
	case NameBasedMD5:
		uuid[6] = (uuid[6] | 0x30) & 0x3f
	case Randomly:
		uuid[6] = (uuid[6] | 0x40) & 0x4f
	case NameBasedSHA1:
		uuid[6] = (uuid[6] | 0x50) & 0x5f
	default:
		panic("uuid: unknwon version of UUID")
	}
}
