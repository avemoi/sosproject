mysql:
	docker run --name lacubamydbcontainer -e MYSQL_ROOT_PASSWORD=mypassword -p 3307:3306 -d mysql:latest
createdb:
	docker exec -it lacubamydbcontainer mysql -u root -pmypassword -e "CREATE DATABASE sosprojectdb;"

dropdb:
	docker exec -it lacubacontainer dropdb --username=postgres  sosprojectdb

migrateup:
	migrate --path db/migrations --database "mysql://root:mypassword@tcp(127.0.0.1:3307)/sosprojectdb" --verbose up

migratedown:
	migrate --path db/migrations --database "mysql://root:mypassword@tcp(127.0.0.1:3307)/sosprojectdb" --verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

build:
	go build -ldflags="-I cmd/api -I db/sqlc"


.PHONY: postgres createdb dropdb migrateup migratedown