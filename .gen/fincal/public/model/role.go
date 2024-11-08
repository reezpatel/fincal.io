//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import "errors"

type Role string

const (
	Role_Owner     Role = "owner"
	Role_Read      Role = "read"
	Role_ReadWrite Role = "read_write"
	Role_None      Role = "none"
)

func (e *Role) Scan(value interface{}) error {
	var enumValue string
	switch val := value.(type) {
	case string:
		enumValue = val
	case []byte:
		enumValue = string(val)
	default:
		return errors.New("jet: Invalid scan value for AllTypesEnum enum. Enum value has to be of type string or []byte")
	}

	switch enumValue {
	case "owner":
		*e = Role_Owner
	case "read":
		*e = Role_Read
	case "read_write":
		*e = Role_ReadWrite
	case "none":
		*e = Role_None
	default:
		return errors.New("jet: Invalid scan value '" + enumValue + "' for Role enum")
	}

	return nil
}

func (e Role) String() string {
	return string(e)
}
