### SETUP
##### install
    go get github.com/mattes/migrate

##### db migration
    if psql "host=host user=username password=password port=port sslmode=disable" -l | cut -d \| -f 1 | grep -qw $DB_NAME; then
      echo Database exists. Do nothing
    else
      echo Database is not exist
      psql "host=host user=username password=password port=port sslmode=disable" -c "CREATE DATABASE database_name ENCODING 'UTF8' TEMPLATE template0"
    fi

    migrate -url postgres://username:password@host:port/database_name?sslmode=disable -path ./migration_path up
    migrate -url postgres://username:password@host:port/database_name?sslmode=disable -path ./migration_path down