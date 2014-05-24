#Golang UUID Component

##Overview

The package implements UUID [RFC 4122](http://www.ietf.org/rfc/rfc4122.txt).

##Usage

Generate V1 UUID: 

    uuid.NewV1UUID() 
    uuid.NewTimeUUID()

Generate V2 UUID: 

    uuid.NewV2UUID(domain dce.Domain) 
    uuid.NewDceUUID(domain dce.Domain)

Generate V3 UUID:
    
    uuid.NewV3UUID(namespace, name string) 
    uuid.NewMD5UUID(namespace, name string)

Generate V4 UUID: 
	
	uuid.NewV4UUID()
	uuid.NewRandomUUID()

Generate V5 UUID: 
    
    uuid.NewV5UUID(namespace, name string)
    uuid.NewSHA1UUID(namespace, name string)
