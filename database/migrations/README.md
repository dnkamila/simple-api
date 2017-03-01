### SETUP
###### install
    go get github.com/mattes/migrate

###### db migration
    migrate -url driver://url -path ./migrations up
    migrate -url postgres://username:password@host:port/database_name?sslmode=disable