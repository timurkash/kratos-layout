// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// KratosColumns holds the columns for the "kratos" table.
	KratosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// KratosTable holds the schema information for the "kratos" table.
	KratosTable = &schema.Table{
		Name:       "kratos",
		Columns:    KratosColumns,
		PrimaryKey: []*schema.Column{KratosColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		KratosTable,
	}
)

func init() {
}
