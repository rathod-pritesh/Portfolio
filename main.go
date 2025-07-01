package main

import (
	"context"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func initMongo(){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil{
		log.Fatal("Mongo connection error:", err)
	}

	db = client.Database("portfolio")
	log.Println("Connected to MongoDB portfolio DB")
}

type Content struct{
	HeroName 			string				`bson:"hero_name"`
	HeroDescription 	string				`bson:"hero_title"`
	ProfileImage    	string   			`bson:"profile_image"`
	ResumeFile 			string				`bson:"resume_file"`
	AboutTitle 			string				`bson:"about_title"`
	AboutText 			string				`bson:"about_content"`
	AboutHighlights 	[]string			`bson:"about_highlights"`
	Projects 			[]Project			`bson:"projects"`
	Certificates 		[]Certificate		`bson:"certificates"`
	ContactEmail 		string				`bson:"contact_email"`
	ContactLinkedIn 	string				`bson:"contact_linkedin"`
	ContactGitHub 		string				`bson:"contact_github"`
}

type Project struct{
	Title string
	Description string	
	Image string
	Stack []string
	GitHubLink string
}

type Certificate struct{
	Title string
	Issuer string
	Date string
	Icon string
}

type PageData struct{
	Projects []Project
	Certificates []Certificate
}

func main(){
	initMongo()
	//serve static files under /static/
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/admin/login", handleAdminLogin)
	http.HandleFunc("/admin/logout", handleAdminLogout)
	http.HandleFunc("/admin", handleAdminDashboard)
	http.HandleFunc("/admin/update_content", handleUpdateContent)
	http.HandleFunc("/download_resume", handleDownloadResume)

	http.HandleFunc("/", handleHome)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


func handleHome(w http.ResponseWriter, r *http.Request){
	// Only handle exact root path, not all paths
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	content, err := getContent()
	if err != nil{
		http.Error(w, "Failed to load content: "+err.Error(), http.StatusInternalServerError)
		return
	}
	
	renderTemplateWithData(w, "index.html", content)
}

func handleAdminLogin(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		tmplPath := filepath.Join("templates", "admin_login.html")
		t, err := template.ParseFiles(tmplPath)
		if err != nil{
			http.Error(w, "Template not found: "+err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
		return
	}
	// POST Method
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Hardcoded check
	if username == "admin" && password == "12345" {
		// set a cookie to simulate session
		http.SetCookie(w, &http.Cookie{
			Name: "admin_logged_in",
			Value: "true",
			Path: "/",
		})
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}

func handleAdminDashboard(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("admin_logged_in")
	if err != nil || cookie.Value != "true" {
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}

	content, _ := getContent()
	if err != nil{ 
		http.Error(w, "Failed to load content: "+err.Error(), http.StatusInternalServerError)
	}
	renderTemplateWithData(w, "admin_dashboard.html", content)
}

func handleUpdateContent(w http.ResponseWriter, r *http.Request){
	cookie, err := r.Cookie("admin_logged_in")
	if err != nil || cookie.Value != "true"{
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	err = r.ParseMultipartForm(10 << 20)	// up to 10 MB
	if err != nil {
		http.Error(w, "Error parsing formL: "+ err.Error(), http.StatusBadRequest)
		return
	}

	heroName := r.FormValue("hero_name")
	heroTitle := r.FormValue("hero_title")
	aboutTitle := r.FormValue("about_title")
	aboutText := r.FormValue("about_content")
	aboutHighlightsRaw := r.FormValue("about_highlights")

	updateSet := bson.M{}

	if heroName != ""{
		updateSet["hero_name"] = heroName
	}
	if heroTitle != ""{
		updateSet["hero_title"] = heroTitle
	}
	if aboutTitle != ""{
		updateSet["about_title"] = aboutTitle
	}
	if aboutText != ""{
		updateSet["about_content"] = aboutText
	}
	if aboutHighlightsRaw != ""{
		updateSet["about_highlights"] = strings.Split(aboutHighlightsRaw, ",")
	}

	update := bson.M{"$set": updateSet}

	// handle profile image upload
	file, handler, err := r.FormFile("profile_image")
	if err == nil && handler != nil{
		defer file.Close()
		dstPath := filepath.Join("static/images", handler.Filename)
		dst, err := os.Create(dstPath)
		if err == nil {
			defer dst.Close()
			_, err = io.Copy(dst, file)
		}
		if err == nil {
			// also update the DB with new profile image path
			update["$set"].(bson.M)["profile_image"] = handler.Filename
		}
	}

	//handle resume upload
	file, handler, err = r.FormFile("resume")
	if err == nil && handler != nil {
		defer file.Close()

		//save to static/resumes/
		os.MkdirAll("static/resumes", os.ModePerm)
		dstPath := filepath.Join("static/resumes", handler.Filename)
		dst, err := os.Create(dstPath)
		if err == nil {
			defer dst.Close()
			_, err = io.Copy(dst, file)
		}
		if err == nil {
			update["$set"].(bson.M)["resume_file"] = handler.Filename
		}
	}

	//update database
	collection := db.Collection("content")
	filter := bson.M{}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		http.Error(w, "Failed to update content: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func handleAdminLogout(w http.ResponseWriter, r *http.Request) {
	// Clear cookie
	http.SetCookie(w, &http.Cookie{
		Name: "admin_logged_in",
		Value: "",
		Path: "/",
		MaxAge: -1,
	})
	http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
}

func getContent() (Content, error){
	var content Content
	collection := db.Collection("content")
	err := collection.FindOne(context.TODO(), bson.M{}).Decode(&content)
	return content, err
}

func handleDownloadResume(w http.ResponseWriter, r *http.Request){
	content, err := getContent()
	if err != nil || content.ResumeFile == ""{
		http.Error(w, "No resume available", http.StatusNotFound)
		return
	}
	filepath := filepath.Join("static/resumes", content.ResumeFile)
	http.ServeFile(w, r, filepath)
}

func renderTemplateWithData(w http.ResponseWriter, tmpl string, content Content){
	tmplPath := filepath.Join("templates", tmpl)
	t, err := template.ParseFiles(tmplPath)
	if err != nil{
		http.Error(w,"Template not found: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Content
		AboutHighlightsStr string
	}{
		Content: content,
		AboutHighlightsStr: strings.Join(content.AboutHighlights, ","),
	}

	err = t.Execute(w, data)
	if err != nil{
		http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}