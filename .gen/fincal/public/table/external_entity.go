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

var ExternalEntity = newExternalEntityTable("public", "external_entity", "")

type externalEntityTable struct {
	postgres.Table

	// Columns
	ID          postgres.ColumnInteger
	WorkspaceID postgres.ColumnInteger
	Name        postgres.ColumnString
	Icon        postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type ExternalEntityTable struct {
	externalEntityTable

	EXCLUDED externalEntityTable
}

// AS creates new ExternalEntityTable with assigned alias
func (a ExternalEntityTable) AS(alias string) *ExternalEntityTable {
	return newExternalEntityTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ExternalEntityTable with assigned schema name
func (a ExternalEntityTable) FromSchema(schemaName string) *ExternalEntityTable {
	return newExternalEntityTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ExternalEntityTable with assigned table prefix
func (a ExternalEntityTable) WithPrefix(prefix string) *ExternalEntityTable {
	return newExternalEntityTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ExternalEntityTable with assigned table suffix
func (a ExternalEntityTable) WithSuffix(suffix string) *ExternalEntityTable {
	return newExternalEntityTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newExternalEntityTable(schemaName, tableName, alias string) *ExternalEntityTable {
	return &ExternalEntityTable{
		externalEntityTable: newExternalEntityTableImpl(schemaName, tableName, alias),
		EXCLUDED:            newExternalEntityTableImpl("", "excluded", ""),
	}
}

func newExternalEntityTableImpl(schemaName, tableName, alias string) externalEntityTable {
	var (
		IDColumn          = postgres.IntegerColumn("id")
		WorkspaceIDColumn = postgres.IntegerColumn("workspace_id")
		NameColumn        = postgres.StringColumn("name")
		IconColumn        = postgres.StringColumn("icon")
		allColumns        = postgres.ColumnList{IDColumn, WorkspaceIDColumn, NameColumn, IconColumn}
		mutableColumns    = postgres.ColumnList{WorkspaceIDColumn, NameColumn, IconColumn}
	)

	return externalEntityTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		WorkspaceID: WorkspaceIDColumn,
		Name:        NameColumn,
		Icon:        IconColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
