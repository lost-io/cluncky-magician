
runserver:
	go run ./app/main.go

runtests:
	cd app/
	go test ./tests

runtestsv:
	go test ./app/tests -v

build:
	docker build -t cardapi .

compose-build:
	docker-compose build