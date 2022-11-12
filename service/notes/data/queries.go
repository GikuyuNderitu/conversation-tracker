package data

import "fmt"

var (
	todoQuery = fmt.Sprintf("SELECT * FROM %s WHERE id = $id", todoTable)
)
