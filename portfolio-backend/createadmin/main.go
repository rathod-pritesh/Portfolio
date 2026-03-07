package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

type Admin struct {
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	CreatedAt time.Time `bson:"createdAt"`
}

func main() {
	// read connection string for DB
	mongoURI := os.Getenv("MONGO_URI")

	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "portfolio"
	}

	// connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("MongoDB connect error: %v", err)
	}
	defer client.Disconnect(context.Background())

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB ping failed: %v\nMake sure MongoDB is running at %s", err, mongoURI)
	}

	fmt.Println("Connected to MongoDB")

	collection := client.Database(dbName).Collection("admin")

	// Prompt for email
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter admin email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	if email == "" {
		log.Fatal("Email cannot be empty")
	}

	// check if already exists admin
	var existing Admin
	err = collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&existing)

	if err == nil {
		fmt.Printf("An admin with email '%s' already exists.\n", email)
		fmt.Print("Overwrite? (yes/no): ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(strings.ToLower(answer))
		if answer != "yes" && answer != "y" {
			fmt.Println("Aborted.")
			return
		}
	}

	// password
	fmt.Print("Enter password: ")
	passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		log.Fatalf("Failed to read password: %v", err)
	}

	// confirm password
	fmt.Print("Confirm password: ")
	confirmBytes, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		log.Fatalf("Failed to read password: %v", err)
	}

	if string(passwordBytes) != string(confirmBytes) {
		log.Fatal("Passwords do not match")
	}

	if len(passwordBytes) < 6 {
		log.Fatal("Password must be at least 6 characters")
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	admin := Admin{
		Email:     email,
		Password:  string(hash),
		CreatedAt: time.Now(),
	}

	// add into "admin" collection
	filter := bson.M{"email": email}
	update := bson.M{"$set": admin}
	opts := options.Update().SetUpsert(true)

	_, err = collection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		log.Fatalf("Failed to save admin: %v", err)
	}

	fmt.Printf("\n Admin '%s' created successfully in database '%s'.\n", email, dbName)
	fmt.Println(" You can now log in at /admin/login")
}
