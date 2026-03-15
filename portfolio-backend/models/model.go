package models

import "time"

type Home struct {
	ID           string `bson:"_id" json:"_id"`
	Name         string `bson:"name" json:"name"`
	Role         string `bson:"role" json:"role"`
	Tagline      string `bson:"tagline" json:"tagline"`
	ResumeUrl    string `bson:"resumeUrl" json:"resumeUrl"`
	ProfileImage string `bson:"profileImage" json:"profileImage"`
}

type Highlight struct {
	Label       string `bson:"label" json:"label"`
	Description string `bson:"description" json:"description"`
}

type About struct {
	ID          string      `bson:"_id" json:"_id"`
	Title       string      `bson:"title" json:"title"`
	Description string      `bson:"description" json:"description"`
	Highlights  []Highlight `bson:"highlights" json:"highlights"`
}

type Skill struct {
	Name string `bson:"name" json:"name"`
	Img  string `bson:"img" json:"img"`
}

type Category struct {
	Name   string  `bson:"name" json:"name"`
	Icon   string  `bson:"icon" json:"icon"`
	Order  int     `bson:"order" json:"order"`
	Skills []Skill `bson:"skills" json:"skills"`
}

type Skills struct {
	ID         string     `bson:"_id" json:"_id"`
	Title      string     `bson:"title" json:"title"`
	Categories []Category `bson:"categories" json:"categories"`
}

type Automation struct {
	ID          string              `bson:"_id" json:"_id"`
	Flow        string              `bson:"flow" json:"flow"`
	Title       string              `bson:"title" json:"title"`
	Description string              `bson:"description" json:"description"`
	WorkflowJson map[string]interface{} `bson:"workflowJson" json:"workflowJson"`
	Order       int                 `bson:"order" json:"order"`
}

type Certification struct {
	ID        string `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string `bson:"name" json:"name"`
	Company   string `bson:"company" json:"company"`
	Link      string `bson:"link" json:"link"`
	IssueDate string `bson:"issueDate" json:"issueDate"`
	Order     int    `bson:"order" json:"order"`
}

type Project struct {
	ID           string   `bson:"_id" json:"_id"`
	Title        string   `bson:"title" json:"title"`
	Description  string   `bson:"description" json:"description"`
	Image        string   `bson:"image" json:"image"`
	GitHub       string   `bson:"github" json:"github"`
	Technologies []string `bson:"technologies" json:"technologies"`
	Gradient     string   `bson:"gradient" json:"gradient"`
	Order        int      `bson:"order" json:"order"`
	IsFeatured   bool     `bson:"isFeatured" json:"isFeatured"`
}

type SocialLink struct {
	Name string `bson:"name" json:"name"`
	Url  string `bson:"url" json:"url"`
	Icon string `bson:"icon" json:"icon"`
}

type Contact struct {
	ID          string       `bson:"_id" json:"_id"`
	Email       string       `bson:"email" json:"email"`
	Location    string       `bson:"location" json:"location"`
	SocialLinks []SocialLink `bson:"sociallinks" json:"sociallinks"`
}

type Admin struct {
	Email     string    `bson:"email" json:"email"`
	Password  string    `bson:"password" json:"password"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}
