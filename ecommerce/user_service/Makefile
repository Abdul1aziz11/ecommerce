DB_URL=postgres://postgres:1@localhost:5432/imtihon3_users?sslmode=disable

gen-proto:
	@protoc \
	--go_out=. \
	--go-grpc_out=. \
	protos/user.proto

migrate_up:
	migrate -path ./db/migrations -database $(DB_URL) up

migrate_down:
	migrate -path ./db/migrations -database $(DB_URL) down

migrate_force:
	migrate -path=db/migrations -database ${DB_URL} -verbose force 1
