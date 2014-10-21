/*

Package md5 implements the name-based version (v3) of UUID that uses MD5 hashing - specified in RFC 4122.

Usage:

	generator := md5.New()
	generator.Configure(namespace, name string) or generator.Configure(name string)
	generator.Generate() ([]byte, error)

*/
package md5
