//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import "errors"

type RecurringType string

const (
	RecurringType_Subscription RecurringType = "subscription"
	RecurringType_Bills        RecurringType = "bills"
	RecurringType_Income       RecurringType = "income"
	RecurringType_Expense      RecurringType = "expense"
)

func (e *RecurringType) Scan(value interface{}) error {
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
	case "subscription":
		*e = RecurringType_Subscription
	case "bills":
		*e = RecurringType_Bills
	case "income":
		*e = RecurringType_Income
	case "expense":
		*e = RecurringType_Expense
	default:
		return errors.New("jet: Invalid scan value '" + enumValue + "' for RecurringType enum")
	}

	return nil
}

func (e RecurringType) String() string {
	return string(e)
}