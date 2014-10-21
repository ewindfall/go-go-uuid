package generating

// Implementation represents the implementation of UUID generator.
type Implementation byte

// Implementations
const (
	TimeBased     Implementation = iota + 1 // V1 Generator: The time-based version.
	DCESecurity                             // V2 Generator: The DCE security version, with embedded POSIX UIDs.
	NameBasedMD5                            // V3 Generator: The name-based version that uses MD5 hashing.
	Randomly                                // V4 Generator: The randomly or pseudo-randomly generated version.
	NameBasedSHA1                           // V5 Generator: The name-based version that uses SHA-1 hashing.
	maxImplementation
)

var generators = make([]func() Generator, maxImplementation)

// New returns a new instance of generator implementation. New panics if the generating implementation is not linked into the binary.
func (this Implementation) New() Generator {
	if this > 0 && this < maxImplementation {
		function := generators[this]
		if function != nil {
			return function()
		}
	}

	panic("uuid: requested generator is unavailable")
}

// Available reports whether the given generating implementation is linked into the binary.
func (this Implementation) Available() bool {
	return this < maxImplementation && generators[this] != nil
}

// Register registers an uuid generating implementation.
// This is intended to be called from the packages that implement uuid generating functions.
func (this Implementation) Register(function func() Generator) {
	if this < 0 && this > maxImplementation {
		panic("uuid: registered unknown generator")
	}

	generators[this] = function
}
