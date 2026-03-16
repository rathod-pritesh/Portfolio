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

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Portfolio
	r.GET("/api/home", controllers.GetHome)
	r.GET("/api/about", controllers.GetAbout)
	r.GET("/api/skills", controllers.GetSkills)
	r.GET("/api/automations", controllers.GetAutomations)
	r.GET("/api/projects", controllers.GetProjects)
	r.GET("/api/certifications", controllers.GetCertifications)

	// Admin Auth
	r.POST("/api/admin/login", controllers.AdminLoginHandler)
	r.GET("/api/admin/verify", controllers.AdminVerifyHandler)
	r.POST("/api/admin/logout", controllers.AdminLogoutHandler)

	// Admin CRUD
	admin := r.Group("/api/admin", controllers.AdminAuthMiddleware)
	{
		// Home section
		admin.GET("/home", controllers.AdminGetHome)
		admin.PUT("/home", controllers.AdminUpdateHome)
	}

	for _, route := range r.Routes() {
		println(route.Method, route.Path)
	}

	r.Run(":8080")
}
