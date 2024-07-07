package routes

import (
	"log"
	"net/http"
	"paleta-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	log.Print("Setting up routes and handlers")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })

	/** Seller */
	r.GET("/sellers", controllers.GetSellers)
	r.GET("/sellers/:id", controllers.GetSeller)
	r.POST("/sellers", controllers.CreateSeller)
	r.PUT("/sellers/:id", controllers.UpdateSeller)
	r.DELETE("/sellers/:id", controllers.DeleteSeller)

	r.GET("/sellers/:id/address", controllers.GetSellerAddress)

	/** Address **/
	r.POST("/address", controllers.CreateAddress)

	return r
}
