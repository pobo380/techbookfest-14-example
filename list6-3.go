package main

import (
	"fmt"
	"github.com/pobo380/techbookfest-14-example/example_proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func main() {
	// @<strong>{初期値を設定しておく}
	msg := &example_proto.ExampleMessage{UserId: 1, Nickname: "John"}

	// @<strong>{protoreflect.Message と Descriptor の取得}
	refMsg := msg.ProtoReflect()
	desc := refMsg.Descriptor()

	// @<strong>{フィールド名からフィールド情報を取得}
	userIdField := desc.Fields().ByName("user_id")

	// @<strong>{フィールド番号からフィールド情報を取得}
	nicknameField := desc.Fields().ByNumber(protoreflect.FieldNumber(2))

	// @<strong>{protoreflect.Message の Get メソッドで値を取得}
	userIdVal := refMsg.Get(userIdField)
	nicknameVal := refMsg.Get(nicknameField)

	// @<strong>{新しくセットする値を生成}
	newUserId := protoreflect.ValueOfUint64(userIdVal.Uint() + 99)
	newNickname := protoreflect.ValueOfString(nicknameVal.String() + "Manjiro")

	// @<strong>{protoreflect.Message の Set メソッドにフィールド情報を渡して値をセット}
	refMsg.Set(userIdField, newUserId)
	refMsg.Set(nicknameField, newNickname)

	// @<strong>{ExampleMessage の内容が変更されているか出力してみる}
	fmt.Println(msg) // → user_id:100 nickname:"JohnManjiro"
}
