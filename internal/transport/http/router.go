package http

import (
	_ "awesomeProject/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) UserRoutes(incomingRoutes *gin.Engine) {
	//transactions
	t := incomingRoutes.Group("/trans")
	t.POST("/create", s.handler.CreateTrans())
	t.GET("/list", s.handler.ListTrans())
	t.GET("/get_by_id/:id", s.handler.GetByID())
	t.GET("/delete/:id", s.handler.DeleteTrans())

	//swagger
	incomingRoutes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
