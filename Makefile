build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server/build/server -tags=jsoniter server/main.go
	docker build --rm -f server/Dockerfile -t herd/backend:latest server

	npm run build
	docker build --rm -f client/Dockerfile -t herd/client:latest client
run-dev:
	docker-compose -f docker-compose.dev.yml -p herd-dev up
setup-dev:
	docker-compose -f docker-compose.dev.yml -p herd-dev up --build
migrate:
	go run migrations/migrations.go
#^^^ add target db variable
