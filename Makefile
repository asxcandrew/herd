deploy:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server/build/server -tags=jsoniter server/main.go

	export $(cat .env | xargs) && yarn build --prefix ./client
	eval $(docker-machine env herd-production)
	### This is temporary solution
	rsync -rvz --rsh='docker-machine ssh herd-production' --progress ./config/nginx.prod.conf.template :/home/nginx.prod.conf.template
	rsync -rvz --rsh='docker-machine ssh herd-production' --progress ./client/dist :/home/
	###
	docker-compose up --no-deps --build -d
run-dev:
	eval $(docker-machine env -u)
	docker-compose -f docker-compose.dev.yml -p herd-dev up
run-api-tests:

setup-dev:
	docker-compose -f docker-compose.dev.yml -p herd-dev build
