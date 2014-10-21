/*

Package dcesecurity implements the DCE Security version (v2) of UUID that embeded POSIX UIDs - specified in RFC 4122.

Usage:

	generator := dcesecurity.New()
	generator.Configure(dcesecurity.Domain)
	generator.Generate() ([]byte, error)

*/
package dcesecurity
