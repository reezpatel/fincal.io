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

var RecurringTransaction = newRecurringTransactionTable("public", "recurring_transaction", "")

type recurringTransactionTable struct {
	postgres.Table

	// Columns
	ID             postgres.ColumnInteger
	CategoryID     postgres.ColumnInteger
	UserID         postgres.ColumnInteger
	Amount         postgres.ColumnFloat
	FromID         postgres.ColumnInteger
	ToID           postgres.ColumnInteger
	Source         postgres.ColumnString
	Dest           postgres.ColumnString
	Frequency      postgres.ColumnString
	NextOccurrence postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type RecurringTransactionTable struct {
	recurringTransactionTable

	EXCLUDED recurringTransactionTable
}

// AS creates new RecurringTransactionTable with assigned alias
func (a RecurringTransactionTable) AS(alias string) *RecurringTransactionTable {
	return newRecurringTransactionTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new RecurringTransactionTable with assigned schema name
func (a RecurringTransactionTable) FromSchema(schemaName string) *RecurringTransactionTable {
	return newRecurringTransactionTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new RecurringTransactionTable with assigned table prefix
func (a RecurringTransactionTable) WithPrefix(prefix string) *RecurringTransactionTable {
	return newRecurringTransactionTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new RecurringTransactionTable with assigned table suffix
func (a RecurringTransactionTable) WithSuffix(suffix string) *RecurringTransactionTable {
	return newRecurringTransactionTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newRecurringTransactionTable(schemaName, tableName, alias string) *RecurringTransactionTable {
	return &RecurringTransactionTable{
		recurringTransactionTable: newRecurringTransactionTableImpl(schemaName, tableName, alias),
		EXCLUDED:                  newRecurringTransactionTableImpl("", "excluded", ""),
	}
}

func newRecurringTransactionTableImpl(schemaName, tableName, alias string) recurringTransactionTable {
	var (
		IDColumn             = postgres.IntegerColumn("id")
		CategoryIDColumn     = postgres.IntegerColumn("category_id")
		UserIDColumn         = postgres.IntegerColumn("user_id")
		AmountColumn         = postgres.FloatColumn("amount")
		FromIDColumn         = postgres.IntegerColumn("from_id")
		ToIDColumn           = postgres.IntegerColumn("to_id")
		SourceColumn         = postgres.StringColumn("source")
		DestColumn           = postgres.StringColumn("dest")
		FrequencyColumn      = postgres.StringColumn("frequency")
		NextOccurrenceColumn = postgres.TimestampzColumn("next_occurrence")
		allColumns           = postgres.ColumnList{IDColumn, CategoryIDColumn, UserIDColumn, AmountColumn, FromIDColumn, ToIDColumn, SourceColumn, DestColumn, FrequencyColumn, NextOccurrenceColumn}
		mutableColumns       = postgres.ColumnList{CategoryIDColumn, UserIDColumn, AmountColumn, FromIDColumn, ToIDColumn, SourceColumn, DestColumn, FrequencyColumn, NextOccurrenceColumn}
	)

	return recurringTransactionTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:             IDColumn,
		CategoryID:     CategoryIDColumn,
		UserID:         UserIDColumn,
		Amount:         AmountColumn,
		FromID:         FromIDColumn,
		ToID:           ToIDColumn,
		Source:         SourceColumn,
		Dest:           DestColumn,
		Frequency:      FrequencyColumn,
		NextOccurrence: NextOccurrenceColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}