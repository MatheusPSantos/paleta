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
	/** Customer **/
	r.POST("/customers", controllers.CreateCustomer)
	r.GET("/customers", controllers.GetCustomers)
	r.GET("/customers/:id", controllers.GetCustomerById)
	r.PUT("/customers/:id", controllers.UpdateCustomer)
	r.DELETE("/customers/:id", controllers.DeleteCustomer)
	/** Address **/
	r.POST("/address", controllers.CreateAddress)
	r.GET("/users/:id/address", controllers.GetUserAddress)
	r.PUT("/users/:id/address/:addressId", controllers.UpdateUserAddress)
	r.DELETE("/users/:id/address/:addressId", controllers.DeleteUserAddress)
	/** Loyalty **/
	r.POST("/loyalty/campaing", controllers.CreateLoyaltyCampaing)
	r.GET("/loyalty/seller/:id/campaings", controllers.ListSellerLoyaltyCampaings)
	// r.GET("/loyalty/seller/:id/campaings/:campaingId", controllers.GetSellerLoyaltyCampaingById)
	// r.GET("/loyalty/customer/:id/campaings", controllers.ListCustomerLoyaltyCampaings)
	// r.GET("/loyalty/customer/:id/campaings/:campaingId", controllers.GetCustomerLoyaltyCampaingById)
	return r
}
