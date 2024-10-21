Sagaの実装の練習かねて、色々と書いてみて気付いた点を残しておくリポジトリ

以下に記載されている例を参考に書いてみる

https://amzn.asia/d/8Obsjx7

https://amzn.asia/d/f623yoI

Sagaの実装

[https://github.com/tkame123/ddd-sample//app/order_api/domain/service/create_order_saga
](https://github.com/tkame123/ddd-sample/tree/main/app/order_api/domain/service)

# 構成

## Message

SQS/SNS

### スキーマ共有：Protobuf
送信時はひとまずProtoJsonでシリアライズ、デシリアライズを実装してみたが、結構使いやすい印象
https://protobuf.dev/programming-guides/proto3/#json

```
// すべてのMessageに付与する識別情報
message Subject {
  Type type = 1;
  Service source = 2;
  // MEMO:　冪等キー等もおそらくここに追加する
}

// 送信されるMessage
message Message {
  Subject subject = 1;
  google.protobuf.Any envelope = 2;
}
```

Any or oneOfを迷ったけど、Anyを試してみる

参考：ChatGPT
```
解説
Envelope メッセージは google.protobuf.Any 型を持ち、任意のメッセージを格納できます。
ptypes.UnmarshalAny を使って、Any 型から特定のメッセージ型 (LoginEvent や LogoutEvent) を動的にアンマーシャルし、型に応じて処理します。
もし Any 型のデータがアンマーシャルできない場合、他の型としての処理を試みます。
結論
Golang で protobuf のメッセージを型別に分岐させる方法としては、以下の2つのアプローチが有効です：

oneof を使用: 複数のメッセージ型が1つのメッセージ内に含まれる場合、型アサーションを使ってメッセージを分岐させる。
Any を使用: 動的なメッセージタイプを扱い、受信時に適切な型にアンマーシャリングして分岐させる。
使う場面に応じて、どちらの方法も柔軟に適用できます。
```

Sample OneOf
```
package main

import (
    "fmt"
    pb "path/to/protobuf/package"  // Protobufで生成されたコードのパス
)

func handleEvent(event *pb.Event) {
    // oneof フィールドの型に応じて分岐
    switch e := event.EventType.(type) {
    case *pb.Event_LoginEvent:
        fmt.Println("Login event for user:", e.LoginEvent.Username)
    case *pb.Event_LogoutEvent:
        fmt.Println("Logout event for user:", e.LogoutEvent.Username)
    default:
        fmt.Println("Unknown event type")
    }
}

func main() {
    // サンプルデータの作成
    loginEvent := &pb.Event{
        EventType: &pb.Event_LoginEvent{
            LoginEvent: &pb.LoginEvent{
                Username: "john_doe",
            },
        },
    }

    // イベントを処理
    handleEvent(loginEvent)

    logoutEvent := &pb.Event{
        EventType: &pb.Event_LogoutEvent{
            LogoutEvent: &pb.LogoutEvent{
                Username: "john_doe",
            },
        },
    }

    handleEvent(logoutEvent)
}
```

Sample Any
```
package main

import (
    "fmt"
    "github.com/golang/protobuf/ptypes"
    "github.com/golang/protobuf/ptypes/any"
    pb "path/to/protobuf/package"  // Protobufで生成されたコードのパス
)

func handleEnvelope(envelope *pb.Envelope) {
    var loginEvent pb.LoginEvent
    var logoutEvent pb.LogoutEvent

    // `Any` 型の `payload` フィールドを動的にアンマーシャルして型を判別
    if err := ptypes.UnmarshalAny(envelope.Payload, &loginEvent); err == nil {
        fmt.Println("Login event for user:", loginEvent.Username)
        return
    }

    if err := ptypes.UnmarshalAny(envelope.Payload, &logoutEvent); err == nil {
        fmt.Println("Logout event for user:", logoutEvent.Username)
        return
    }

    fmt.Println("Unknown event type")
}

func main() {
    // サンプルデータの作成
    loginEvent := pb.LoginEvent{Username: "john_doe"}
    anyLoginEvent, _ := ptypes.MarshalAny(&loginEvent)

    envelope := &pb.Envelope{
        Payload: anyLoginEvent,
    }

    // イベントを処理
    handleEnvelope(envelope)

    // LogoutEvent も同様に処理
    logoutEvent := pb.LogoutEvent{Username: "john_doe"}
    anyLogoutEvent, _ := ptypes.MarshalAny(&logoutEvent)

    envelope.Payload = anyLogoutEvent
    handleEnvelope(envelope)
}

```
## OrderAPI

### ORM: ent.

https://entgo.io/ja/

migrationやShemaのVisualize対応等含めて、ワンセットで便利な感じ

### ORM: Connect

gRPCとgrpc gateewayを自分で対応するより全然便利。。。

https://connectrpc.com/

### FSM: looplab/fsm

https://github.com/looplab/fsm

sagaの状態管理につかってみたけど、機能が方法で使いやすい


## KitchenAPI

### ORM: 未定

ent.以外をためしてみる予定

