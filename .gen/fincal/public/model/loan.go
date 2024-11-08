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

type Loan struct {
	ID                 int64 `sql:"primary_key"`
	WorkspaceID        int64
	UserID             int64
	Type               string
	InterestRate       float64
	TermMonths         int32
	LoanAmount         float64
	PrincipalAmount    float64
	LoanBalance        float64
	PrincipalBalance   float64
	MonthlyInstallment float64
	StartDate          time.Time
	Status             LoanStatus
}
