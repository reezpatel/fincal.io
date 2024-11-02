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

var User = newUserTable("public", "user", "")

type userTable struct {
	postgres.Table

	// Columns
	ID       postgres.ColumnInteger
	Email    postgres.ColumnString
	Icon     postgres.ColumnString
	Password postgres.ColumnString
	Name     postgres.ColumnString
	Hash     postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type UserTable struct {
	userTable

	EXCLUDED userTable
}

// AS creates new UserTable with assigned alias
func (a UserTable) AS(alias string) *UserTable {
	return newUserTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new UserTable with assigned schema name
func (a UserTable) FromSchema(schemaName string) *UserTable {
	return newUserTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new UserTable with assigned table prefix
func (a UserTable) WithPrefix(prefix string) *UserTable {
	return newUserTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new UserTable with assigned table suffix
func (a UserTable) WithSuffix(suffix string) *UserTable {
	return newUserTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newUserTable(schemaName, tableName, alias string) *UserTable {
	return &UserTable{
		userTable: newUserTableImpl(schemaName, tableName, alias),
		EXCLUDED:  newUserTableImpl("", "excluded", ""),
	}
}

func newUserTableImpl(schemaName, tableName, alias string) userTable {
	var (
		IDColumn       = postgres.IntegerColumn("id")
		EmailColumn    = postgres.StringColumn("email")
		IconColumn     = postgres.StringColumn("icon")
		PasswordColumn = postgres.StringColumn("password")
		NameColumn     = postgres.StringColumn("name")
		HashColumn     = postgres.StringColumn("hash")
		allColumns     = postgres.ColumnList{IDColumn, EmailColumn, IconColumn, PasswordColumn, NameColumn, HashColumn}
		mutableColumns = postgres.ColumnList{EmailColumn, IconColumn, PasswordColumn, NameColumn, HashColumn}
	)

	return userTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:       IDColumn,
		Email:    EmailColumn,
		Icon:     IconColumn,
		Password: PasswordColumn,
		Name:     NameColumn,
		Hash:     HashColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
