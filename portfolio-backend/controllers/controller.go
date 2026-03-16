package controllers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/rathod-pritesh/portfolio/db"
	"github.com/rathod-pritesh/portfolio/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Home API
func GetHome(c *gin.Context) {

	client := db.GetClient()
	collection := client.Database("portfolio").Collection("home")

	var home models.Home

	err := collection.FindOne(context.TODO(), bson.M{"_id": "home"}).Decode(&home)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, home)
}

// About API
func GetAbout(c *gin.Context) {

	client := db.GetClient()
	collection := client.Database("portfolio").Collection("about")

	var about models.About

	collection.FindOne(context.TODO(), bson.M{"_id": "about"}).Decode(&about)

	c.JSON(200, about)
}

// Skills API
func GetSkills(c *gin.Context) {

	client := db.GetClient()
	collection := client.Database("portfolio").Collection("skills")

	var skills models.Skills

	collection.FindOne(context.TODO(), bson.M{"_id": "skills"}).Decode(&skills)

	c.JSON(200, skills)
}

// Automation API
func GetAutomations(c *gin.Context) {
	client := db.GetClient()
	collection := client.Database("portfolio").Collection("automations")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var automations []models.Automation
	cursor.All(context.TODO(), &automations)

	c.JSON(200, automations)
}

// Certification API
func GetCertifications(c *gin.Context) {
	client := db.GetClient()
	collection := client.Database("portfolio").Collection("certifications")

	opts := options.Find().SetSort(bson.D{{Key: "order", Value: 1}})

	cursor, err := collection.Find(context.TODO(), bson.M{}, opts)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var certifications []models.Certification
	if err := cursor.All(context.TODO(), &certifications); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, certifications)
}

// Projects API
func GetProjects(c *gin.Context) {

	client := db.GetClient()
	collection := client.Database("portfolio").Collection("projects")

	cursor, _ := collection.Find(context.TODO(), bson.M{})

	var projects []models.Project

	cursor.All(context.TODO(), &projects)

	c.JSON(200, projects)
}
