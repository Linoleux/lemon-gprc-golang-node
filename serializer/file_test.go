package serializer_test

import (
	"grpc/pb"
	"grpc/sample"
	"grpc/serializer"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	file := "../tmp/laptop.bin"
	// json := "../tmp/laptop.json"
	laptop1 := sample.NewLaptop()

	err := serializer.WriteProtobufToBinaryFile(laptop1, file)
	require.NoError(t, err)

	// err = serializer.WriteProtobufToJSONFile(laptop1, json)
	// require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(file, laptop2)
	require.NoError(t, err)

	// require.True(t, proto.Equal(laptop1, laptop2))
}
