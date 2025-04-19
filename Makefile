.PHONY: build dev
dev:
	@docker compose up
build:
	@docker compose up --build