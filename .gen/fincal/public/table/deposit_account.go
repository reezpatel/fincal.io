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

var DepositAccount = newDepositAccountTable("public", "deposit_account", "")

type depositAccountTable struct {
	postgres.Table

	// Columns
	ID             postgres.ColumnInteger
	WorkspaceID    postgres.ColumnInteger
	UserID         postgres.ColumnInteger
	CurrencyCode   postgres.ColumnString
	OpeningBalance postgres.ColumnFloat
	Name           postgres.ColumnString
	Icon           postgres.ColumnString
	Type           postgres.ColumnString
	Balance        postgres.ColumnFloat
	CreatedAt      postgres.ColumnTimestamp
	JSON           postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type DepositAccountTable struct {
	depositAccountTable

	EXCLUDED depositAccountTable
}

// AS creates new DepositAccountTable with assigned alias
func (a DepositAccountTable) AS(alias string) *DepositAccountTable {
	return newDepositAccountTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new DepositAccountTable with assigned schema name
func (a DepositAccountTable) FromSchema(schemaName string) *DepositAccountTable {
	return newDepositAccountTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new DepositAccountTable with assigned table prefix
func (a DepositAccountTable) WithPrefix(prefix string) *DepositAccountTable {
	return newDepositAccountTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new DepositAccountTable with assigned table suffix
func (a DepositAccountTable) WithSuffix(suffix string) *DepositAccountTable {
	return newDepositAccountTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newDepositAccountTable(schemaName, tableName, alias string) *DepositAccountTable {
	return &DepositAccountTable{
		depositAccountTable: newDepositAccountTableImpl(schemaName, tableName, alias),
		EXCLUDED:            newDepositAccountTableImpl("", "excluded", ""),
	}
}

func newDepositAccountTableImpl(schemaName, tableName, alias string) depositAccountTable {
	var (
		IDColumn             = postgres.IntegerColumn("id")
		WorkspaceIDColumn    = postgres.IntegerColumn("workspace_id")
		UserIDColumn         = postgres.IntegerColumn("user_id")
		CurrencyCodeColumn   = postgres.StringColumn("currency_code")
		OpeningBalanceColumn = postgres.FloatColumn("opening_balance")
		NameColumn           = postgres.StringColumn("name")
		IconColumn           = postgres.StringColumn("icon")
		TypeColumn           = postgres.StringColumn("type")
		BalanceColumn        = postgres.FloatColumn("balance")
		CreatedAtColumn      = postgres.TimestampColumn("created_at")
		JSONColumn           = postgres.StringColumn("json")
		allColumns           = postgres.ColumnList{IDColumn, WorkspaceIDColumn, UserIDColumn, CurrencyCodeColumn, OpeningBalanceColumn, NameColumn, IconColumn, TypeColumn, BalanceColumn, CreatedAtColumn, JSONColumn}
		mutableColumns       = postgres.ColumnList{WorkspaceIDColumn, UserIDColumn, CurrencyCodeColumn, OpeningBalanceColumn, NameColumn, IconColumn, TypeColumn, BalanceColumn, CreatedAtColumn, JSONColumn}
	)

	return depositAccountTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:             IDColumn,
		WorkspaceID:    WorkspaceIDColumn,
		UserID:         UserIDColumn,
		CurrencyCode:   CurrencyCodeColumn,
		OpeningBalance: OpeningBalanceColumn,
		Name:           NameColumn,
		Icon:           IconColumn,
		Type:           TypeColumn,
		Balance:        BalanceColumn,
		CreatedAt:      CreatedAtColumn,
		JSON:           JSONColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}