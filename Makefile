dev:
	POSTGRES_DSN="postgresql://postgres:123123@localhost:5434/snowboardsdb?sslmode=disable" go run ./cmd/snowboardsdb

gen:
	go generate ./...