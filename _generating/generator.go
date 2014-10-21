package generating

// Generator represents the interface of an UUID generator.
type Generator interface {
	// Configure configures the generator.
	Configure(args ...interface{}) error

	// Generate generates a new UUID.
	Generate() ([]byte, error)
}
