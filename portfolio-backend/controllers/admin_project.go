package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rathod-pritesh/portfolio/db"
	"github.com/rathod-pritesh/portfolio/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Order Asc get all projects
func AdminGetProjects(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	col := db.GetClient().Database("portfolio").Collection("projects")
	opts := options.Find().SetSort(bson.D{{Key: "order", Value: 1}})

	cursor, err := col.Find(ctx, bson.M{}, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch projects"})
		return
	}
	defer cursor.Close(ctx)

	var projects []models.Project
	if err := cursor.All(ctx, &projects); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode projects" + err.Error()})
		return
	}

	if projects == nil {
		projects = []models.Project{}
	}

	c.JSON(http.StatusOK, projects)
}

// POST
func AdminAddProject(c *gin.Context) {
	var input struct {
		Title string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
		Image string `json:"image"`
		GitHub string `json:"github"`
		Technologies []string `json:"technologies"`
		Gradient string `json:"gradient"`
		Order int `json:"order"`
		IsFeatured bool `json:"isFeatured"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	col := db.GetClient().Database("portfolio").Collection("projects")

	newID := primitive.NewObjectID().Hex()

	projects := bson.M{
		"_id": newID,
		"title": input.Title,
		"description": input.Description,
		"image": input.Image,
		"github": input.GitHub,
		"technologies": input.Technologies,
		"gradient": input.Gradient,
		"order": input.Order,
		"isFeatured": input.IsFeatured,
	}

	_, err := col.InsertOne(ctx, projects)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add project"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Project added successfully",
		"insertedId": newID,
	})
}

// PUT
func AdminUpdateProject(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Title *string `json:"title"`
		Description *string `json:"description"`
		Image *string `json:"image"`
		GitHub *string `json:"github"`
		Technologies []string `json:"technologies"`
		Gradient *string `json:"gradient"`
		Order *int `json:"order"`
		IsFeatured *bool `json:"isFeatured"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateFields := bson.M{}
	if input.Title != nil { updateFields["title"] = *input.Title }
	if input.Description != nil { updateFields["description"] = *input.Description }
	if input.Image != nil { updateFields["image"] = *input.Image }
	if input.GitHub != nil { updateFields["github"] = *input.GitHub }
	if input.Technologies != nil { updateFields["technologies"] = input.Technologies }
	if input.Gradient != nil { updateFields["gradient"] = *input.Gradient }
	if input.Order != nil { updateFields["order"] = *input.Order }
	if input.IsFeatured != nil { updateFields["isFeatured"] = *input.IsFeatured }

	if len(updateFields) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No fields to update"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	col := db.GetClient().Database("portfolio").Collection("projects")
	result, err := col.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": updateFields},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project updated successfully"})
}

// DELETE
func AdminDeleteProject(c *gin.Context) {
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	col := db.GetClient().Database("portfolio").Collection("projects")
	result, err := col.DeleteOne(
		ctx,
		bson.M{"_id": id},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project"})
		return
	}
	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}

// PUT - reorder projects
func AdminReorderProject(c *gin.Context) {
	var items []struct {
		ID string `json:"id" binding:"required"`
		Order int `json:"order"`
	}

	if err := c.ShouldBindJSON(&items); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	col := db.GetClient().Database("portfolio").Collection("projects")
	for _, item := range items {
		col.UpdateOne(ctx, bson.M{"_id": item.ID}, bson.M{"$set": bson.M{"order": item.Order}})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project reordered"})
}


