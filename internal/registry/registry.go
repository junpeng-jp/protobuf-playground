package registry

import (
	"fmt"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
)

// https://buf.build/docs/reference/descriptors#dynamic-messaging
// first compile the .proto files into its binary format
func NewProtoRegistry(filepath string) *protoregistry.Files {

	fmt.Printf("Reading binpb file from `%s`\n", filepath)
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Unable to read file from %s: %v\n", filepath, err)
	}

	// then unmarshal the binary into a FileDescriptorSet
	var files descriptorpb.FileDescriptorSet
	if err := proto.Unmarshal(data, &files); err != nil {
		log.Fatalf("failed to process descriptors in %s: %v\n", filepath, err)
	}

	descriptorRegistry, err := protodesc.NewFiles(&files)
	if err != nil {
		log.Fatalf("failed to parse fileDescriptors: %v\n", err)
	}

	return descriptorRegistry
}
