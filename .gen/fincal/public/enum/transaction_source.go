//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package enum

import "github.com/go-jet/jet/v2/postgres"

var TransactionSource = &struct {
	DepositAccount  postgres.StringExpression
	CreditCard      postgres.StringExpression
	SecurityAccount postgres.StringExpression
	Entity          postgres.StringExpression
	Asset           postgres.StringExpression
	Loan            postgres.StringExpression
}{
	DepositAccount:  postgres.NewEnumValue("deposit_account"),
	CreditCard:      postgres.NewEnumValue("credit_card"),
	SecurityAccount: postgres.NewEnumValue("security_account"),
	Entity:          postgres.NewEnumValue("entity"),
	Asset:           postgres.NewEnumValue("asset"),
	Loan:            postgres.NewEnumValue("loan"),
}