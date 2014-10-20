/*

Package uuid implement UUID specified in RFC 4122.

Usage:

Import Package:

		import "github.com/landjur/go-uuid"

Generate Time-Based UUID:

        uuid.NewTimeBased() (UUID, error)
		uuid.NewV1() (UUID, error)

Generate DCE Security UUID:

		uuid.NewDCESecurity(uuid.UserDomain or uuid.GroupDomain)  (UUID, error)
		uuid.NewV2(uuid.UserDomain or uuid.GroupDomain) (UUID, error)

Generate Name-Based UUID uses MD5 hashing:

		uuid.NewNameBasedMD5(namespace, name string) (UUID, error)
		uuid.NewV3(namespace, name string) (UUID, error)

Generate V4 UUID:

        uuid.NewRandomly() (UUID, error)
        uuid.NewV4() (UUID, error)

Generate V5 UUID:

		uuid.NewNameBasedSHA1(namespace, name string) (UUID, error)
        uuid.NewV5(namespace, name string) (UUID, error)

*/
package uuid
