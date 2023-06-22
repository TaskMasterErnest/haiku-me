run:
	go run ./cmd/web
unit-tests:
	go test -v ./cmd/web
clean-cache:
	go clean -testcache
build-db:
	docker build -f Dockerfile.db -t maria-db:""$GIT_COMMIT"" --build-arg MYSQL_USER=ernest --build-arg MYSQL_PASSWORD=connect@db1 --build-arg MYSQL_ROOT_PASSWORD=p@ssw0r/d .
build-app:
	docker build -f Dockerfile.app -t haiku-app:""$GIT_COMMIT"" .
compose:
	docker-compose -f docker-compose.yml up --detach