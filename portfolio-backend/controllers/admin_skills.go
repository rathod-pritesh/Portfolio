package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rathod-pritesh/portfolio/db"
	"github.com/rathod-pritesh/portfolio/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func skillsCol() *mongo.Collection {
	return db.GetClient().Database("portfolio").Collection("skills")
}

// fetch the full skills doc
func getSkillsDoc() (models.Skills, error) {
	var skills models.Skills
	err := skillsCol().FindOne(context.TODO(), bson.M{"_id": "skills"}).Decode(&skills)
	return skills, err
}

// save
func saveSkillsDoc(skills models.Skills) error {
	_, err := skillsCol().UpdateOne(
		context.TODO(),
		bson.M{"_id": "skills"},
		bson.M{"$set": bson.M{
			"title":      skills.Title,
			"categories": skills.Categories,
		}},
	)
	return err
}

// GET /api/admin/skills  — fetch current skills data for the admin form
func AdminGetSkills(c *gin.Context) {
	skills, err := getSkillsDoc()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, skills)
}

// PUT /api/admin/skills/title  — update skills data
func AdminUpdateSkillsTitle(c *gin.Context) {
	var body struct {
		Title string `json:"title"`
	}
	if err := c.ShouldBindJSON(&body); err != nil || body.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}

	_, err := skillsCol().UpdateOne(
		context.TODO(),
		bson.M{"_id": "skills"},
		bson.M{"$set": bson.M{"title": body.Title}},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Title updated"})
}

// POST /api/admin/skills/category
func AdminAddCategory(c *gin.Context) {
	var newCat models.Category
	if err := c.ShouldBindJSON(&newCat); err != nil || newCat.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category name is required"})
		return
	}
	newCat.Skills = []models.Skill{}

	skills, err := getSkillsDoc()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newCat.Order = len(skills.Categories) + 1
	skills.Categories = append(skills.Categories, newCat)

	if err := saveSkillsDoc(skills); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category added", "data": skills})
}

// ── PUT /api/admin/skills/category/:ci
func AdminUpdateCategory(c *gin.Context) {
	ci, err := strconv.Atoi(c.Param("ci"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid index"})
		return
	}

	var body struct {
		Name string `json:"name"`
		Icon string `json:"icon"`
	}
	if err := c.ShouldBindJSON(&body); err != nil || body.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	skills, err := getSkillsDoc()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if ci < 0 || ci >= len(skills.Categories) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category index out of range"})
		return
	}

	skills.Categories[ci].Name = body.Name
	skills.Categories[ci].Icon = body.Icon

	if err := saveSkillsDoc(skills); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category updated", "data": skills})
}

// ── DELETE /api/admin/skills/category/:ci
func AdminDeleteCategory(c *gin.Context) {
	ci, err := strconv.Atoi(c.Param("ci"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid index"})
		return
	}

	skills, err := getSkillsDoc()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if ci < 0 || ci >= len(skills.Categories) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category index out of range"})
		return
	}

	skills.Categories = append(skills.Categories[:ci], skills.Categories[ci+1:]...)

	if err := saveSkillsDoc(skills); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted", "data": skills})
}

// ── POST /api/admin/skills/category/:ci/skill
func AdminAddSkills(c *gin.Context) {
	ci, err := strconv.Atoi(c.Param("ci"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid index"})
		return
	}

	var newSkills models.Skill
	if err := c.ShouldBindJSON(&newSkills); err != nil || newSkills.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "skill name is required"})
		return
	}

	skills, err := getSkillsDoc()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if ci < 0 || ci >= len(skills.Categories) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category index out of range"})
		return
	}

	skills.Categories[ci].Skills = append(skills.Categories[ci].Skills, newSkills)

	if err := saveSkillsDoc(skills); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Skill added", "data": skills})

}

func AdminUpdateSkill(c *gin.Context) {
	ci, err1 := strconv.Atoi(c.Param("ci"))
	si, err2 := strconv.Atoi(c.Param("si"))
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid index"})
		return
	}

	var body models.Skill
	if err := c.ShouldBindJSON(&body); err != nil || body.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "skill name is required"})
		return
	}

	skills, err := getSkillsDoc()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if ci < 0 || ci >= len(skills.Categories) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category index out of range"})
		return
	}
	if si < 0 || si >= len(skills.Categories[ci].Skills) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "skill index out of range"})
		return
	}

	skills.Categories[ci].Skills[si] = body

	if err := saveSkillsDoc(skills); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Skill updated", "data": skills})
}

// ── DELETE /api/admin/skills/category/:ci/skill/:si
func AdminDeleteSkill(c *gin.Context) {
	ci, err1 := strconv.Atoi(c.Param("ci"))
	si, err2 := strconv.Atoi(c.Param("si"))
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid index"})
		return
	}

	skills, err := getSkillsDoc()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if ci < 0 || ci >= len(skills.Categories) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category index out of range"})
		return
	}
	if si < 0 || si >= len(skills.Categories[ci].Skills) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "skill index out of range"})
		return
	}

	cat := &skills.Categories[ci]
	cat.Skills = append(cat.Skills[:si], cat.Skills[si+1:]...)

	if err := saveSkillsDoc(skills); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Skill deleted", "data": skills})
}
