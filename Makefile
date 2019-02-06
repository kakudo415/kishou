main: package
	go build -o bin/kishow src/main.go
package:
	go get -u github.com/labstack/echo
	go get -u github.com/go-redis/redis
	go get -u github.com/jinzhu/gorm
	go get -u github.com/go-sql-driver/mysql
	go get -u github.com/mmcdole/gofeed
	go get -u github.com/google/uuid