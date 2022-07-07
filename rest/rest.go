package rest

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	h, err := NewHandler()
	if err != nil {
		return err
	}
	return RunAPIWithHandler(address, h)
}

func RunAPIWithHandler(address string, h HandlerInterface) error {
	//Get gin's default engine
	r := gin.Default()
	r.Use(MyCustomLogger())

	//get products
	r.GET("/products", h.GetProducts)
	/*
		//post user sign in
		r.POST("/user/signin", h.SignIn)
		//post user sign out
		r.POST("/user/:id/signout", h.SignOut)
		//get user orders
		r.GET("/user/:id/orders", h.GetOrders)
		//post purchase charge
		r.POST("/user/charge", h.Charge)
	*/

	return r.Run(address)
	//return r.RunTLS(address, "cert.pem", "key.pem")
}

func MyCustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("************************************")
		c.Next()
		fmt.Println("************************************")
	}
}
