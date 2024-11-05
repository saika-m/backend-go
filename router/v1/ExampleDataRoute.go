package route

import (
	"github.com/Takina-Space/backend-go/app/controller"
	"github.com/Takina-Space/backend-go/app/repository"
	"github.com/Takina-Space/backend-go/app/service"
	"github.com/Takina-Space/backend-go/config/database"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ExampleRoute(router *gin.RouterGroup) {
	validate := validator.New()
	PostgreDB := database.PostgreDB

	exampleRepository := repository.NewExampleRepository(PostgreDB)
	exampleService := service.NewExampleService(exampleRepository)
	authService := service.NewAuthService()
	exampleController := controller.NewExampleController(exampleService, authService, validate)

	//ExampleData
	router.POST("/example-data", exampleController.CreateExampleData)
	router.GET("/example-data", exampleController.GetExampleDatas)
	router.GET("/example-data/:example_data_id", exampleController.GetExampleDataById)
	router.PUT("/example-data/:example_data_id", exampleController.EditExampleData)
	router.DELETE("/example-data/:example_data_id", exampleController.DeleteExampleData)
}
