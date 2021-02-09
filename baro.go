package baro

type ObjectType string

type Item interface {
	GetID() uint64
	Encode() string
	Decode(string) error
}

type User interface {
    GetID() uint64
    GetObject() string
    Encode() string
    Decode(string) error
}

type Role interface {
    GetID() uint64
    GetParentID() uint64
    GetObject() string
    Encode() string
    Decode(string) error
}

type Object interface {
    GetID() uint64
    GetParentID() uint64
    GetObject() string
    GetObjectType() ObjectType
    Encode() string
    Decode(string) error
}

type Domain interface {
    GetID() uint64
    GetObject() string
    Encode() string
    Decode(string) error
}

type currentProvider interface {
    Get() (User, Domain, error)
}

type IB interface {
    GetUsersForRoleInDomain(Role, Domain) []User
    GetRolesForUserInDomain(User, Domain) []Role


}

type MetaDataBase interface {

    // Domain API
    GetAllDomain() []Domain
}
