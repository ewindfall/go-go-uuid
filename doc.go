/*

Package uuid implement UUID specified in RFC 4122.

Usage:

Generate Time-Based UUID:

        uuid.NewTimeBased() (UUID, error)
		uuid.NewV1() (UUID, error)

Generate DCE Security UUID:

		uuid.NewDCESecurity(uuid.DomainUser or uuid.DomainGroup)  (UUID, error)
		uuid.NewV2(uuid.DomainUser or uuid.DomainGroup) (UUID, error)

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
