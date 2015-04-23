#Go-UUID 

The go-uuid package implements UUID [RFC 4122](http://www.ietf.org/rfc/rfc4122.txt) for golang.

##Usage

###Generating 

####Time-Based (Version 1)

    import "github.com/wayn3h0/go-uuid/timebased"

    timebased.New() (uuid.UUID, error)

####DCE Security (Version 2)

    import "github.com/wayn3h0/go-uuid/dcesecurity"

    dcesecurity.New(dcesecurity.Domain) (uuid.UUID, error)

####Name-Based uses MD5 hashing (Version 3)

    import "github.com/wayn3h0/go-uuid/namebased/md5"

    md5.New(namespace, name string) (uuid.UUID, error)

####Random (Version 4)

    import "github.com/wayn3h0/go-uuid/random"

    random.New() (uuid.UUID, error)

####Name-Based uses SHA-1 hashing (Version 5)

    import "github.com/wayn3h0/go-uuid/namebased/sha1"

    sha1.New(namespace, name string) (uuid.UUID, error)

###Styles

* Standard: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx (8-4-4-4-12, length: 36)
* Without Dash: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx (length: 32)

###Formatting & Parsing

    import "github.com/wayn3h0/go-uuid"
    
    (UUID Instance).String() string             // format to standard style
    (UUID Instance).Format(uuid.Style) string   // format to uuid.StyleStandard or uuid.StyleWithoutDash

    uuid.Parse(string) (uuid.UUID, error)       // parse from UUID string
