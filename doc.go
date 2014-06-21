/*

Package uuid implement UUID V1, V2, V3, V4, V5 specified in RFC 4122.

Usage:

Generate V1 UUID:

        uuid.NewV1UUID()
        uuid.NewTimeUUID()

Generate V2 UUID:

        uuid.NewV2UUID(domain Domain)
        uuid.NewDceUUID(domain Domain)

Generate V3 UUID:

        uuid.NewV3UUID(namespace, name string)
        uuid.NewMD5UUID(namespace, name string)

Generate V4 UUID:

        uuid.NewV4UUID()
        uuid.NewRandomUUID()

Generate V5 UUID:

        uuid.NewV5UUID(namespace, name string)
        uuid.NewSHA1UUID(namespace, name string)

*/
package uuid
