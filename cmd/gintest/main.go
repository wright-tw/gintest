package main

import (
    "gintest/configs"
    "gintest/internal/app/gintest/routes"
    "gintest/pkg/logger"
    "github.com/gin-gonic/gin"
    "os"
)

func main() {

    configs.Setting()
    logger.Info("設定環境變數完成", logger.SERVER)

    server := gin.Default()
    logger.Info("gin 框架載入完成", logger.SERVER)

    router := routes.Router{}
    router.Setting(server)
    logger.Info("路由設定完成", logger.SERVER)

    logger.Info("啟動 http server", logger.SERVER)
    server.Run(":" + os.Getenv("PORT"))
}