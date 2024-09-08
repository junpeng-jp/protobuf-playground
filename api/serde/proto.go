package serde

import (
	"fmt"

	"github.com/junpeng.ong/protobuf-playground/api/dto"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
)

func UnmarshalJsonToLocation(body []byte) ([]byte, error) {
	var locationDTO dto.Location
	err := protojson.Unmarshal(body, &locationDTO)

	if err != nil {
		return []byte{}, fmt.Errorf("Could not unmarshal request body as a dto.Location")
	}

	text, err := prototext.Marshal(&locationDTO)

	if err != nil {
		return []byte{}, fmt.Errorf("Failed to marshal dto.Location object to text")
	}

	return text, nil
}

func UnmarshalJsonToError(body []byte) ([]byte, error) {
	var errorDTO dto.Error
	err := protojson.Unmarshal(body, &errorDTO)

	if err != nil {
		return []byte{}, fmt.Errorf("Could not unmarshal request body as a dto.Error")
	}

	text, err := prototext.Marshal(&errorDTO)

	if err != nil {
		return []byte{}, fmt.Errorf("Failed to marshal dto.Error object to text")
	}

	return text, nil
}

func UnmarshalJsonToContainer(body []byte) ([]byte, error) {
	var containerDTO dto.Container
	err := protojson.Unmarshal(body, &containerDTO)

	if err != nil {
		return []byte{}, fmt.Errorf("Could not unmarshal request body as a dto.Container")
	}

	text, err := prototext.Marshal(&containerDTO)

	if err != nil {
		return []byte{}, fmt.Errorf("Failed to marshal dto.Container object to text")
	}

	return text, nil
}
