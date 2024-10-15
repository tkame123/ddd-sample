
tool.install:
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/golang/mock/mockgen
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install entgo.io/contrib/entproto/cmd/protoc-gen-entgrpc@master
	go install github.com/air-verse/air@latest

localstack.start:
	DOCKER_FLAGS='-v $(CURDIR)/dev/locakstack-init-aws.sh:/etc/localstack/init/ready.d/init-aws.sh' localstack start
	#DOCKER_FLAGS='-v $(CURDIR)/dev/locakstack-init-aws.sh:/etc/localstack/init/ready.d/init-aws.sh' localstack start -d

localstack.stop:
	localstack stop

buf.generate:
	buf generate

buf.lint:
	buf lint

order.wire:
	cd app/order_api/di && wire

order.test.mock:
	mockgen -destination=app/order_api/domain/port/mock/domain_event.go --package=mock github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event EventHandler,Publisher
	mockgen -destination=app/order_api/domain/port/mock/external_api.go -package=mock github.com/tkame123/ddd-sample/app/order_api/domain/port/external_service BillingAPI,KitchenAPI
	mockgen -destination=app/order_api/domain/port/mock/repository.go -package=mock github.com/tkame123/ddd-sample/app/order_api/domain/port/repository Repository,CreateOrderSagaState,Order
	mockgen -destination=app/order_api/domain/port/mock/service.go -package=mock github.com/tkame123/ddd-sample/app/order_api/domain/port/service CreateOrder

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

kitchen.wire:
	cd app/kitchen_api/di && wire

billing.wire:
	cd app/billilng_api/di && wire
