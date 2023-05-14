package main

import (
	"fmt"
	"github.com/pobo380/techbookfest-14-example/example_proto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func Censor(refMsg protoreflect.Message) {
	desc := refMsg.Descriptor()

	// @<strong>{すべてのフィールドを走査する}
	for i := 0; i < desc.Fields().Len(); i++ {
		// @<strong>{フィールド情報を取得}
		field := desc.Fields().Get(i)

		// @<strong>{フィールドの Value を取得}
		fieldVal := refMsg.Get(field)

		// @<strong>{map 型の場合の処理}
		if field.IsMap() {
			// @<strong>{map の value の型がメッセージ型の場合に再帰呼び出しをする}
			valField := field.MapValue()
			if valField.Kind() == protoreflect.MessageKind {
				// @<strong>{map の各要素に対して再帰呼び出し}
				fieldVal.Map().Range(
					func(_ protoreflect.MapKey,
						val protoreflect.Value) bool {
						Censor(val.Message())
						return true
					})
			}

			// @<strong>{map フィールド自体の削除処理はサポートしないことにする}
			continue
		}

		// @<strong>{Message 型の場合の処理}
		if field.Kind() == protoreflect.MessageKind {
			// @<strong>{List かどうかを確認}
			if field.IsList() {
				// @<strong>{List としての値を取得}
				list := fieldVal.List()

				// @<strong>{List のすべて要素を走査}
				for j := 0; j < list.Len(); j++ {
					// @<strong>{List の要素の値を取得}
					listVal := list.Get(j)

					// @<strong>{Message に対して再帰呼び出し}
					Censor(listVal.Message())
				}
			} else {
				// @<strong>{Message に対して再帰呼び出し}
				Censor(fieldVal.Message())
			}

			// @<strong>{削除処理はスキップ}
			continue
		}

		// @<strong>{カスタムオプションを取得}
		sensitive := proto.GetExtension(
			field.Options().(*descriptorpb.FieldOptions),
			example_proto.E_Sensitive).(bool)

		// @<strong>{カスタムオプションがセットされていたら値をクリア}
		if sensitive {
			refMsg.Clear(field)
		}
	}
}

func main() {
	// @<strong>{初期値を設定しておく}
	msg := &example_proto.AdvancedExampleMessage{
		Child: &example_proto.ExampleMessage{
			Nickname: "A",
			Address:  "Roppongi",
		},
		Children: []*example_proto.ExampleMessage{
			{
				Nickname: "B",
				Address:  "Shibuya",
			},
			{
				Nickname: "C",
				Address:  "Meguro",
			},
		},
		MapField: map[string]*example_proto.ExampleMessage{
			"key-1": {
				Nickname: "D",
				Address:  "Shinjuku",
			},
		},
	}

	// @<strong>{カスタムオプションが有効なフィールドの値の削除を実行}
	Censor(msg.ProtoReflect())

	// @<strong>{ExampleMessage の内容が変更されているか出力してみる}
	fmt.Println(msg)
	// → child:{nickname:"A"}
	//   children:{nickname:"B"}
	//   children:{nickname:"C"}
	//   map_field:{key:"key-1" value:{nickname:"D"}}
}
