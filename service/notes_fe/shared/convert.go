package shared

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ToJson(p protoreflect.ProtoMessage) (string, error) {
	bytes, err := protojson.Marshal(p)
	if err != nil {
		return "	", err
	}

	return string(bytes), nil
}
