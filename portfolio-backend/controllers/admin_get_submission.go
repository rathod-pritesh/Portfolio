package controllers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/rathod-pritesh/portfolio/db"
	"github.com/rathod-pritesh/portfolio/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AdminGetSubmissions(c *gin.Context) {
	client := db.GetClient()
	collection := client.Database("portfolio").Collection("contact_submission")

	opts := options.Find().SetSort(bson.D{{Key: "order", Value: 1}})
	cursor, err := collection.Find(context.TODO(), bson.M{}, opts)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var submissions []models.ContactSubmission
	cursor.All(context.TODO(), &submissions)

	c.JSON(200, submissions)
}
