package database // import "gnorm.org/gnorm/database"

// Info is the collection of schema info from a database.
type Info struct {
	Schemas []Schema // the list of schema info
}

// Schema is the information on a single named schema in the database.
type Schema struct {
	Name   string  // the name of this schema
	Tables []Table // the list of tables in this schema
	Enums  []Enum  // the list of enums in this schema
}

// Enum represents a type that has a set of allowed values.
type Enum struct {
	Schema string      // the name of the schema this enum is in
	Name   string      // the name of the enum type
	Values []EnumValue // the list of possible values for this enum
}

// EnumValue is one of the named values for an enum.
type EnumValue struct {
	Name  string // the label for this enum value
	Value int    // the value for this enum value (order)
}

// Table contains the definiiton of a database table.
type Table struct {
	Schema  string   // the name of the schema this table is in
	Name    string   // the name of the table
	Columns []Column // ordered list of columns in this table
}

// Column contains data about a column in a table.
type Column struct {
	Name        string      // the name of the column
	Type        string      // the mapped type of the column
	DBType      string      // the original type of the column in the db
	IsArray     bool        // true if the column type is an array
	Length      int         // non-zero if the type has a length (e.g. varchar[16])
	UserDefined bool        // true if the type is user-defined
	Nullable    bool        // true if the column is not NON NULL
	HasDefault  bool        // true if the column has a default
	Orig        interface{} // the raw database column data
}
