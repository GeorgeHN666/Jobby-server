package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	// AUTHORIZATION FIELDS
	CreatedAt time.Time `json:"created_at"  bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at"   bson:"updated_at"`
	IsAuth    string    `json:"is_auth"   bson:"is_auth"`
	EmaiCode  string    `json:"email_code" bson:"email_code"`
	ResetCode string    `json:"reset_code" bson:"reset_code"`
	New       string    `json:"new"  bson:"new"`
}

type Business struct {
	Id             primitive.ObjectID `json:"_id"  bson:"_id"`
	Verified       string             `json:"verified"  bson:"verified"`
	About          string             `json:"about,omitempty"   bson:"about,omitempty"`
	ProfilePicture primitive.ObjectID `json:"profile_pic,omitempty"   bson:"profile_pic,omitempty"`
	SocialMedia    []string           `json:"social_media,omitempty"  bson:"social_media,omitempty"`
	// AUTHORIZATION FIELDS
	CreatedAt          time.Time `json:"created_at"  bson:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"   bson:"updated_at"`
	IsAuth             string    `json:"is_auth"   bson:"is_auth"`
	VerificationStatus string    `json:"verification_status"   bson:"verification_status"`
	EmaiCode           string    `json:"email_code" bson:"email_code"`
	ResetCode          string    `json:"reset_code" bson:"reset_code"`
	New                string    `json:"new"  bson:"new"`
}
