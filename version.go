package uuid

// Version represents version of UUID.
type Version byte

const (
	VersionUnknown Version = iota // Unknown Version
	VersionTime                   // V1: The time-based version specified in RFC 4122.
	VersionDCE                    // V2: DCE Security version, with embedded POSIX UIDs.
	VersionMD5                    // V3: The name-based version specified in RFC 4122 that uses MD5 hashing.
	VersionRandom                 // V4: The randomly or pseudo-randomly generated version specified in RFC 4122.
	VersionSHA1                   // V5: The name-based version specified in RFC 4122 that uses SHA-1 hashing.
	maxVersion
)

// String returns English description of version.
func (this Version) String() string {
	switch this {
	case VersionTime:
		return "Version: Time Based(V1 specified in RFC 4122)"
	case VersionDCE:
		return "Version: DCE Security(V2 specified in RFC 4122)"
	case VersionMD5:
		return "Version: Name Based uses MD5 hashing(V3 specified in RFC 4122)"
	case VersionRandom:
		return "Version: Randomly or Pseudo-Randomly(V4 specified in RFC 4122)"
	case VersionSHA1:
		return "Version: Name Based uses SHA-1 hashing(V5 specified in RFC 4122)"
	}

	return "Unknown Version"
}
