package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rathod-pritesh/portfolio/db"
	"github.com/rathod-pritesh/portfolio/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func certCol() *mongo.Collection {
	return db.GetClient().Database("portfolio").Collection("certifications")
}

// ── GET /api/admin/certifications
func AdminGetCertifications(c *gin.Context) {
	opts := options.Find().SetSort(bson.D{{Key: "order", Value: 1}})
	cursor, err := certCol().Find(context.TODO(), bson.M{}, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var certs []models.Certification
	if err := cursor.All(context.TODO(), &certs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, certs)
}

// ── POST /api/admin/certifications
func AdminAddCertification(c *gin.Context) {
	var input models.Certification
	if err := c.ShouldBindJSON(&input); err != nil || input.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	// auto-set order if not provided
	if input.Order == 0 {
		count, _ := certCol().CountDocuments(context.TODO(), bson.M{})
		input.Order = int(count) + 1
	}

	result, err := certCol().InsertOne(context.TODO(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	input.ID = result.InsertedID.(primitive.ObjectID).Hex()
	c.JSON(http.StatusOK, gin.H{"message": "Certification added", "data": input})
}

// ── PUT /api/admin/certifications/:id
func AdminUpdateCertification(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var input models.Certification
	if err := c.ShouldBindJSON(&input); err != nil || input.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	_, err = certCol().UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": bson.M{
			"name":      input.Name,
			"company":   input.Company,
			"link":      input.Link,
			"issueDate": input.IssueDate,
			"order":     input.Order,
		}},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Certification updated"})
}

// ── DELETE /api/admin/certifications/:id
func AdminDeleteCertification(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	_, err = certCol().DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Certification deleted"})
}
