//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type SecurityAccount struct {
	ID          int64 `sql:"primary_key"`
	UserID      int64
	WorkspaceID int64
	MarketID    string
	Name        string
}