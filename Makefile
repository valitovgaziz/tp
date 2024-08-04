build:
	@docker compose build

run: build
	@docker compose up

clean:
	@docker builder prune

.DEFAULT_GOAL=run