.PHONY: *
export
COMPOSE_FILE = docker-compose.yml

build:
	docker compose -f ${COMPOSE_FILE} build $(s) --no-cache
up:
	docker compose -f ${COMPOSE_FILE} up $(s) -d
logs:
	docker compose -f ${COMPOSE_FILE} logs $(s) -f
down:
	docker compose -f ${COMPOSE_FILE} down

