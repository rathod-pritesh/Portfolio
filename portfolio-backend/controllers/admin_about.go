package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rathod-pritesh/portfolio/db"
	"github.com/rathod-pritesh/portfolio/models"
	"go.mongodb.org/mongo-driver/bson"
)

// GET /api/admin/about  — fetch current about data for the admin form
func AdminGetAbout(c *gin.Context) {
	collection := db.GetClient().Database("portfolio").Collection("about")

	var about models.About
	err := collection.FindOne(context.TODO(), bson.M{"_id": "about"}).Decode(&about)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, about)
}

// PUT /api/admin/about  — update about data
func AdminUpdateAbout(c *gin.Context) {
	var input models.About

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	collection := db.GetClient().Database("portfolio").Collection("about")

	update := bson.M{
		"$set": bson.M{
			"title":         input.Title,
			"description":   input.Description,
			"highlights":    input.Highlights,
		},
	}

	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": "about"}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "About updated successfully"})
}