mysql:
	docker run --name sosdbcontainer -e MYSQL_ROOT_PASSWORD=mypassword -p 3308:3306 -d mysql:latest
createdb:
	docker exec -it sosdbcontainer mysql -u root -pmypassword -e "CREATE DATABASE sosprojectdb;"

dropdb:
	docker exec -it sosdbcontainer mysql -u root -pmypassword -e "DROP DATABASE sosprojectdb;"

migrateup:
	migrate --path db/migrations --database "mysql://root:mypassword@tcp(127.0.0.1:3308)/sosprojectdb" --verbose up

migratedown:
	migrate --path db/migrations --database "mysql://root:mypassword@tcp(127.0.0.1:3308)/sosprojectdb" --verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

build:
	go build -o sosproject github.com/avemoi/sosproject/cmd/api

prod:
	CGO_ENABLED=0 go build  github.com/avemoi/sosproject/cmd/api


.PHONY: postgres createdb dropdb migrateup migratedown