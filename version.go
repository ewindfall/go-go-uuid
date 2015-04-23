package uuid

// Version represents the version of UUID. See page 7 in RFC 4122.
type Version byte

const (
	VersionUnknown       Version = iota // Unknwon
	VersionTimeBased                    // V1: The time-based version
	VersionDCESecurity                  // V2: The DCE security version, with embedded POSIX UIDs
	VersionNameBasedMD5                 // V3: The name-based version that uses MD5 hashing
	VersionRandom                       // V4: The randomly or pseudo-randomly generated version
	VersionNameBasedSHA1                // V5: The name-based version that uses SHA-1 hashing
)

// String returns English description of version.
func (this Version) String() string {
	switch this {
	case VersionTimeBased:
		return "Version: Time-Based"
	case VersionDCESecurity:
		return "Version: DCE Security With Embedded POSIX UIDs"
	case VersionNameBasedMD5:
		return "Version: Name-Based (MD5)"
	case VersionRandom:
		return "Version: Randomly OR Pseudo-Randomly Generated"
	case VersionNameBasedSHA1:
		return "Version: Name-Based (SHA-1)"
	default:
		return "Version: Unknwon"
	}
}

// ParseVersion parses the version of UUID.
func ParseVersion(uuid UUID) Version {
	if len(uuid) == 16 {
		version := uuid[6] >> 4
		if version > 0 && version < 6 {
			return Version(version)
		}
	}

	return VersionUnknown
}

// SetVersion sets version of UUID.
func SetVersion(uuid UUID, version Version) {
	if len(uuid) != 16 {
		panic("uuid: length of UUID is invalid")
	}

	switch version {
	case VersionTimeBased:
		uuid[6] = (uuid[6] | 0x10) & 0x1f
	case VersionDCESecurity:
		uuid[6] = (uuid[6] | 0x20) & 0x2f
	case VersionNameBasedMD5:
		uuid[6] = (uuid[6] | 0x30) & 0x3f
	case VersionRandom:
		uuid[6] = (uuid[6] | 0x40) & 0x4f
	case VersionNameBasedSHA1:
		uuid[6] = (uuid[6] | 0x50) & 0x5f
	default:
		panic("uuid: version of UUID is unknown")
	}
}
