package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/samandar2605/practice_api/api/v1"
	"github.com/samandar2605/practice_api/config"
	"github.com/samandar2605/practice_api/storage"

	_ "github.com/samandar2605/practice_api/api/docs" // for swagger

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type RouterOptions struct {
	Cfg     *config.Config
	Storage storage.StorageI
}

// @title           Swagger for blog api
// @version         1.0
// @description     This is a blog service api.
// @host      		localhost:8000
// @BasePath  		/v1
func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Cfg:     opt.Cfg,
		Storage: opt.Storage,
	})

	apiV1 := router.Group("/v1")

	// Students
	apiV1.GET("/students", handlerV1.GetAllStudent)
	apiV1.POST("/students", handlerV1.CreateStudent)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
