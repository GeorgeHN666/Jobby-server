package main

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
	ID       primitive.ObjectID `json:"_id"   bson:"_id"`
	FullName string             `json:"full_name" bson:"full_name"`
	Email    string             `json:"email"   bson:"email"`
	Password string             `json:"password"  bson:"password"`
	// AUTHORIZATION FIELDS
	Role      int       `json:"role"  bson:"role"`
	CreatedAt time.Time `json:"created_at"  bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at"   bson:"updated_at"`
	IsAuth    string    `json:"is_auth"   bson:"is_auth"`
	EmaiCode  string    `json:"email_code" bson:"email_code"`
	ResetCode string    `json:"reset_code" bson:"reset_code"`
	New       string    `json:"new"  bson:"new"`
}

type User struct {
	ID          primitive.ObjectID `json:"_id"  bson:"_id"`
	FullName    string             `json:"full_name"  bson:"full_name"`
	ProfilePic  primitive.ObjectID `json:"profile_pic,omitempty" bson:"profile_pic,omitempty"`
	Age         int                `json:"age,omitempty" bson:"age,omitempty"`
	Schooling   []string           `json:"schooling,omitempty" bson:"schooling,omitempty"`
	CV          primitive.ObjectID `json:"cv,omitempty" bson:"cv,omitempty"`
	Email       string             `json:"email"   bson:"email"`
	Password    string             `json:"password"  bson:"password"`
	Location    string             `json:"location,omitempty"   bson:"location,omitempty"`
	SocialMedia []string           `json:"social_media,omitempty"  bson:"social_media,omitempty"`
	Hobbies     []string           `json:"hobbies,omitempty"  bson:"hobbies,omitempty"`
	Skills      []string           `json:"skills,omitempty"  bson:"skills,omitempty"`
	Collections []string           `json:"collections"   bson:"collections"`
	// AUTHORIZATION FIELDS
	Role      int       `json:"role"  bson:"role"`
	CreatedAt time.Time `json:"created_at"  bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at"   bson:"updated_at"`
	IsAuth    string    `json:"is_auth"   bson:"is_auth"`
	EmaiCode  string    `json:"email_code" bson:"email_code"`
	ResetCode string    `json:"reset_code" bson:"reset_code"`
	New       string    `json:"new"  bson:"new"`
}

type Business struct {
	ID                 primitive.ObjectID `json:"_id"  bson:"_id"`
	Verified           string             `json:"verified"  bson:"verified"`
	About              string             `json:"about,omitempty"   bson:"about,omitempty"`
	ProfilePicture     primitive.ObjectID `json:"profile_pic,omitempty"   bson:"profile_pic,omitempty"`
	SocialMedia        []string           `json:"social_media,omitempty"  bson:"social_media,omitempty"`
	VerificationStatus string             `json:"verification_status"   bson:"verification_status"`
	Location           string             `json:"location"`
	// AUTHORIZATION FIELDS
	Role      int       `json:"role"  bson:"role"`
	CreatedAt time.Time `json:"created_at"  bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at"   bson:"updated_at"`
	IsAuth    string    `json:"is_auth"   bson:"is_auth"`
	EmaiCode  string    `json:"email_code" bson:"email_code"`
	ResetCode string    `json:"reset_code" bson:"reset_code"`
	New       string    `json:"new"  bson:"new"`
}

type SeekerPost struct {
	ID             primitive.ObjectID  `json:"_id"  bson:"_id"`
	UserInfo       User                `json:"user_info"   bson:"user_info"`
	JobDescription string              `json:"job_description"  bson:"job_description"`
	Collection     string              `json:"collection"   bson:"collection"`
	Location       string              `json:"location"  bson:"location"`
	ProofImages    map[int]interface{} `json:"proof_images"   bson:"proof_images"`
	CreatedAt      time.Time           `json:"created_at"   bson:"created_at"`
	UpdatedAt      time.Time           `json:"updated_at"   bson:"updated_at"`
}

type BusinessPost struct {
	ID             primitive.ObjectID  `json:"_id"   bson:"_id"`
	BusinessInfo   Business            `json:"business_info"   bson:"business_info"`
	JobDescription string              `json:"job_description"  bson:"job_description"`
	Vacants        int                 `json:"vacants"   bson:"vacants"`
	Collection     string              `json:"collection"   bson:"collection"`
	Location       string              `json:"location"   bson:"location"`
	Requisites     []string            `json:"requisites"`
	JobImages      map[int]interface{} `json:"job_images"`
	CreatedAt      time.Time           `json:"created_at"   bson:"created_at"`
	UpdatedAt      time.Time           `json:"updated_at"  bson:"updated_at"`
}

// Token Models
type Claims struct {
	Email string             `json:"emal"`
	ID    primitive.ObjectID `json:"_id"`
	jwt.StandardClaims
}
