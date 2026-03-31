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
	r.GET("/api/education", controllers.GetEducation)

	// Admin Auth
	r.POST("/api/admin/login", controllers.AdminLoginHandler)
	r.GET("/api/admin/verify", controllers.AdminVerifyHandler)
	r.POST("/api/admin/logout", controllers.AdminLogoutHandler)

	// Admin CRUD
	admin := r.Group("/api/admin", controllers.AdminAuthMiddleware)
	{
		// Home 
		admin.GET("/home", controllers.AdminGetHome)
		admin.PUT("/home", controllers.AdminUpdateHome)

		// About
		admin.GET("/about",controllers.AdminGetAbout)
		admin.PUT("/about", controllers.AdminUpdateAbout)

		// Skills
		admin.GET("/skills", controllers.AdminGetSkills)
		admin.PUT("/skills/title", controllers.AdminUpdateSkillsTitle)
		admin.POST("/skills/category", controllers.AdminAddCategory)
		admin.PUT("/skills/category/:ci", controllers.AdminUpdateCategory)
		admin.DELETE("/skills/category/:ci", controllers.AdminDeleteCategory)
		admin.POST("/skills/category/:ci/skill", controllers.AdminAddSkills)
		admin.PUT("/skills/category/:ci/skill/:si", controllers.AdminUpdateSkill)
		admin.DELETE("/skills/category/:ci/skill/:si", controllers.AdminDeleteSkill)

		// Certification
		admin.GET("/certifications", controllers.AdminGetCertifications)
		admin.POST("/certifications", controllers.AdminAddCertification)
		admin.PUT("/certifications/:id", controllers.AdminUpdateCertification)
		admin.DELETE("/certifications/:id", controllers.AdminDeleteCertification)

		// Project
		admin.GET("/projects", controllers.AdminGetProjects)
		admin.POST("/projects", controllers.AdminAddProject)
		admin.PUT("/projects/reorder", controllers.AdminReorderProject)
		admin.PUT("/projects/:id", controllers.AdminUpdateProject)
		admin.DELETE("/projects/:id", controllers.AdminDeleteProject)
	}

	for _, route := range r.Routes() {
		println(route.Method, route.Path)
	}

	r.Run(":8080")
}
