package app

import (
	"0byte/app/middleware"
	"0byte/controllers/proof"
	"0byte/services/proofsvc"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MapUrl() {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	// proof handler
	proofSvc := proofsvc.Handler()
	proofHanlder := proof.Handler(proofSvc)

	// health check route
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	router.POST("/proof", proofHanlder.GenerateProof)

	err := router.Run()
	if err != nil {
		panic(err.Error() + "router not able to run")
	}
}
