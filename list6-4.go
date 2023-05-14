package main

import (
	"fmt"
	"github.com/pobo380/techbookfest-14-example/example_proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func main() {
	// @<strong>{初期値を設定しておく}
	msg := &example_proto.ExampleMessage{}

	// @<strong>{protoreflect.Message と Descriptor の取得}
	refMsg := msg.ProtoReflect()
	desc := refMsg.Descriptor()

	// @<strong>{すべてのフィールドを走査する}
	for i := 0; i < desc.Fields().Len(); i++ {
		// @<strong>{フィールド情報を取得}
		field := desc.Fields().Get(i)

		// @<strong>{string 型のフィールドかどうかを判定して値をセットする}
		if field.Kind() == protoreflect.StringKind {
			refMsg.Set(field, protoreflect.ValueOfString("Go"))
		}
	}

	// @<strong>{ExampleMessage の内容が変更されているか出力してみる}
	fmt.Println(msg) // → nickname:"Go" address:"Go"
}
