package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rathod-pritesh/portfolio/db"
	"github.com/rathod-pritesh/portfolio/models"
	"go.mongodb.org/mongo-driver/bson"
)

// GET /api/admin/home  — fetch current home data for the admin form
func AdminGetHome(c *gin.Context) {
	collection := db.GetClient().Database("portfolio").Collection("home")

	var home models.Home
	err := collection.FindOne(context.TODO(), bson.M{"_id": "home"}).Decode(&home)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, home)
}

// PUT /api/admin/home  — update home data
func AdminUpdateHome(c *gin.Context) {
	var input models.Home

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	collection := db.GetClient().Database("portfolio").Collection("home")

	update := bson.M{
		"$set": bson.M{
			"name":         input.Name,
			"role":         input.Role,
			"tagline":      input.Tagline,
			"resumeUrl":    input.ResumeUrl,
			"profileImage": input.ProfileImage,
		},
	}

	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": "home"}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Home updated successfully"})
}