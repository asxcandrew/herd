deploy:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server/build/server -tags=jsoniter server/main.go
	# docker build --rm -f server/Dockerfile -t herd/backend:latest server
	npm run build --prefix ./client
	eval $(docker-machine env production)

	docker-compose up --no-deps --build -d
run-dev:
	docker-compose -f docker-compose.dev.yml -p herd-dev up
setup-dev:
	docker-compose -f docker-compose.dev.yml -p herd-dev up --build
