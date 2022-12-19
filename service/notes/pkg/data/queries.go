package data

import (
	"fmt"
)

var (
	todoQuery  = fmt.Sprintf("SELECT * FROM %s WHERE id = $id", todoTable)
	convoQuery = fmt.Sprintf("SELECT * FROM %s WHERE id = $id", convoTable)
)
