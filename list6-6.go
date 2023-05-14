package main

import (
	"fmt"
	"github.com/pobo380/techbookfest-14-example/example_proto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func main() {
	msg := &example_proto.ExampleMessage{}
	refMsg := msg.ProtoReflect()
	desc := refMsg.Descriptor()

	// @<strong>{フィールドのオプション情報を取得}
	fieldOpts := desc.Fields().ByName("address").Options()

	// @<strong>{カスタムオプションを取得}
	sensitive := proto.GetExtension(
		fieldOpts.(*descriptorpb.FieldOptions),
		example_proto.E_Sensitive).(bool)

	// @<strong>{カスタムオプションの値を確認}
	fmt.Println(sensitive) // → true
}
