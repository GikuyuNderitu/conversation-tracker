package data

import (
	"errors"

	service_pb "atypicaldev.com/conversation/notes/internal/proto/service/v1"
)

var (
	ErrConvoRequestValidationError = errors.New("title must be populated")
)

func validateConvoRequest(req *service_pb.CreateConversationRequest) error {
	if req.GetTitle() == "" {
		return ErrConvoRequestValidationError
	}
	return nil
}
