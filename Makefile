build:
	@docker compose build

run: build
	@docker compose up

clean:
	@docker builder prune

test:
	@go test ./api/src/auth/... -v

tc:
	@go test -cover

.DEFAULT_GOAL=run