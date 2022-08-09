run: migrate server

migrate: ## dbmate migrate
	docker-compose up -d mysql
	docker-compose run --rm dockerize -wait tcp://mysql:3306 -timeout 20s
	docker-compose run --rm migrate

server: ## go run server
	docker-compose up -d mysql
	docker-compose run --rm dockerize -wait tcp://mysql:3306 -timeout 20s
	docker-compose up api
