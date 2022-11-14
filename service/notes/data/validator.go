package data

import (
	"errors"

	pb "atypicaldev.com/conversation/api/notes"
)

var (
	ErrConvoRequestValidationError = errors.New("title must be populated")
)

func validateConvoRequest(req *pb.CreateConversationRequest) error {
	if req.GetTitle() == "" {
		return ErrConvoRequestValidationError
	}
	return nil
}
