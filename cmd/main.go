package main

import (
	"log"

	"github.com/AnanyaDevGo/pkg/auth"
	"github.com/AnanyaDevGo/pkg/config"
	"github.com/AnanyaDevGo/pkg/order"
	"github.com/AnanyaDevGo/pkg/product"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)

}
