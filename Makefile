install:
	go install github.com/golang/mock/mockgen

test.mock:
	mockgen -source=app/order_api/domain/port/domain_event/publisher.go -destination=app/order_api/domain/port/mock/domain_event/publisher.go -package=mock
	mockgen -source=app/order_api/domain/port/external_service/billing_api.go -destination=app/order_api/domain/port/mock/external_service/billing_api.go -package=mock
	mockgen -source=app/order_api/domain/port/external_service/ticket_api.go -destination=app/order_api/domain/port/mock/external_service/ticket_api.go -package=mock
	mockgen -source=app/order_api/domain/port/repository/repository.go -destination=app/order_api/domain/port/mock/repository/repository.go -package=mock
	mockgen -source=app/order_api/domain/port/repository/create_order_saga_state.go -destination=app/order_api/domain/port/mock/repository/create_order_saga_state.go -package=mock
	mockgen -source=app/order_api/domain/port/repository/order.go -destination=app/order_api/domain/port/mock/repository/order.go -package=mock
	mockgen -source=app/order_api/domain/port/service/create_order.go -destination=app/order_api/domain/port/mock/service/create_order.go -package=mock

