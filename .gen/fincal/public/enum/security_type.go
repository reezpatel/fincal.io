//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package enum

import "github.com/go-jet/jet/v2/postgres"

var SecurityType = &struct {
	Stock      postgres.StringExpression
	MutualFund postgres.StringExpression
	Bond       postgres.StringExpression
	Etf        postgres.StringExpression
	Other      postgres.StringExpression
}{
	Stock:      postgres.NewEnumValue("stock"),
	MutualFund: postgres.NewEnumValue("mutual_fund"),
	Bond:       postgres.NewEnumValue("bond"),
	Etf:        postgres.NewEnumValue("etf"),
	Other:      postgres.NewEnumValue("other"),
}