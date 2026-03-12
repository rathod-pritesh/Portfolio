package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rathod-pritesh/portfolio/controllers"
	"github.com/rathod-pritesh/portfolio/db"
)

func main() {
	db.ConnectDB()

	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/api/home", controllers.GetHome)
	r.GET("/api/about", controllers.GetAbout)
	r.GET("/api/skills", controllers.GetSkills)
	r.GET("/api/automations", controllers.GetAutomations)
	r.GET("/api/projects", controllers.GetProjects)
	r.GET("/api/certifications", controllers.GetCertifications)

	for _, route := range r.Routes() {
		println(route.Method, route.Path)
	}

	r.Run(":8080")
}
