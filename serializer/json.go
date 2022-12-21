package serializer

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJSON(message proto.Message) (string, error) {
	/*
		marshaler := jsonpb.Marshaler{
			EnumsAsInts:  false,
			EmitDefaults: true,
			Indent:       "  ",
			OrigName:     true,
		}
	*/

	marshaler, err := protojson.Marshal(message)

	if err != nil {
		return "", err
	}

	return string(marshaler), nil
}

func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to JSON: %w", err)
	}

	err = os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("cannot write JSON data to file: %w", err)
	}

	return nil
}
