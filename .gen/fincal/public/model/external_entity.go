//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type ExternalEntity struct {
	ID          int64 `sql:"primary_key"`
	WorkspaceID int64
	Name        string
	Icon        string
}
