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

type FileUpload struct {
	ID          int64 `sql:"primary_key"`
	Path        string
	WorkspaceID int64
	Name        string
	MimeType    string
	CreatedAt   time.Time
	Linked      bool
}