package serializer

import (
	"encoding/json"
	"fmt"
	"grpc/pb"
	"os"
)

func WriteProtobufToBinaryFile(message *pb.Laptop, filename string) error {
	// data, err := proto.Marshal(message)
	data, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to binary: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot write binary data to file: %w", err)
	}

	return nil
}

func ReadProtobufFromBinaryFile(filename string, message *pb.Laptop) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary data from file: %w", err)
	}

	err = json.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal binary to proto message: %w", err)
	}

	return nil
}
