package main

import (
	"fmt"
	"github.com/pobo380/techbookfest-14-example/example_proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func main() {
	// @<strong>{オリジナルのメッセージの初期化}
	originalMsg := &example_proto.ExampleMessage{Nickname: "Hamachi"}

	// @<strong>{新しいメッセージの生成}
	refMsg := originalMsg.ProtoReflect().New()
	desc := refMsg.Descriptor()

	// @<strong>{新しいメッセージに値をセット}
	refMsg.Set(
		desc.Fields().ByName("nickname"),
		protoreflect.ValueOfString("Buri"))

	// @<strong>{Interface メソッドで通常のメッセージのインターフェースを変換}
	newMsg := refMsg.Interface()

	// @<strong>{それぞれ異なる値がセットされていることを確認}
	fmt.Println(originalMsg) // → nickname:"Hamachi"
	fmt.Println(newMsg)      // → nickname:"Buri"

	// @<strong>{Interface メソッドで取得した proto.Message 型はダウンキャストが可能}
	fmt.Println(newMsg.(*example_proto.ExampleMessage).Nickname) // → Buri
}
