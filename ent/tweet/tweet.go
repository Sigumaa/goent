// Code generated by entc, DO NOT EDIT.

package tweet

const (
	// Label holds the string label denoting the tweet type in the database.
	Label = "tweet"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// Table holds the table name of the tweet in the database.
	Table = "tweets"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "tweets"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "user_tweets"
)

// Columns holds all SQL columns for tweet fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldContent,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tweets"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_tweets",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}