//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type DepositAccount struct {
	ID             int64 `sql:"primary_key"`
	WorkspaceID    int64
	UserID         int64
	CurrencyCode   string
	OpeningBalance float64
	Name           string
	Icon           string
	Type           AccountType
	Balance        float64
	CreatedAt      time.Time
	JSON           string
}
