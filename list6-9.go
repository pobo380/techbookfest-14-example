package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
	"os"
)

func main() {
	// @<strong>{FileDescriptorSet を読み込む}
	bytes, _ := os.ReadFile("example_proto/fds.pb")
	fds := &descriptorpb.FileDescriptorSet{}
	proto.Unmarshal(bytes, fds)

	// @<strong>{registry を生成する}
	registry, _ := protodesc.NewFiles(fds)

	// @<strong>{名前を元に registry から Descriptor を取得する}
	desc, _ := registry.FindDescriptorByName("example_proto.ExampleMessage")

	// @<strong>{MessageDescriptor へキャストする}
	msgDesc := desc.(protoreflect.MessageDescriptor)

	// @<strong>{FieldDescriptor を取得する}
	nicknameField := msgDesc.Fields().ByName("nickname")

	// @<strong>{新しいメッセージを生成しフィールドに値をセットする}
	newMsg := dynamicpb.NewMessage(msgDesc)
	newMsg.Set(nicknameField, protoreflect.ValueOfString("dynamic!"))

	// @<strong>{メッセージにセットされた値を確認してみる}
	fmt.Println(newMsg) // → nickname:"dynamic!"
}
