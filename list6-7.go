package main

import (
	"fmt"
	"github.com/pobo380/techbookfest-14-example/example_proto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func main() {
	// @<strong>{初期値を設定しておく}
	msg := &example_proto.ExampleMessage{Nickname: "John", Address: "Roppongi"}

	refMsg := msg.ProtoReflect()
	desc := refMsg.Descriptor()

	// @<strong>{すべてのフィールドを走査する}
	for i := 0; i < desc.Fields().Len(); i++ {
		// @<strong>{フィールド情報を取得}
		field := desc.Fields().Get(i)

		// @<strong>{カスタムオプションを取得}
		sensitive := proto.GetExtension(
			field.Options().(*descriptorpb.FieldOptions),
			example_proto.E_Sensitive).(bool)

		// @<strong>{カスタムオプションがセットされていたら値をクリア}
		if sensitive {
			refMsg.Clear(field)
		}
	}

	// @<strong>{ExampleMessage の内容が変更されているか出力してみる}
	fmt.Println(msg) // → nickname:"John"
}
