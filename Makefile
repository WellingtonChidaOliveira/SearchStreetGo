build:
	go build -o ./bin/server
run: build
	 @PORT=:8080 \
    APP_ENV=local \
    DB_HOST=localhost \
    MANAGER_PORT=8081 \
    DB_PORT=27017 \
    DB_USERNAME=root \
    DB_PASS=example \
    DB_COLLECTION=streets \
    DB_DATABASENAME=searchstreet \
    ./bin/server