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

type Goal struct {
	ID          int64 `sql:"primary_key"`
	WorkspaceID int64
	UserID      int64
	Amount      float64
	Icon        string
	TargetDate  time.Time
	Status      GoalStatus
	Deadline    time.Time
}
