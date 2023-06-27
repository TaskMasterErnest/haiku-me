run:
	go run ./cmd/web
unit-tests:
	go test -v ./cmd/web
clean-cache:
	go clean -testcache
build-db:
	docker build -f Dockerfile.db -t ernestklu/haibox:db-v1 --build-arg MYSQL_USER=ernest --build-arg MYSQL_PASSWORD=connect@db1 --build-arg MYSQL_ROOT_PASSWORD=p@ssw0r/d .
build-app:
	docker build -f Dockerfile.app -t ernestklu/haibox:app-v1 .
compose:
	docker-compose -f docker-compose.yml up --detach