package main

import (
	framework "go_server/framework"
	"log"
)

func StartServer() *WE.Engine {
	e := framework.New()
	c := Container{"ssss"}
	e.GET("/login", c.Login)
	e.Static("/", "dist")
	e.Run(":2325")
	return e
}

type Container struct {
	DB string
}

func (sc *Container) Login(c *WE.Context) {
	log.Print("Login", sc.DB)
	c.Success(nil)
}

func TestMiddleware(c *WE.Context) {
	log.Print("TestMiddleware")
	c.Next()
}
