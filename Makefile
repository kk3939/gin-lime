.PHONY: up
up:
	docker-compose --env-file ./.env.local up --build

.PHONY: down
down:
	docker-compose --env-file ./.env.local down

.PHONY: test
test:
	go test -v -cover ./.../
