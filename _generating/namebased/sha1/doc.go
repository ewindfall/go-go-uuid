/*

Package sha1 implements the name-based version (v5) of UUID that uses MD5 hashing - specified in RFC 4122.

Usage:

	generator := sha1.New()
	generator.Configure(namespace, name string) or generator.Configure(name string)
	generator.Generate() ([]byte, error)

*/
package sha1
