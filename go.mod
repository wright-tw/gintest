module gintest

// go 版本
go 1.14

// gin 框架
require github.com/gin-gonic/gin v1.6.3

// log 套件
require github.com/sirupsen/logrus v1.6.0

// gorm
require github.com/jinzhu/gorm v1.9.14

// 環境套件 .env
require github.com/joho/godotenv v1.3.0

require (
	github.com/go-redis/redis/v8 v8.2.3
	// DI 套件
	go.uber.org/dig v1.10.0
)
