package sha1

import (
	"crypto/sha1"
	"errors"
	"github.com/landjur/go-uuid/generating"
	"github.com/landjur/go-uuid/generating/namebased"
	"github.com/landjur/go-uuid/layout"
	"github.com/landjur/go-uuid/version"
)

// Errors
var (
	ErrConfigurationInvalid = errors.New("uuid: wrong configuration for name-based (SHA-1) generator, it should be Configure(namespace, name string) or Configure(name string)")
)

// init registers the generator self.
func init() {
	generating.NameBasedSHA1.Register(New)
}

// New returns a new generator for generating name-based uuid that uses sha1 hashing.
func New() generating.Generator {
	instance := new(generator)
	instance.namespace = namebased.NamespaceDNS

	return instance
}

// generator represents the time-based uuid generator.
type generator struct {
	namespace string
	name      string
}

// Configure configures the generator.
// Usage: Configure(namespace, name string) or Configure(name string)
func (this *generator) Configure(args ...interface{}) error {
	switch len(args) {
	case 1:
		name, ok := args[0].(string)
		if !ok {
			return ErrConfigurationInvalid
		}
		this.name = name
		return nil
	case 2:
		namespace, ok := args[0].(string)
		if !ok {
			return ErrConfigurationInvalid
		}
		name, ok := args[1].(string)
		if !ok {
			return ErrConfigurationInvalid
		}
		this.namespace = namespace
		this.name = name
		return nil
	default:
		return ErrConfigurationInvalid
	}
}

// Generate returns a new name-based uses sha1 hashing uuid.
func (this generator) Generate() ([]byte, error) {
	hash := sha1.New()
	_, err := hash.Write([]byte(this.namespace))
	if err != nil {
		return nil, err
	}
	_, err = hash.Write([]byte(this.name))
	if err != nil {
		return nil, err
	}

	sum := hash.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, sum)

	version.Set(uuid, version.NameBasedSHA1)
	layout.Set(uuid, generating.Layout)

	return uuid, nil
}
