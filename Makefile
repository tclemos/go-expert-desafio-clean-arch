.PHONY: proto-gen
proto-gen:
	protoc --go_out=. --go-grpc_out=. internal/infra/grpc/proto/order.proto

.PHONY: run
run:
	docker compose up -d --build

.PHONY: stop
stop:
	docker compose down --remove-orphans

.PHONY: restart
restart: stop run

