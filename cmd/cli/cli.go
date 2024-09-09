package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/junpeng.ong/protobuf-playground/registry"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

func main() {
	filepath := os.Args[1]
	messageType := protoreflect.FullName(os.Args[2])
	if !messageType.IsValid() {
		log.Fatalf("message type %q is not a valid fully-qualified type name\n", messageType)
	}

	// Initialize a new proto registry from compiled FileDescriptorSet
	registry := registry.NewProtoRegistry(filepath)

	// Retrieve a descriptor from the registry by name
	descriptor, err := registry.FindDescriptorByName(messageType)
	if err != nil {
		log.Fatalf("failed to find message type %q in given descriptors: %v\n", messageType, err)
	}
	// Assert that descriptor is a MessageDescriptor
	messageDescriptor, ok := descriptor.(protoreflect.MessageDescriptor)
	if !ok {
		log.Fatalf("element named %q is not a message (%T)\n", messageType, descriptor)
	}
	message := dynamicpb.NewMessage(messageDescriptor)

	// Read json data from stdin and unmarshal to the dynamic message
	messageData, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read message data from stdin: %v\n", err)
	}

	if err := protojson.Unmarshal(messageData, message); err != nil {
		log.Fatalf("failed to process input data for message type %q: %v\n", messageType, err)
	}

	// And write text format to stdout
	fmt.Print(prototext.Format(message))
}
