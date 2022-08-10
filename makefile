build_dev:
	docker-compose -f docker-compose.dev.yaml build

start_dev:
	docker-compose -f docker-compose.dev.yaml up -d

start_prod:
	docker-compose up

build_prod:
	docker-compose build