//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type CreditCardAccount struct {
	ID           int64 `sql:"primary_key"`
	UserID       int64
	WorkspaceID  int64
	CurrencyCode string
	Balance      float64
	CreditLimit  float64
	InterestRate float64
}