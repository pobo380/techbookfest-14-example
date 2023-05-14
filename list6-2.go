package main

import (
	"fmt"
	"github.com/pobo380/techbookfest-14-example/example_proto"
)

func main() {
	// @<strong>{protoc-gen-go の生成した Message struct を取得}
	msg := &example_proto.ExampleMessage{}

	// @<strong>{Message struct から protoreflect.Message を取得}
	refMsg := msg.ProtoReflect()

	// @<strong>{protoreflect.Message から Descriptor を取得}
	desc := refMsg.Descriptor()

	// @<strong>{型名を出力してみる}
	fmt.Println(desc.Name()) // → ExampleMessage
}
