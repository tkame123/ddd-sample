order.install:
	go install github.com/golang/mock/mockgen
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install entgo.io/contrib/entproto/cmd/protoc-gen-entgrpc@master

order.test.mock:
	mockgen -source=app/order_api/domain/port/domain_event/publisher.go -destination=app/order_api/domain/port/mock/domain_event/publisher.go -package=mock
	mockgen -source=app/order_api/domain/port/external_service/billing_api.go -destination=app/order_api/domain/port/mock/external_service/billing_api.go -package=mock
	mockgen -source=app/order_api/domain/port/external_service/ticket_api.go -destination=app/order_api/domain/port/mock/external_service/ticket_api.go -package=mock
	mockgen -source=app/order_api/domain/port/repository/repository.go -destination=app/order_api/domain/port/mock/repository/repository.go -package=mock
	mockgen -source=app/order_api/domain/port/repository/create_order_saga_state.go -destination=app/order_api/domain/port/mock/repository/create_order_saga_state.go -package=mock
	mockgen -source=app/order_api/domain/port/repository/order.go -destination=app/order_api/domain/port/mock/repository/order.go -package=mock
	mockgen -source=app/order_api/domain/port/service/create_order.go -destination=app/order_api/domain/port/mock/service/create_order.go -package=mock

order.ent.generate:
	go generate ./app/order_api/...

order.ent.schema:
	go run -mod=mod entgo.io/ent/cmd/ent describe ./app/order_api/adapter/database/ent/schema

TARGET=hoge
order.ent.add:
	go run -mod=mod entgo.io/ent/cmd/ent new --target app/order_api/adapter/database/ent/schema $(TARGET)

order.ent.atlas:
	 atlas schema inspect \
      -u "ent://app/order_api/adapter/database/ent/schema" \
      --dev-url "sqlite://file?mode=memory&_fk=1" \
      -w
