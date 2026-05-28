package controllers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rathod-pritesh/portfolio/db"
	"github.com/rathod-pritesh/portfolio/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// Education API
func GetEducation(c *gin.Context) {
	client := db.GetClient()
	collection := client.Database("portfolio").Collection("education")

	opts := options.Find().SetSort(bson.D{{Key: "order", Value: -1}}) //descending

	cursor, err := collection.Find(context.TODO(), bson.M{}, opts)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var education []models.Education
	if err := cursor.All(context.TODO(), &education); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, education)
}

// contact form submit
func SubmitContact(c *gin.Context) {
	var body struct {
		Name string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required"`
		Message string `json:"message" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "All fields are required"})
		return
	}

	client := db.GetClient()
	collection := client.Database("portfolio").Collection("contact_submission")

	submission := models.ContactSubmission{
		ID: primitive.NewObjectID(),
		Name: body.Name,
		Email: body.Email,
		Message: body.Message,
		CreatedAt: time.Now(),
	}

	_, err := collection.InsertOne(context.TODO(), submission)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to save message"})
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}