build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server/build/server -tags=jsoniter server/main.go
	docker build --rm -f server/Dockerfile -t herd/backend:latest server

	npm run build
	docker build --rm -f client/Dockerfile -t herd/client:latest client
run-dev:
	docker run --rm -it -p 8000:8080 herd/backend:latest
migrate:
	go run migrations/migrations.go
#^^^ add target db variable
