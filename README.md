Sagaの実装の練習かねて、色々と書いてみて気付いた点を残しておくリポジトリ

以下に記載されている例を参考に書いてみる

https://amzn.asia/d/8Obsjx7

https://amzn.asia/d/f623yoI

Sagaの実装

https://github.com/tkame123/ddd-sample/tree/main/app/order_api/domain/service

# Saga

MicroServiceにおいては外部サービスの状況を信頼出来ない（コントロール出来ない）

各サービス毎に自治がおこなうことでのメリットを活かす為

よってサービス間のやり取りをメッセージを使い非同期にする

各サービスはメッセージの発行までを責務として、そのメッセージを受け取った側がいつ処理をするかは、受け取り側が決めることである

このような非同期処理の中で、複数のサービスが連携して一つのビジネスロジックを実現する必要がある場合の手法が大きく２つ存在する

- Choreography
- Saga

Choreographyは、各サービスが自身の状態を管理し、他のサービスとの連携を行うことでビジネスロジックを実現する手法である

Sagaは特定の統率者が、状態を管理し行うべきアクションをコントロールする手法である

個人的な感触では、データのレプリケート等、一連のアクションに関係性が低ければChoreographyも可能だが、ビジネストランザクションへの対応の場合は管理面からSaga一択となる印象である

これは現実的にサービス間に散らした知識を管理する事は出来ないと感じるため

## Sagaのポイント

### 全体概要

- FSMをつかって管理するのが良い
- EventによりStateが変化する
- Eventにより処理を依頼する
- Eventにより結果を受けとる

例

![image](https://github.com/user-attachments/assets/8081ef92-475e-49a5-857f-2c514d8b8bd4)
  
### 設計ポイント

**処理のトランザクションを分類する**

- 補償可能トランザクション -  補償トランザクションでロールバックされる可能性があるトランザクション（例：在庫切れで注文を受けれない状態だった際に元のOrderをキャンセルする等）
- ピポットトランザクション -  処理の最終決定の判断のポイントとなるトランザクション（例だと支払い処理　-> 支払い不可能であれば何があってもこのSagaは実行完了まで行き着かない）
- 再試行可能トランザクション - 成功を前提とし再試行により解決するもの

再試行可能トランザクションに無駄な手厚いケアを入れるのは必要ない等

この考えにより、FSMで表現したFlowを明確にシンプルに設計できると学べる

**カウンターメジャー**

- semantic lock（セマンティックロック）
- commutative update(交換可能な更新）
- Passimistic view(悲観的天格ぃ）
- Reread vallue(値の読み直し）
- Version FIle
- By Value

ひとまずは、、semantic lockをいれるのは必須かなあ。。

簡単にいえば、Saga実施中に変更を受けてほしくない対象をロックするような形

例；Orderキャンセル時には、キャンセル実施中として、競合する処理が実施されないようにする　等

これは、Microservice文脈でなくても、Sagaとして処理を切り離して別に管理する間、に作業中という状況で対象を保護するのは有効な手法かなと思う

# 構成

## 認証・認可

IdaasにAuth0をつかった、SPAにおけるOIDC構成とする

以下理解を整頓

### 全体概要

- auth0を認可サーバ、BEをRPとしたOIDC構成
- 利用者の本人確認はauth0の責務となり、結果としてIdTokenとAccessTokenを利用者に発行する
- RPとなるBEは認可サーバに認証を委任するため、認可サーバが発行したAccessTokenの検証でBEの認証行為とする

### 認証

BEが行う認証行為はAccessTokenの検証となる

- issuerがAuth0の想定するテナントである点
- 署名がissuerによるものである点
- 有効期限の確認
- audienceが自分を指す点
- 上記確認がIdTokenでは出来ない（idTokenのAudienceは利用者を指すのがOIDCの推奨定義の為）ので、BEはAccessTokenを検証する

参考

https://auth0.com/blog/id-token-access-token-what-is-the-difference/

https://qiita.com/TakahikoKawasaki/items/8f0e422c7edd2d220e06

https://qiita.com/TakahikoKawasaki/items/970548727761f9e02bcd


実装は、Auth0が提供するライブラリを利用

https://github.com/auth0/go-jwt-middleware

### 認可

- Auth0のRBAC機能を利用してRBACベースで実装する
- Auth0のAPI単位でPermissionを定義する
- Roleに対してPermissionを設定する
- AccessTokenにPermissionをClaim追加する設定を行う
- BE側はAccessTokenから取り出したpermissionを評価する

Permissionの評価はCasbinを利用

- Modelのサンプルも豊富で、柔軟な変更もできそうでいい感じ

具体的には以下とした

```
sub: だれが 
obj: 何に 
act: 何をする
```

- Permission を　{act} : {sub} と定義
- objを protoのService単位とする

Auth0 Permissio Sample
```
read:orders
write:orders
```

Casbin Policy Sample
```
p, orders, /order_api.v1.OrderService/FindOrder, write
p, orders, /order_api.v1.OrderService/FindOrder, read
p, orders, /order_api.v1.OrderService/CreateOrder, write
```

参考

https://auth0.com/docs/get-started/apis/enable-role-based-access-control-for-apis

https://casbin.org/docs/supported-models

### その他

SPA側でのアクセストークンの保存先をローカルストレージでなくメモリへという点
Auth0のSDKだと対応している模様

https://qiita.com/kura_lab/items/8eda9b2899e00e95a50c

Cognitoはこの点、AWS側のSDKがLocalStorageに格納して定期的に話題にあがっている理解

https://github.com/aws-amplify/amplify-js/issues/3436

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

### API: Connect

gRPCとgrpc gateewayを自分で対応するより全然便利。。。

https://connectrpc.com/

### FSM: looplab/fsm

https://github.com/looplab/fsm

sagaの状態管理につかってみたけど、機能が方法で使いやすい


## KitchenAPI

### ORM: 未定

ent.以外をためしてみる予定

