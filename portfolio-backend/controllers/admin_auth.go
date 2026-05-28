package controllers

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/rathod-pritesh/portfolio/db"
	"github.com/rathod-pritesh/portfolio/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func jwtSecret() []byte {
	s := os.Getenv("JWT_SECRET")
	if s == "" {
		s = "avcrojdlaofndnsdowpawfa"
	}

	return  []byte(s)
}

func makeToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret())
}

func parseToken(tokeStr string) (string, bool) {
	token, err := jwt.Parse(tokeStr, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret(), nil
	})

	if err != nil || !token.Valid{
		return "", false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", false
	}

	email, _ := claims["email"].(string)
	return email, true
}

// Admin login handler
func AdminLoginHandler(c *gin.Context) {
	var input models.Admin

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// find admin in DB
	var stored models.Admin
	collection := db.GetClient().Database("portfolio").Collection("admin")
	err := collection.FindOne(context.TODO(), bson.M{"email": input.Email}).Decode(&stored)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(stored.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Issue JWT
	token, err := makeToken(stored.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
		return
	}

	// HTTP cookie - browser stores it
	c.SetCookie(
		"admin_token", //name
		token,				// value
		3600*24,				// 24 hour
		"/",					// path
		"",	//domain
		false,				// secure
		true,				// httpOnly
	)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

// Admin verify for session check
func AdminVerifyHandler(c *gin.Context) {
	token, err := c.Cookie("admin_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	email, ok := parseToken(token)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"email": email})
}

// Admin logout
func AdminLogoutHandler(c *gin.Context) {
	// clear cookie
	c.SetCookie("admin_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}

func AdminAuthMiddleware(c *gin.Context) {
	token, err := c.Cookie("admin_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		c.Abort()
		return
	}

	_, ok := parseToken(token)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired session"})
		c.Abort()
		return
	}

	c.Next()
}