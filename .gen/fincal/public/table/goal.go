//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Goal = newGoalTable("public", "goal", "")

type goalTable struct {
	postgres.Table

	// Columns
	ID          postgres.ColumnInteger
	WorkspaceID postgres.ColumnInteger
	UserID      postgres.ColumnInteger
	Amount      postgres.ColumnFloat
	Icon        postgres.ColumnString
	TargetDate  postgres.ColumnTimestampz
	Status      postgres.ColumnString
	Deadline    postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type GoalTable struct {
	goalTable

	EXCLUDED goalTable
}

// AS creates new GoalTable with assigned alias
func (a GoalTable) AS(alias string) *GoalTable {
	return newGoalTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new GoalTable with assigned schema name
func (a GoalTable) FromSchema(schemaName string) *GoalTable {
	return newGoalTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new GoalTable with assigned table prefix
func (a GoalTable) WithPrefix(prefix string) *GoalTable {
	return newGoalTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new GoalTable with assigned table suffix
func (a GoalTable) WithSuffix(suffix string) *GoalTable {
	return newGoalTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newGoalTable(schemaName, tableName, alias string) *GoalTable {
	return &GoalTable{
		goalTable: newGoalTableImpl(schemaName, tableName, alias),
		EXCLUDED:  newGoalTableImpl("", "excluded", ""),
	}
}

func newGoalTableImpl(schemaName, tableName, alias string) goalTable {
	var (
		IDColumn          = postgres.IntegerColumn("id")
		WorkspaceIDColumn = postgres.IntegerColumn("workspace_id")
		UserIDColumn      = postgres.IntegerColumn("user_id")
		AmountColumn      = postgres.FloatColumn("amount")
		IconColumn        = postgres.StringColumn("icon")
		TargetDateColumn  = postgres.TimestampzColumn("target_date")
		StatusColumn      = postgres.StringColumn("status")
		DeadlineColumn    = postgres.TimestampzColumn("deadline")
		allColumns        = postgres.ColumnList{IDColumn, WorkspaceIDColumn, UserIDColumn, AmountColumn, IconColumn, TargetDateColumn, StatusColumn, DeadlineColumn}
		mutableColumns    = postgres.ColumnList{WorkspaceIDColumn, UserIDColumn, AmountColumn, IconColumn, TargetDateColumn, StatusColumn, DeadlineColumn}
	)

	return goalTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		WorkspaceID: WorkspaceIDColumn,
		UserID:      UserIDColumn,
		Amount:      AmountColumn,
		Icon:        IconColumn,
		TargetDate:  TargetDateColumn,
		Status:      StatusColumn,
		Deadline:    DeadlineColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
