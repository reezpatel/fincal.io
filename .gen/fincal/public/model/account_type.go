//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import "errors"

type AccountType string

const (
	AccountType_Cash   AccountType = "cash"
	AccountType_Wallet AccountType = "wallet"
	AccountType_Bank   AccountType = "bank"
)

func (e *AccountType) Scan(value interface{}) error {
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
	case "cash":
		*e = AccountType_Cash
	case "wallet":
		*e = AccountType_Wallet
	case "bank":
		*e = AccountType_Bank
	default:
		return errors.New("jet: Invalid scan value '" + enumValue + "' for AccountType enum")
	}

	return nil
}

func (e AccountType) String() string {
	return string(e)
}