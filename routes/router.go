package routes

import (
	"jwt-project/controller"
	"jwt-project/middleware"

	"github.com/gin-gonic/gin"
)

type Router interface {
	Setup()
	Run(string)
}

type router struct {
	engine     *gin.Engine
	controller controller.Controller
}

func NewRouter(
	controller controller.Controller,
) Router {
	return &router{
		controller: controller,
	}
}

func (r *router) Run(serverHost string) {
	r.engine.Run(serverHost)
}

func (r *router) Setup() {
	r.engine = gin.New()
	r.engine.Use(gin.Logger())
	r.engine.Use(middleware.Autheticate())

	r.authenticationRoutes()
	r.personRoutes()
}

func (r *router) authenticationRoutes() {
	r.engine.POST("/person/signup", r.controller.SignUp)
	r.engine.POST("/person/login", r.controller.Login)
}

func (r *router) personRoutes() {
	r.engine.GET("/person/:userid", r.controller.GetUser)
	r.engine.GET("/personall", r.controller.GetUsers)
}
