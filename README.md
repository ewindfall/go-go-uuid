#Golang UUID Package

The package implements UUID [RFC 4122](http://www.ietf.org/rfc/rfc4122.txt).

##Usage

###Import Package

    import "github.com/landjur/go-uuid"

###Generate Time-Based UUID

    uuid.NewTimeBased() (UUID, error)
    uuid.NewV1() (UUID, error)

###Generate DCE Security UUID

    uuid.NewDCESecurity(uuid.UserDomain or uuid.GroupDomain)  (UUID, error)
    uuid.NewV2(uuid.UserDomain or uuid.GroupDomain) (UUID, error)

###Generate Name-Based UUID uses MD5 hashing

    uuid.NewNameBasedMD5(namespace, name string) (UUID, error)
    uuid.NewV3(namespace, name string) (UUID, error)

###Generate Randomly UUID

    uuid.NewRandomly() (UUID, error)
    uuid.NewV4() (UUID, error)

###Generate Name-Based UUID uses SHA-1 hashing

    uuid.NewNameBasedSHA1(namespace, name string) (UUID, error)
    uuid.NewV5(namespace, name string) (UUID, error)

##COPYRIGHT & LICENSE
Copyright 2014 Landjur, Inc. Code released under the Apache License, Version 2.0.
