Directory Structure

gqlgengin
└─internal
    └─pkg
        └─db
            └─migrations
                └─mysql


Initialization

go get -u github.com/go-sql-driver/mysql
go build -tags 'mysql' -ldflags="-X main.Version=1.0.0" -o $GOPATH/bin/migrate github.com/golang-migrate/migrate/v4/cmd/migrate/
cd internal/pkg/db/migrations/
migrate create -ext sql -dir mysql -seq create_users_table
migrate create -ext sql -dir mysql -seq create_links_table

Run this command in your project root directory:
migrate -database "mysql://root:root@tcp(localhost:3306)/gqlgengin" -path internal/pkg/db/migrations/mysql -verbose up