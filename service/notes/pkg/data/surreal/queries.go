package surreal

import (
	"fmt"
)

var (
	noteQuery           = fmt.Sprintf("SELECT * FROM %s WHERE id = $id", noteTable)
	convoQuery          = fmt.Sprintf("SELECT * FROM %s WHERE id = $id", convoTable)
	createNoteQuery     = fmt.Sprintf("CREATE %s SET content = $content, reply = $reply", noteTable)
	relateNoteWithConvo = "RELATE $note -> belongsto -> $convo"
)
