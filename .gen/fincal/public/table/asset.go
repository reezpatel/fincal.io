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

var Asset = newAssetTable("public", "asset", "")

type assetTable struct {
	postgres.Table

	// Columns
	ID            postgres.ColumnInteger
	WorkspaceID   postgres.ColumnInteger
	UserID        postgres.ColumnInteger
	Name          postgres.ColumnString
	Icon          postgres.ColumnString
	OriginalValue postgres.ColumnFloat
	CurrentValue  postgres.ColumnFloat
	AcquiredAt    postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type AssetTable struct {
	assetTable

	EXCLUDED assetTable
}

// AS creates new AssetTable with assigned alias
func (a AssetTable) AS(alias string) *AssetTable {
	return newAssetTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AssetTable with assigned schema name
func (a AssetTable) FromSchema(schemaName string) *AssetTable {
	return newAssetTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new AssetTable with assigned table prefix
func (a AssetTable) WithPrefix(prefix string) *AssetTable {
	return newAssetTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new AssetTable with assigned table suffix
func (a AssetTable) WithSuffix(suffix string) *AssetTable {
	return newAssetTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newAssetTable(schemaName, tableName, alias string) *AssetTable {
	return &AssetTable{
		assetTable: newAssetTableImpl(schemaName, tableName, alias),
		EXCLUDED:   newAssetTableImpl("", "excluded", ""),
	}
}

func newAssetTableImpl(schemaName, tableName, alias string) assetTable {
	var (
		IDColumn            = postgres.IntegerColumn("id")
		WorkspaceIDColumn   = postgres.IntegerColumn("workspace_id")
		UserIDColumn        = postgres.IntegerColumn("user_id")
		NameColumn          = postgres.StringColumn("name")
		IconColumn          = postgres.StringColumn("icon")
		OriginalValueColumn = postgres.FloatColumn("original_value")
		CurrentValueColumn  = postgres.FloatColumn("current_value")
		AcquiredAtColumn    = postgres.TimestampzColumn("acquired_at")
		allColumns          = postgres.ColumnList{IDColumn, WorkspaceIDColumn, UserIDColumn, NameColumn, IconColumn, OriginalValueColumn, CurrentValueColumn, AcquiredAtColumn}
		mutableColumns      = postgres.ColumnList{WorkspaceIDColumn, UserIDColumn, NameColumn, IconColumn, OriginalValueColumn, CurrentValueColumn, AcquiredAtColumn}
	)

	return assetTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:            IDColumn,
		WorkspaceID:   WorkspaceIDColumn,
		UserID:        UserIDColumn,
		Name:          NameColumn,
		Icon:          IconColumn,
		OriginalValue: OriginalValueColumn,
		CurrentValue:  CurrentValueColumn,
		AcquiredAt:    AcquiredAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
